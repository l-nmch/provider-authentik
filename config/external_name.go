package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigurations applies the external name configuration to every
// resource. Authentik assigns the identifier of everything it stores (a UUID,
// or an integer primary key for providers, outposts and a few others) and
// returns it as the "id" attribute, which is also what "terraform import"
// takes. So every resource takes its external name from the provider, and
// there is no per resource table to maintain.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	}
}
