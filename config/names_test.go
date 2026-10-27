package config

import (
	"strings"
	"testing"
)

// TestResourceNames checks every Terraform resource has a name, that no two
// share one, and that the table holds nothing the schema lacks.
func TestResourceNames(t *testing.T) {
	p := GetProvider()
	seen := map[[2]string]string{}
	for name, r := range p.Resources {
		n, ok := resourceNames[name]
		if !ok {
			t.Errorf("%s has no entry in resourceNames", name)
			continue
		}
		if r.ShortGroup != n.group || r.Kind != n.kind {
			t.Errorf("%s: got %s/%s, want %s/%s", name, r.ShortGroup, r.Kind, n.group, n.kind)
		}
		key := [2]string{n.group, n.kind}
		if other, dup := seen[key]; dup {
			t.Errorf("%s and %s are both %s/%s", name, other, n.group, n.kind)
		}
		seen[key] = name
	}
	for name := range resourceNames {
		if !strings.HasPrefix(name, resourcePrefix+"_") {
			t.Errorf("%s does not start with %q", name, resourcePrefix)
		}
		if _, ok := p.Resources[name]; !ok {
			t.Errorf("%s is in resourceNames but not in the provider", name)
		}
	}
}
