# provider-authentik

`provider-authentik` is a [Crossplane](https://crossplane.io/) provider for
[authentik](https://goauthentik.io/), generated with
[Upjet](https://github.com/crossplane/upjet) from
[`goauthentik/terraform-provider-authentik`](https://github.com/goauthentik/terraform-provider-authentik)
**v2026.8.0**. It targets Crossplane v2 and exposes every one of the 103
resources of the Terraform provider, both as namespaced
(`*.authentik.m.crossplane.io`) and cluster-scoped
(`*.authentik.crossplane.io`) managed resources.

The Terraform provider follows authentik's own versioning (`2026.8.x` for
authentik 2026.8), and the API it talks to changes between releases: keep
`TERRAFORM_PROVIDER_VERSION` on the major and minor version of the server.

## Resources

API groups and kinds are not the ones Upjet would derive from the Terraform
names (group `provider` with kind `Oauth2`, group `policy` with kind
`Binding`): they follow the model names of authentik's blueprints, so
`authentik_providers_oauth2.oauth2provider` is `OAuth2Provider`, and a stage,
a policy, a source or a property mapping always ends with `Stage`, `Policy`,
`Source` or `PropertyMapping`. The whole table is
[`config/resources.go`](config/resources.go), and a test checks it against the
schema.

| Group | Kinds |
|---|---|
| `core` | Agent, Application, ApplicationEntitlement, Blueprint, Brand, CertificateKeyPair, Group, InitialPermissions, ObjectAttribute, Role, RolePermission, SystemSettings, TaskSchedule, Token, User, UserOffboarding, UserPermission |
| `flow` | Flow, StageBinding |
| `stage` | 27 kinds: `CaptchaStage`, `IdentificationStage`, `PasswordStage`, `PromptStage`, `UserLoginStage`, `Authenticator...Stage`, ... and `PromptField` |
| `policy` | `DummyPolicy`, `EventMatcherPolicy`, `ExpiryPolicy`, `ExpressionPolicy`, `GeoIPPolicy`, `PasswordPolicy`, `ReputationPolicy`, `UniquePasswordPolicy` and `PolicyBinding` |
| `provider` | `OAuth2Provider`, `OAuth2DCRProvider`, `SAMLProvider`, `ProxyProvider`, `LDAPProvider`, `RACProvider`, `RACEndpoint`, `RadiusProvider`, `SCIMProvider`, `SSFProvider`, `WSFederationProvider`, `GoogleWorkspaceProvider`, `MicrosoftEntraProvider` |
| `source` | `OAuthSource`, `SAMLSource`, `LDAPSource`, `KerberosSource`, `PlexSource`, `SCIMSource`, `TelegramSource` |
| `propertymapping` | 14 kinds: `ProviderScopePropertyMapping`, `ProviderSAMLPropertyMapping`, `SourceLDAPPropertyMapping`, `NotificationPropertyMapping`, ... |
| `outpost` | `Outpost`, `ProviderAttachment`, `DockerServiceConnection`, `KubernetesServiceConnection` |
| `events` | `NotificationRule`, `NotificationTransport` |
| `endpoints` | `ConnectorAgent`, `ConnectorAgentEnrollmentToken`, `DeviceAccessGroup`, `GoogleChromeConnector` |
| `request` | `Rule`, `RuleBinding`, `RuleChildBinding` |
| `enterprise` | `License` |

RBAC kinds (`Role`, `RolePermission`, `UserPermission`, `InitialPermissions`)
are in `core` rather than in a group of their own: a group refers to its roles
and a role permission refers to a user, and Go does not allow the two packages
to import each other.

### Identity

authentik assigns the identifier of everything it stores, a UUID, or an
integer primary key for users, providers, outposts and a few others, and
returns it as the Terraform `id`. Every resource therefore takes its external
name from the provider: `metadata.name` is not the authentik identity (a
group's name is `spec.forProvider.name`), and the identifier becomes the
`crossplane.io/external-name` annotation once the object exists. Set that
annotation to an existing identifier to adopt an object.

### Cross-resource references

The Terraform provider's documentation has no example that links two
resources, so Upjet finds no reference by itself. [`config/references.go`](config/references.go)
declares them: 115 arguments of 49 resources get a `...Ref` and a `...Selector`
beside them, for the arguments that can only point at one kind of resource:

| Argument | Points at |
|---|---|
| `authorizationFlow`, `authenticationFlow`, `invalidationFlow`, `enrollmentFlow`, `configureFlow`, `recoveryFlow`, ... | `Flow` |
| `signingKey`, `signingKp`, `encryptionKp`, `verificationKp`, `certificate`, `webCertificate`, ... | `CertificateKeyPair` |
| `propertyMappings` | the property mapping kind of that provider or source |
| `groups`, `parents`, `filterGroup`, ... | `Group` |
| `users`, `user` | `User` |
| `roles`, `role` | `Role` |
| `application` of an entitlement | `Application` |

Where the target can be one of many kinds, there is no reference and the field
takes the identifier itself, the external name of the referenced object: an
application's or an outpost's `protocolProvider`, a policy binding's `policy`
and `target`, a flow stage binding's `stage`.

## Configuration

Create an authentik service account with a role that holds the permissions on
the objects to manage, create an API token for it, put them in a Secret as a
JSON document, and point a `ProviderConfig` (or `ClusterProviderConfig`) at it:

```json
{
  "url": "https://authentik.example.org",
  "token": "..."
}
```

| Key | Description |
|---|---|
| `url` | authentik URL (required). |
| `token` | API token (required). |
| `insecure` | Skip the verification of the server certificate. |
| `headers` | Extra HTTP headers sent with every request. |
| `ca_cert` | PEM certificate authority to trust in addition to the system ones. |

Keys this provider does not know are rejected rather than ignored, so a
misspelled `token` fails instead of looking like a missing one.

```yaml
apiVersion: authentik.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: my-namespace
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: my-namespace
      name: authentik-creds
      key: credentials
```

`ca_cert` is written into a trust bundle the Terraform provider plugin reads
through `SSL_CERT_FILE`. The plugin is a separate process whose trust store
belongs to the provider pod rather than to one ProviderConfig, so the bundle
holds the system certificates plus every authority any ProviderConfig
supplied: several of them compose instead of overwriting each other. Mounting
the authority into the pod and setting `SSL_CERT_FILE` through a
`DeploymentRuntimeConfig` has the same effect and keeps it out of the Secret.

## Usage notes

- **Secrets are references.** A user's `password`, a token's `key`, an OAuth
  client secret and the like are `...SecretRef` fields pointing at a
  Kubernetes Secret key, and the values authentik generates (a token's key, a
  client secret it makes up) are written to the connection secret
  (`writeConnectionSecretToRef`).
- **The scope of the token is the scope of the provider.** A token that can
  only read one group can only manage that group.
- **Blueprints and Crossplane overlap.** An object created by an authentik
  blueprint can be managed from here by adopting it with the
  `crossplane.io/external-name` annotation, but then both would try to own it.
  Pick one owner per object.

## Developing

Tools: Go (see `go.mod`), `goimports` on `PATH` (`go install tool` installs it),
Docker for images, `python3` for the package metadata.

```console
make generate        # schema, docs, CRDs, controllers, package metadata
make reviewable      # generate + lint + tests
make build           # provider image and xpkg under _output/
```

`config/provider-metadata.yaml` (the Terraform documentation the generated
CRDs take their field descriptions from) is produced with the scraper that
ships with Upjet, after `make pull-docs`:

```console
go run github.com/crossplane/upjet/v2/cmd/scraper -n authentik --resource-prefix authentik \
  -r .work/goauthentik/authentik/docs -o config/provider-metadata.yaml
```

Provider configuration lives in [`config/`](config/): the names in
[`resources.go`](config/resources.go), the references in
[`references.go`](config/references.go) and the external name in
[`external_name.go`](config/external_name.go). Credentials handling is in
[`internal/clients/authentik.go`](internal/clients/authentik.go).

### Implementation notes

The provider runs the `terraform` CLI with the Terraform provider plugin, served
from a filesystem mirror inside the image, the way `upjet-provider-template`
does (see [`cluster/images/provider-authentik`](cluster/images/provider-authentik)).
The plugin is built from source at the pinned tag, with `golang.org/x/net` and
`golang.org/x/text` raised ahead of upstream, and `hack/upstream/*.patch`
applied if there are any.

## License

Apache-2.0 for this provider. The image also ships Terraform 1.5.7 (MPL-2.0)
and `terraform-provider-authentik` (GPL-3.0), built from the unmodified source
of the pinned upstream tag and run as a separate process, as the other Upjet
providers do with theirs.
