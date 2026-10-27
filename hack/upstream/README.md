# Upstream patches

Changes this provider needs in
[`goauthentik/terraform-provider-authentik`](https://github.com/goauthentik/terraform-provider-authentik),
kept here until they are merged and released. The provider image builds
`terraform-provider-authentik` from source at `TERRAFORM_PROVIDER_VERSION` (see
`cluster/images/provider-authentik/Dockerfile`) and applies every `*.patch` of
this directory before compiling, in file name order. Patches are plain `git
apply` diffs against that tag; a patch that no longer applies fails the image
build. Remove a patch once the pinned upstream release carries it.

There are none at the moment.
