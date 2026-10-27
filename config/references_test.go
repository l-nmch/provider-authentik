package config

import "testing"

// TestReferenceExtractors checks that a reference to a flow or to an
// application carries its UUID. The Terraform id of both is their slug, so the
// default extractor would put a slug where authentik expects a UUID.
func TestReferenceExtractors(t *testing.T) {
	p := GetProvider()
	for name, fields := range references {
		for field, target := range fields {
			got := p.Resources[name].References[field]
			if got.TerraformName != target {
				t.Errorf("%s.%s: points at %q, want %q", name, field, got.TerraformName, target)
			}
			if want := uuidTargets[target]; want != (got.Extractor == uuidExtractor) {
				t.Errorf("%s.%s -> %s: extractor %q, UUID wanted: %t", name, field, target, got.Extractor, want)
			}
		}
	}
}
