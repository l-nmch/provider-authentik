package clients

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	clusterv1beta1 "github.com/l-nmch/provider-authentik/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/l-nmch/provider-authentik/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal authentik credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds, err := parseCredentials(data)
		if err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		cfg, err := creds.terraformConfiguration(os.TempDir())
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		ps.Configuration = cfg
		return ps, nil
	}
}

// parseCredentials reads the credentials document, rejecting any key it does
// not know. An ignored key is worse than an error here: a misspelled token
// would look like a missing one, and options the Authentik Terraform provider
// has no equivalent for would look supported.
func parseCredentials(data []byte) (Credentials, error) {
	creds := Credentials{}
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&creds); err != nil {
		return Credentials{}, err
	}
	return creds, nil
}

// Credentials is the JSON document expected in the ProviderConfig credentials
// source. It mirrors the configuration arguments of the Terraform provider.
type Credentials struct {
	// URL of the Authentik server, e.g. https://authentik.example.org.
	URL string `json:"url"`
	// Token is an Authentik API token, usually the one of a service account.
	Token string `json:"token"`
	// Insecure skips the verification of the server certificate.
	Insecure bool `json:"insecure,omitempty"`
	// Headers are extra HTTP headers sent with every request.
	Headers map[string]string `json:"headers,omitempty"`
	// CACert is a PEM encoded certificate authority to trust in addition to
	// the system ones, for a server whose certificate it signed.
	CACert string `json:"ca_cert,omitempty"`
}

// terraformConfiguration validates the credentials and returns the matching
// Terraform provider configuration, leaving out every argument that was not
// set.
func (c Credentials) terraformConfiguration(baseDir string) (map[string]any, error) {
	if c.URL == "" {
		return nil, errors.New("url is required")
	}
	if !strings.HasPrefix(c.URL, "https://") && !strings.HasPrefix(c.URL, "http://") {
		return nil, errors.New("url must start with https:// or http://")
	}
	if c.Token == "" {
		return nil, errors.New("token is required")
	}

	cfg := map[string]any{"url": c.URL, "token": c.Token}
	if c.Insecure {
		cfg["insecure"] = true
	}
	if len(c.Headers) > 0 {
		headers := make(map[string]any, len(c.Headers))
		for k, v := range c.Headers {
			headers[k] = v
		}
		cfg["headers"] = headers
	}

	if c.CACert != "" {
		if err := trustCA(baseDir, c.CACert); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

// systemCABundles are the usual locations of the system trust store, tried in
// order. The provider image is Alpine based, where ca-certificates installs
// the first one.
var systemCABundles = []string{
	"/etc/ssl/certs/ca-certificates.crt",
	"/etc/pki/tls/certs/ca-bundle.crt",
	"/etc/ssl/ca-bundle.pem",
	"/etc/ssl/cert.pem",
}

var caTrust = struct {
	sync.Mutex
	// added is the set of certificates already in the bundle, keyed by digest.
	added map[string]struct{}
	// path of the bundle, once it has been written.
	path string
}{added: map[string]struct{}{}}

// trustCA adds a PEM encoded certificate authority to the trust store of the
// Terraform provider plugin. The plugin is a separate Go process, and neither
// the Authentik Terraform provider nor upjet's provider configuration can carry
// a certificate authority, so the only way in is SSL_CERT_FILE, which Go reads
// instead of the system trust store. That makes trust a property of the
// provider pod rather than of a single ProviderConfig: the bundle holds the
// system certificates plus every authority any ProviderConfig supplied, so
// several of them compose instead of overwriting each other. Mounting the
// authority into the pod and setting SSL_CERT_FILE through a
// DeploymentRuntimeConfig has the same effect and keeps it out of the Secret.
func trustCA(baseDir, cert string) error {
	if p, _ := pem.Decode([]byte(cert)); p == nil {
		return errors.New("ca_cert is not a PEM encoded certificate")
	}
	sum := sha256.Sum256([]byte(cert))
	digest := hex.EncodeToString(sum[:])

	caTrust.Lock()
	defer caTrust.Unlock()
	if _, ok := caTrust.added[digest]; ok {
		return nil
	}

	bundle := caTrust.path
	if bundle == "" {
		dir := filepath.Join(baseDir, "provider-authentik-ca")
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return errors.Wrap(err, "cannot create the trust store directory")
		}
		bundle = filepath.Join(dir, "ca-bundle.crt")
		if err := writeFileAtomic(bundle, systemCertificates()); err != nil {
			return err
		}
		if err := os.Setenv("SSL_CERT_FILE", bundle); err != nil {
			return errors.Wrap(err, "cannot point the trust store at the bundle")
		}
		caTrust.path = bundle
	}

	f, err := os.OpenFile(bundle, os.O_APPEND|os.O_WRONLY, 0o600) //nolint:gosec // path is ours
	if err != nil {
		return errors.Wrap(err, "cannot open the trust bundle")
	}
	defer f.Close() //nolint:errcheck // the write below is what matters
	if _, err := f.WriteString("\n" + strings.TrimSpace(cert) + "\n"); err != nil {
		return errors.Wrap(err, "cannot add the certificate authority to the trust bundle")
	}
	caTrust.added[digest] = struct{}{}
	return nil
}

// systemCertificates returns the contents of the system trust store, so that
// adding an authority to the bundle does not drop the public ones. It returns
// nothing when there is none, which only means the server certificate has to
// be signed by the supplied authority.
func systemCertificates() []byte {
	for _, p := range systemCABundles {
		if b, err := os.ReadFile(p); err == nil { //nolint:gosec // fixed list of system paths
			return b
		}
	}
	return nil
}

// writeFileAtomic writes content to path with owner-only permissions, unless
// the file already holds exactly that content.
func writeFileAtomic(path string, content []byte) error {
	if existing, err := os.ReadFile(path); err == nil && bytes.Equal(existing, content) { //nolint:gosec // path is built from a hash
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return errors.Wrap(err, "cannot create the trust store directory")
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".tmp-*")
	if err != nil {
		return errors.Wrap(err, "cannot create temporary file")
	}
	defer os.Remove(tmp.Name()) //nolint:errcheck // best effort cleanup, renamed on success
	if _, err := tmp.Write(content); err != nil {
		tmp.Close() //nolint:errcheck,gosec // already failing
		return errors.Wrap(err, "cannot write temporary file")
	}
	if err := tmp.Close(); err != nil {
		return errors.Wrap(err, "cannot close temporary file")
	}
	return errors.Wrap(os.Rename(tmp.Name(), path), "cannot write the trust bundle")
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged: //nolint:staticcheck // still handling cluster-scoped behavior
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) { //nolint:staticcheck // still handling cluster-scoped behavior
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New(" is not an Object")
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
