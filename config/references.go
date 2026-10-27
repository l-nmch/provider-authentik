package config

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// references lists the arguments that hold the identifier of another resource
// and name that resource's Terraform type, so that each gets a
// <field>Ref/<field>Selector pair next to it. The Terraform provider's
// documentation has no example that links two resources, so Upjet finds none
// of these by itself.
//
// Only arguments with a single possible target are listed. A protocol provider
// (an application's, an outpost's or a policy binding's) or a stage (a flow
// stage binding's) can be any of a dozen kinds, and a reference names one, so
// those take the identifier itself: the external name of the referenced
// object. The same goes for a policy binding's target.
var references = map[string]map[string]string{
	"authentik_application_entitlement": {
		"application": "authentik_application",
	},
	"authentik_brand": {
		"client_certificates": "authentik_certificate_key_pair",
		"flow_authentication": "authentik_flow",
		"flow_device_code":    "authentik_flow",
		"flow_invalidation":   "authentik_flow",
		"flow_lockdown":       "authentik_flow",
		"flow_recovery":       "authentik_flow",
		"flow_unenrollment":   "authentik_flow",
		"flow_user_settings":  "authentik_flow",
		"web_certificate":     "authentik_certificate_key_pair",
	},
	"authentik_endpoints_connector_agent": {
		"authorization_flow": "authentik_flow",
	},
	"authentik_event_rule": {
		"destination_group": "authentik_group",
	},
	"authentik_event_transport": {
		"webhook_mapping_body":    "authentik_property_mapping_notification",
		"webhook_mapping_headers": "authentik_property_mapping_notification",
	},
	"authentik_flow_stage_binding": {
		"target": "authentik_flow",
	},
	"authentik_group": {
		"parents": "authentik_group",
		"roles":   "authentik_rbac_role",
		"users":   "authentik_user",
	},
	"authentik_outpost_provider_attachment": {
		"outpost": "authentik_outpost",
	},
	"authentik_policy_binding": {
		"group": "authentik_group",
		"user":  "authentik_user",
	},
	"authentik_provider_google_workspace": {
		"filter_group":      "authentik_group",
		"property_mappings": "authentik_property_mapping_provider_google_workspace",
	},
	"authentik_provider_ldap": {
		"bind_flow":   "authentik_flow",
		"certificate": "authentik_certificate_key_pair",
		"unbind_flow": "authentik_flow",
	},
	"authentik_provider_microsoft_entra": {
		"filter_group":      "authentik_group",
		"property_mappings": "authentik_property_mapping_provider_microsoft_entra",
	},
	"authentik_provider_oauth2": {
		"authentication_flow": "authentik_flow",
		"authorization_flow":  "authentik_flow",
		"invalidation_flow":   "authentik_flow",
		"property_mappings":   "authentik_property_mapping_provider_scope",
		"signing_key":         "authentik_certificate_key_pair",
	},
	"authentik_provider_oauth2_dcr": {
		"default_application_group":   "authentik_group",
		"oauth2_provider":             "authentik_provider_oauth2",
		"override_authorization_flow": "authentik_flow",
		"override_invalidation_flow":  "authentik_flow",
		"override_property_mappings":  "authentik_property_mapping_provider_scope",
	},
	"authentik_provider_proxy": {
		"authentication_flow": "authentik_flow",
		"authorization_flow":  "authentik_flow",
		"invalidation_flow":   "authentik_flow",
		"property_mappings":   "authentik_property_mapping_provider_scope",
	},
	"authentik_provider_rac": {
		"authentication_flow": "authentik_flow",
		"authorization_flow":  "authentik_flow",
		"property_mappings":   "authentik_property_mapping_provider_rac",
	},
	"authentik_provider_radius": {
		"authorization_flow": "authentik_flow",
		"certificate":        "authentik_certificate_key_pair",
		"invalidation_flow":  "authentik_flow",
		"property_mappings":  "authentik_property_mapping_provider_radius",
	},
	"authentik_provider_saml": {
		"authentication_flow":             "authentik_flow",
		"authn_context_class_ref_mapping": "authentik_property_mapping_provider_saml",
		"authorization_flow":              "authentik_flow",
		"encryption_kp":                   "authentik_certificate_key_pair",
		"invalidation_flow":               "authentik_flow",
		"name_id_mapping":                 "authentik_property_mapping_provider_saml",
		"property_mappings":               "authentik_property_mapping_provider_saml",
		"signing_kp":                      "authentik_certificate_key_pair",
		"verification_kp":                 "authentik_certificate_key_pair",
	},
	"authentik_provider_scim": {
		"property_mappings": "authentik_property_mapping_provider_scim",
	},
	"authentik_provider_ssf": {
		"signing_key": "authentik_certificate_key_pair",
	},
	"authentik_provider_ws_federation": {
		"authentication_flow":             "authentik_flow",
		"authn_context_class_ref_mapping": "authentik_property_mapping_provider_saml",
		"authorization_flow":              "authentik_flow",
		"encryption_kp":                   "authentik_certificate_key_pair",
		"invalidation_flow":               "authentik_flow",
		"name_id_mapping":                 "authentik_property_mapping_provider_saml",
		"signing_kp":                      "authentik_certificate_key_pair",
	},
	"authentik_rac_endpoint": {
		"property_mappings": "authentik_property_mapping_provider_rac",
	},
	"authentik_rbac_initial_permissions": {
		"role": "authentik_rbac_role",
	},
	"authentik_rbac_permission_role": {
		"role": "authentik_rbac_role",
	},
	"authentik_rbac_permission_user": {
		"user": "authentik_user",
	},
	"authentik_request_rule": {
		"request_flow": "authentik_flow",
	},
	"authentik_source_kerberos": {
		"authentication_flow": "authentik_flow",
		"enrollment_flow":     "authentik_flow",
	},
	"authentik_source_ldap": {
		"property_mappings": "authentik_property_mapping_source_ldap",
		"sync_parent_group": "authentik_group",
	},
	"authentik_source_oauth": {
		"authentication_flow": "authentik_flow",
		"enrollment_flow":     "authentik_flow",
		"property_mappings":   "authentik_property_mapping_source_oauth",
	},
	"authentik_source_plex": {
		"authentication_flow": "authentik_flow",
		"enrollment_flow":     "authentik_flow",
	},
	"authentik_source_saml": {
		"authentication_flow":     "authentik_flow",
		"encryption_kp":           "authentik_certificate_key_pair",
		"enrollment_flow":         "authentik_flow",
		"pre_authentication_flow": "authentik_flow",
		"property_mappings":       "authentik_property_mapping_source_saml",
		"signing_kp":              "authentik_certificate_key_pair",
		"verification_kp":         "authentik_certificate_key_pair",
	},
	"authentik_source_scim": {
		"property_mappings": "authentik_property_mapping_source_scim",
	},
	"authentik_source_telegram": {
		"authentication_flow":     "authentik_flow",
		"enrollment_flow":         "authentik_flow",
		"pre_authentication_flow": "authentik_flow",
	},
	"authentik_stage_account_lockdown": {
		"self_service_completion_flow": "authentik_flow",
	},
	"authentik_stage_authenticator_duo": {
		"configure_flow": "authentik_flow",
	},
	"authentik_stage_authenticator_email": {
		"configure_flow": "authentik_flow",
	},
	"authentik_stage_authenticator_endpoint_gdtc": {
		"configure_flow": "authentik_flow",
	},
	"authentik_stage_authenticator_sms": {
		"configure_flow": "authentik_flow",
		"mapping":        "authentik_property_mapping_notification",
	},
	"authentik_stage_authenticator_static": {
		"configure_flow": "authentik_flow",
	},
	"authentik_stage_authenticator_totp": {
		"configure_flow": "authentik_flow",
	},
	"authentik_stage_authenticator_webauthn": {
		"configure_flow": "authentik_flow",
	},
	"authentik_stage_identification": {
		"captcha_stage":     "authentik_stage_captcha",
		"enrollment_flow":   "authentik_flow",
		"password_stage":    "authentik_stage_password",
		"passwordless_flow": "authentik_flow",
		"recovery_flow":     "authentik_flow",
		"webauthn_stage":    "authentik_stage_authenticator_webauthn",
	},
	"authentik_stage_mutual_tls": {
		"certificate_authorities": "authentik_certificate_key_pair",
	},
	"authentik_stage_password": {
		"configure_flow": "authentik_flow",
	},
	"authentik_stage_redirect": {
		"target_flow": "authentik_flow",
	},
	"authentik_stage_user_write": {
		"create_users_group": "authentik_group",
	},
	"authentik_token": {
		"user": "authentik_user",
	},
	"authentik_user": {
		"groups": "authentik_group",
		"roles":  "authentik_rbac_role",
	},
	"authentik_user_offboarding": {
		"user": "authentik_user",
	},
}

// uuidExtractor makes a reference resolve to the "uuid" attribute of the
// target instead of its external name. The Terraform id of a flow or of an
// application is its slug, while every argument that points at one of them
// takes its UUID.
const uuidExtractor = "github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath(\"uuid\",true)"

// uuidTargets are the Terraform types whose id is not what a reference to them
// has to carry.
var uuidTargets = map[string]bool{
	"authentik_flow":        true,
	"authentik_application": true,
}

// addReferences sets the references of a resource.
func addReferences(r *ujconfig.Resource) {
	for field, target := range references[r.Name] {
		ref := ujconfig.Reference{TerraformName: target}
		if uuidTargets[target] {
			ref.Extractor = uuidExtractor
		}
		r.References[field] = ref
	}
}
