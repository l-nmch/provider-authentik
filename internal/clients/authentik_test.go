package clients

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	tlsURL = "https://authentik.example.org"
	secret = "t0ps3cr3t11"
)

func TestParseCredentials(t *testing.T) {
	got, err := parseCredentials([]byte(`{"url":"https://authentik.example.org","token":"t0ps3cr3t11","insecure":true,"headers":{"X-Test":"1"}}`))
	if err != nil {
		t.Fatalf("parseCredentials: %v", err)
	}
	want := Credentials{URL: tlsURL, Token: secret, Insecure: true, Headers: map[string]string{"X-Test": "1"}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("credentials: %s", diff)
	}
}

func TestParseCredentialsRejectsUnknownKeys(t *testing.T) {
	// A key we silently ignored would be worse than an error: a misspelled
	// token reads as a missing one, and an option the Authentik Terraform
	// provider has no equivalent for reads as supported.
	for name, doc := range map[string]string{
		"Misspelled":  `{"url":"https://authentik.example.org","tokem":"secret"}`,
		"Unsupported": `{"url":"https://authentik.example.org","insecure_skip_verify":true}`,
		"OpenFGAKeys": `{"api_url":"https://authentik.example.org","api_token":"secret"}`,
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := parseCredentials([]byte(doc)); err == nil {
				t.Error("expected an error, got none")
			}
		})
	}
}

func TestTerraformConfiguration(t *testing.T) {
	for name, tc := range map[string]struct {
		creds Credentials
		want  map[string]any
	}{
		"Minimal": {
			creds: Credentials{URL: tlsURL, Token: secret},
			want:  map[string]any{"url": tlsURL, "token": secret},
		},
		"InsecureAndHeaders": {
			creds: Credentials{URL: tlsURL, Token: secret, Insecure: true, Headers: map[string]string{"X-Test": "1"}},
			want: map[string]any{
				"url":      tlsURL,
				"token":    secret,
				"insecure": true,
				"headers":  map[string]any{"X-Test": "1"},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := tc.creds.terraformConfiguration(t.TempDir())
			if err != nil {
				t.Fatalf("terraformConfiguration: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("configuration: %s", diff)
			}
		})
	}
}

func TestTerraformConfigurationErrors(t *testing.T) {
	for name, tc := range map[string]struct {
		creds   Credentials
		wantErr string
	}{
		"NoURL":   {creds: Credentials{Token: secret}, wantErr: "url is required"},
		"BadURL":  {creds: Credentials{URL: "authentik.example.org", Token: secret}, wantErr: "must start with"},
		"NoToken": {creds: Credentials{URL: tlsURL}, wantErr: "token is required"},
		"BadCA":   {creds: Credentials{URL: tlsURL, Token: secret, CACert: "nope"}, wantErr: "not a PEM"},
	} {
		t.Run(name, func(t *testing.T) {
			_, err := tc.creds.terraformConfiguration(t.TempDir())
			if err == nil {
				t.Fatal("expected an error, got none")
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("error %q does not mention %q", err, tc.wantErr)
			}
		})
	}
}

// testCA is a throwaway self-signed certificate, only ever parsed as PEM.
const testCA = `-----BEGIN CERTIFICATE-----
MIIBITAdBgNVBAMTFmNyb3NzcGxhbmUtdGVzdC1jYS1vbmx5MA0GCSqGSIb3DQEB
CwUAA0EAqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq
qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq
-----END CERTIFICATE-----`

const testCA2 = `-----BEGIN CERTIFICATE-----
MIIBITAdBgNVBAMTFmNyb3NzcGxhbmUtdGVzdC1jYS10d28wDQYJKoZIhvcNAQEL
BQADQQCrq6urq6urq6urq6urq6urq6urq6urq6urq6urq6urq6urq6urq6urq6ur
q6urq6urq6urq6urq6urq6urq6urq6urq6ur
-----END CERTIFICATE-----`

// TestTrustCA checks that authorities compose instead of overwriting each
// other, since the trust store belongs to the provider pod rather than to a
// single ProviderConfig.
func TestTrustCA(t *testing.T) {
	t.Setenv("SSL_CERT_FILE", "")
	caTrust.Lock()
	caTrust.added = map[string]struct{}{}
	caTrust.path = ""
	caTrust.Unlock()
	t.Cleanup(func() {
		caTrust.Lock()
		caTrust.added = map[string]struct{}{}
		caTrust.path = ""
		caTrust.Unlock()
	})

	dir := t.TempDir()
	for _, cert := range []string{testCA, testCA2, testCA} {
		if err := trustCA(dir, cert); err != nil {
			t.Fatalf("trustCA: %v", err)
		}
	}

	bundle := os.Getenv("SSL_CERT_FILE")
	if bundle == "" {
		t.Fatal("SSL_CERT_FILE was not set")
	}
	b, err := os.ReadFile(bundle) //nolint:gosec // path comes from trustCA
	if err != nil {
		t.Fatalf("cannot read the bundle: %v", err)
	}
	got := string(b)

	// Both authorities are in, each exactly once despite the repeat.
	for _, cert := range []string{testCA, testCA2} {
		if n := strings.Count(got, strings.SplitN(cert, "\n", 2)[1]); n != 1 {
			t.Errorf("certificate appears %d times, want 1", n)
		}
	}
	// The system authorities, when the host has any, must survive: pointing
	// SSL_CERT_FILE at a bundle replaces the system trust store rather than
	// adding to it.
	if sys := systemCertificates(); len(sys) > 0 && !strings.HasPrefix(got, string(sys)) {
		t.Error("the bundle does not start with the system certificates")
	}
}
