# provider-authentik

Crossplane provider for [authentik](https://goauthentik.io/): manage
applications, providers, flows, stages, policies, users, groups and the rest of
what authentik exposes as Kubernetes resources.

Generated with [Upjet](https://github.com/crossplane/upjet) from
[`goauthentik/terraform-provider-authentik`](https://github.com/goauthentik/terraform-provider-authentik)
v2026.8.0, so it follows the versioning of authentik itself: use the provider
release that matches the major and minor version of your server. Requires
**Crossplane v2**. Source, issues and full documentation:
https://github.com/l-nmch/provider-authentik

## Install

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-authentik
spec:
  package: xpkg.upbound.io/l-nmch/provider-authentik:v0.1.0
```

```console
kubectl apply -f provider.yaml
kubectl get providers.pkg.crossplane.io provider-authentik   # wait for HEALTHY=True
```

The provider pods must be able to reach the authentik API over HTTPS.

## Configure

### 1. Create an API token

Create a service account in authentik, give it a role holding the permissions
it needs on the objects you want to manage, and create an API token for it.
The provider acts with exactly the permissions of that account.

### 2. Store the credentials in a Secret

The Secret holds a JSON document. `url` and `token` are required; unknown keys
are rejected rather than ignored.

| Key | Description |
|---|---|
| `url` | authentik URL (required), e.g. `https://authentik.example.org` |
| `token` | API token (required) |
| `insecure` | Skip the verification of the server certificate |
| `headers` | Extra HTTP headers sent with every request, as a JSON object |
| `ca_cert` | PEM certificate authority to trust in addition to the system ones |

```console
kubectl -n crossplane-system create secret generic authentik-creds \
  --from-literal=credentials='{"url":"https://authentik.example.org","token":"'"${AUTHENTIK_TOKEN}"'"}'
```

### 3. Create a ClusterProviderConfig

```yaml
apiVersion: authentik.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: authentik-creds
      key: credentials
```

Namespaced managed resources use the `ClusterProviderConfig` named `default`
unless they set `providerConfigRef`. A namespaced `ProviderConfig` is also
available.

## Example

```yaml
apiVersion: core.authentik.m.crossplane.io/v1alpha1
kind: Group
metadata:
  name: developers
  namespace: default
spec:
  forProvider:
    name: developers
---
apiVersion: provider.authentik.m.crossplane.io/v1alpha1
kind: OAuth2Provider
metadata:
  name: grafana
  namespace: default
spec:
  forProvider:
    name: grafana
    clientId: grafana
    # The flows are Flow managed resources, referenced by name:
    authorizationFlowRef:
      name: grafana-authorization
    invalidationFlowRef:
      name: grafana-invalidation
---
apiVersion: core.authentik.m.crossplane.io/v1alpha1
kind: Application
metadata:
  name: grafana
  namespace: default
spec:
  forProvider:
    name: Grafana
    slug: grafana
    protocolProvider: 1
```

## Resources

All 103 resources of the Terraform provider, each available namespaced
(`*.authentik.m.crossplane.io`) and cluster-scoped (`*.authentik.crossplane.io`).
Kinds are named after authentik's blueprint models, so `OAuth2Provider`,
`ExpressionPolicy` or `IdentificationStage` are what
`authentik_providers_oauth2.oauth2provider` and the like are called in a
blueprint.

| Group | Kinds |
|---|---|
| `core` | Application, ApplicationEntitlement, Brand, Group, User, Token, Role, RolePermission, UserPermission, CertificateKeyPair, Blueprint, SystemSettings, ... |
| `flow` | Flow, StageBinding |
| `stage` | one `...Stage` kind per stage type, and PromptField |
| `policy` | one `...Policy` kind per policy type, and PolicyBinding |
| `provider` | OAuth2Provider, SAMLProvider, ProxyProvider, LDAPProvider, ... |
| `source` | OAuthSource, SAMLSource, LDAPSource, ... |
| `propertymapping` | one `...PropertyMapping` kind per mapping type |
| `outpost` | Outpost, ProviderAttachment, service connections |
| `events` | NotificationRule, NotificationTransport |
| `endpoints`, `request`, `enterprise` | the corresponding authentik features |

Good to know:

- authentik assigns the identifier of everything it stores, and it becomes the
  `crossplane.io/external-name` annotation once the object exists. Set that
  annotation to an existing identifier to adopt an object.
- Fields that hold the identifier of another object have a `...Ref` and a
  `...Selector` next to them where the target is a single kind of resource
  (flows, key pairs, groups, users, roles, property mappings). Where it can be
  one of many kinds, as for an application's `protocolProvider`, the field
  takes the identifier itself.
- Secrets (a user's password, a client secret, a token key) are given as
  references to Kubernetes Secrets, never inline.

See the [README](https://github.com/l-nmch/provider-authentik#readme) for details.

## License

Apache-2.0. The image also ships Terraform 1.5.7 (MPL-2.0) and
terraform-provider-authentik v2026.8.0 (GPL-3.0), built from the unmodified
upstream source at that tag and run as a separate process. This is a community
provider, not affiliated with authentik or Crossplane.
