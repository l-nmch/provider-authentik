// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	agent "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/agent"
	application "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/application"
	applicationentitlement "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/applicationentitlement"
	blueprint "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/blueprint"
	brand "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/brand"
	certificatekeypair "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/certificatekeypair"
	group "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/group"
	initialpermissions "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/initialpermissions"
	objectattribute "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/objectattribute"
	role "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/role"
	rolepermission "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/rolepermission"
	systemsettings "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/systemsettings"
	taskschedule "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/taskschedule"
	token "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/token"
	user "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/user"
	useroffboarding "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/useroffboarding"
	userpermission "github.com/l-nmch/provider-authentik/internal/controller/namespaced/core/userpermission"
	connectoragent "github.com/l-nmch/provider-authentik/internal/controller/namespaced/endpoints/connectoragent"
	connectoragentenrollmenttoken "github.com/l-nmch/provider-authentik/internal/controller/namespaced/endpoints/connectoragentenrollmenttoken"
	deviceaccessgroup "github.com/l-nmch/provider-authentik/internal/controller/namespaced/endpoints/deviceaccessgroup"
	googlechromeconnector "github.com/l-nmch/provider-authentik/internal/controller/namespaced/endpoints/googlechromeconnector"
	license "github.com/l-nmch/provider-authentik/internal/controller/namespaced/enterprise/license"
	notificationrule "github.com/l-nmch/provider-authentik/internal/controller/namespaced/events/notificationrule"
	notificationtransport "github.com/l-nmch/provider-authentik/internal/controller/namespaced/events/notificationtransport"
	flow "github.com/l-nmch/provider-authentik/internal/controller/namespaced/flow/flow"
	stagebinding "github.com/l-nmch/provider-authentik/internal/controller/namespaced/flow/stagebinding"
	dockerserviceconnection "github.com/l-nmch/provider-authentik/internal/controller/namespaced/outpost/dockerserviceconnection"
	kubernetesserviceconnection "github.com/l-nmch/provider-authentik/internal/controller/namespaced/outpost/kubernetesserviceconnection"
	outpost "github.com/l-nmch/provider-authentik/internal/controller/namespaced/outpost/outpost"
	providerattachment "github.com/l-nmch/provider-authentik/internal/controller/namespaced/outpost/providerattachment"
	dummypolicy "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/dummypolicy"
	eventmatcherpolicy "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/eventmatcherpolicy"
	expirypolicy "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/expirypolicy"
	expressionpolicy "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/expressionpolicy"
	geoippolicy "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/geoippolicy"
	passwordpolicy "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/passwordpolicy"
	policybinding "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/policybinding"
	reputationpolicy "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/reputationpolicy"
	uniquepasswordpolicy "github.com/l-nmch/provider-authentik/internal/controller/namespaced/policy/uniquepasswordpolicy"
	notificationpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/notificationpropertymapping"
	providergoogleworkspacepropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/providergoogleworkspacepropertymapping"
	providermicrosoftentrapropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/providermicrosoftentrapropertymapping"
	providerracpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/providerracpropertymapping"
	providerradiuspropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/providerradiuspropertymapping"
	providersamlpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/providersamlpropertymapping"
	providerscimpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/providerscimpropertymapping"
	providerscopepropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/providerscopepropertymapping"
	sourcekerberospropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/sourcekerberospropertymapping"
	sourceldappropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/sourceldappropertymapping"
	sourceoauthpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/sourceoauthpropertymapping"
	sourceplexpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/sourceplexpropertymapping"
	sourcesamlpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/sourcesamlpropertymapping"
	sourcescimpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/namespaced/propertymapping/sourcescimpropertymapping"
	googleworkspaceprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/googleworkspaceprovider"
	ldapprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/ldapprovider"
	microsoftentraprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/microsoftentraprovider"
	oauth2dcrprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/oauth2dcrprovider"
	oauth2provider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/oauth2provider"
	proxyprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/proxyprovider"
	racendpoint "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/racendpoint"
	racprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/racprovider"
	radiusprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/radiusprovider"
	samlprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/samlprovider"
	scimprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/scimprovider"
	ssfprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/ssfprovider"
	wsfederationprovider "github.com/l-nmch/provider-authentik/internal/controller/namespaced/provider/wsfederationprovider"
	providerconfig "github.com/l-nmch/provider-authentik/internal/controller/namespaced/providerconfig"
	rule "github.com/l-nmch/provider-authentik/internal/controller/namespaced/request/rule"
	rulebinding "github.com/l-nmch/provider-authentik/internal/controller/namespaced/request/rulebinding"
	rulechildbinding "github.com/l-nmch/provider-authentik/internal/controller/namespaced/request/rulechildbinding"
	kerberossource "github.com/l-nmch/provider-authentik/internal/controller/namespaced/source/kerberossource"
	ldapsource "github.com/l-nmch/provider-authentik/internal/controller/namespaced/source/ldapsource"
	oauthsource "github.com/l-nmch/provider-authentik/internal/controller/namespaced/source/oauthsource"
	plexsource "github.com/l-nmch/provider-authentik/internal/controller/namespaced/source/plexsource"
	samlsource "github.com/l-nmch/provider-authentik/internal/controller/namespaced/source/samlsource"
	scimsource "github.com/l-nmch/provider-authentik/internal/controller/namespaced/source/scimsource"
	telegramsource "github.com/l-nmch/provider-authentik/internal/controller/namespaced/source/telegramsource"
	accountlockdownstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/accountlockdownstage"
	authenticatorduostage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/authenticatorduostage"
	authenticatoremailstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/authenticatoremailstage"
	authenticatorendpointgdtcstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/authenticatorendpointgdtcstage"
	authenticatorsmsstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/authenticatorsmsstage"
	authenticatorstaticstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/authenticatorstaticstage"
	authenticatortotpstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/authenticatortotpstage"
	authenticatorvalidatestage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/authenticatorvalidatestage"
	authenticatorwebauthnstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/authenticatorwebauthnstage"
	captchastage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/captchastage"
	consentstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/consentstage"
	denystage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/denystage"
	dummystage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/dummystage"
	emailstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/emailstage"
	endpointsstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/endpointsstage"
	identificationstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/identificationstage"
	invitationstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/invitationstage"
	mutualtlsstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/mutualtlsstage"
	passwordstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/passwordstage"
	promptfield "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/promptfield"
	promptstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/promptstage"
	redirectstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/redirectstage"
	sourcestage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/sourcestage"
	userdeletestage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/userdeletestage"
	userloginstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/userloginstage"
	userlogoutstage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/userlogoutstage"
	userwritestage "github.com/l-nmch/provider-authentik/internal/controller/namespaced/stage/userwritestage"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agent.Setup,
		application.Setup,
		applicationentitlement.Setup,
		blueprint.Setup,
		brand.Setup,
		certificatekeypair.Setup,
		group.Setup,
		initialpermissions.Setup,
		objectattribute.Setup,
		role.Setup,
		rolepermission.Setup,
		systemsettings.Setup,
		taskschedule.Setup,
		token.Setup,
		user.Setup,
		useroffboarding.Setup,
		userpermission.Setup,
		connectoragent.Setup,
		connectoragentenrollmenttoken.Setup,
		deviceaccessgroup.Setup,
		googlechromeconnector.Setup,
		license.Setup,
		notificationrule.Setup,
		notificationtransport.Setup,
		flow.Setup,
		stagebinding.Setup,
		dockerserviceconnection.Setup,
		kubernetesserviceconnection.Setup,
		outpost.Setup,
		providerattachment.Setup,
		dummypolicy.Setup,
		eventmatcherpolicy.Setup,
		expirypolicy.Setup,
		expressionpolicy.Setup,
		geoippolicy.Setup,
		passwordpolicy.Setup,
		policybinding.Setup,
		reputationpolicy.Setup,
		uniquepasswordpolicy.Setup,
		notificationpropertymapping.Setup,
		providergoogleworkspacepropertymapping.Setup,
		providermicrosoftentrapropertymapping.Setup,
		providerracpropertymapping.Setup,
		providerradiuspropertymapping.Setup,
		providersamlpropertymapping.Setup,
		providerscimpropertymapping.Setup,
		providerscopepropertymapping.Setup,
		sourcekerberospropertymapping.Setup,
		sourceldappropertymapping.Setup,
		sourceoauthpropertymapping.Setup,
		sourceplexpropertymapping.Setup,
		sourcesamlpropertymapping.Setup,
		sourcescimpropertymapping.Setup,
		googleworkspaceprovider.Setup,
		ldapprovider.Setup,
		microsoftentraprovider.Setup,
		oauth2dcrprovider.Setup,
		oauth2provider.Setup,
		proxyprovider.Setup,
		racendpoint.Setup,
		racprovider.Setup,
		radiusprovider.Setup,
		samlprovider.Setup,
		scimprovider.Setup,
		ssfprovider.Setup,
		wsfederationprovider.Setup,
		providerconfig.Setup,
		rule.Setup,
		rulebinding.Setup,
		rulechildbinding.Setup,
		kerberossource.Setup,
		ldapsource.Setup,
		oauthsource.Setup,
		plexsource.Setup,
		samlsource.Setup,
		scimsource.Setup,
		telegramsource.Setup,
		accountlockdownstage.Setup,
		authenticatorduostage.Setup,
		authenticatoremailstage.Setup,
		authenticatorendpointgdtcstage.Setup,
		authenticatorsmsstage.Setup,
		authenticatorstaticstage.Setup,
		authenticatortotpstage.Setup,
		authenticatorvalidatestage.Setup,
		authenticatorwebauthnstage.Setup,
		captchastage.Setup,
		consentstage.Setup,
		denystage.Setup,
		dummystage.Setup,
		emailstage.Setup,
		endpointsstage.Setup,
		identificationstage.Setup,
		invitationstage.Setup,
		mutualtlsstage.Setup,
		passwordstage.Setup,
		promptfield.Setup,
		promptstage.Setup,
		redirectstage.Setup,
		sourcestage.Setup,
		userdeletestage.Setup,
		userloginstage.Setup,
		userlogoutstage.Setup,
		userwritestage.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agent.SetupGated,
		application.SetupGated,
		applicationentitlement.SetupGated,
		blueprint.SetupGated,
		brand.SetupGated,
		certificatekeypair.SetupGated,
		group.SetupGated,
		initialpermissions.SetupGated,
		objectattribute.SetupGated,
		role.SetupGated,
		rolepermission.SetupGated,
		systemsettings.SetupGated,
		taskschedule.SetupGated,
		token.SetupGated,
		user.SetupGated,
		useroffboarding.SetupGated,
		userpermission.SetupGated,
		connectoragent.SetupGated,
		connectoragentenrollmenttoken.SetupGated,
		deviceaccessgroup.SetupGated,
		googlechromeconnector.SetupGated,
		license.SetupGated,
		notificationrule.SetupGated,
		notificationtransport.SetupGated,
		flow.SetupGated,
		stagebinding.SetupGated,
		dockerserviceconnection.SetupGated,
		kubernetesserviceconnection.SetupGated,
		outpost.SetupGated,
		providerattachment.SetupGated,
		dummypolicy.SetupGated,
		eventmatcherpolicy.SetupGated,
		expirypolicy.SetupGated,
		expressionpolicy.SetupGated,
		geoippolicy.SetupGated,
		passwordpolicy.SetupGated,
		policybinding.SetupGated,
		reputationpolicy.SetupGated,
		uniquepasswordpolicy.SetupGated,
		notificationpropertymapping.SetupGated,
		providergoogleworkspacepropertymapping.SetupGated,
		providermicrosoftentrapropertymapping.SetupGated,
		providerracpropertymapping.SetupGated,
		providerradiuspropertymapping.SetupGated,
		providersamlpropertymapping.SetupGated,
		providerscimpropertymapping.SetupGated,
		providerscopepropertymapping.SetupGated,
		sourcekerberospropertymapping.SetupGated,
		sourceldappropertymapping.SetupGated,
		sourceoauthpropertymapping.SetupGated,
		sourceplexpropertymapping.SetupGated,
		sourcesamlpropertymapping.SetupGated,
		sourcescimpropertymapping.SetupGated,
		googleworkspaceprovider.SetupGated,
		ldapprovider.SetupGated,
		microsoftentraprovider.SetupGated,
		oauth2dcrprovider.SetupGated,
		oauth2provider.SetupGated,
		proxyprovider.SetupGated,
		racendpoint.SetupGated,
		racprovider.SetupGated,
		radiusprovider.SetupGated,
		samlprovider.SetupGated,
		scimprovider.SetupGated,
		ssfprovider.SetupGated,
		wsfederationprovider.SetupGated,
		providerconfig.SetupGated,
		rule.SetupGated,
		rulebinding.SetupGated,
		rulechildbinding.SetupGated,
		kerberossource.SetupGated,
		ldapsource.SetupGated,
		oauthsource.SetupGated,
		plexsource.SetupGated,
		samlsource.SetupGated,
		scimsource.SetupGated,
		telegramsource.SetupGated,
		accountlockdownstage.SetupGated,
		authenticatorduostage.SetupGated,
		authenticatoremailstage.SetupGated,
		authenticatorendpointgdtcstage.SetupGated,
		authenticatorsmsstage.SetupGated,
		authenticatorstaticstage.SetupGated,
		authenticatortotpstage.SetupGated,
		authenticatorvalidatestage.SetupGated,
		authenticatorwebauthnstage.SetupGated,
		captchastage.SetupGated,
		consentstage.SetupGated,
		denystage.SetupGated,
		dummystage.SetupGated,
		emailstage.SetupGated,
		endpointsstage.SetupGated,
		identificationstage.SetupGated,
		invitationstage.SetupGated,
		mutualtlsstage.SetupGated,
		passwordstage.SetupGated,
		promptfield.SetupGated,
		promptstage.SetupGated,
		redirectstage.SetupGated,
		sourcestage.SetupGated,
		userdeletestage.SetupGated,
		userloginstage.SetupGated,
		userlogoutstage.SetupGated,
		userwritestage.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		agent.SetupWebhookWithManager,
		application.SetupWebhookWithManager,
		applicationentitlement.SetupWebhookWithManager,
		blueprint.SetupWebhookWithManager,
		brand.SetupWebhookWithManager,
		certificatekeypair.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		initialpermissions.SetupWebhookWithManager,
		objectattribute.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		rolepermission.SetupWebhookWithManager,
		systemsettings.SetupWebhookWithManager,
		taskschedule.SetupWebhookWithManager,
		token.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		useroffboarding.SetupWebhookWithManager,
		userpermission.SetupWebhookWithManager,
		connectoragent.SetupWebhookWithManager,
		connectoragentenrollmenttoken.SetupWebhookWithManager,
		deviceaccessgroup.SetupWebhookWithManager,
		googlechromeconnector.SetupWebhookWithManager,
		license.SetupWebhookWithManager,
		notificationrule.SetupWebhookWithManager,
		notificationtransport.SetupWebhookWithManager,
		flow.SetupWebhookWithManager,
		stagebinding.SetupWebhookWithManager,
		dockerserviceconnection.SetupWebhookWithManager,
		kubernetesserviceconnection.SetupWebhookWithManager,
		outpost.SetupWebhookWithManager,
		providerattachment.SetupWebhookWithManager,
		dummypolicy.SetupWebhookWithManager,
		eventmatcherpolicy.SetupWebhookWithManager,
		expirypolicy.SetupWebhookWithManager,
		expressionpolicy.SetupWebhookWithManager,
		geoippolicy.SetupWebhookWithManager,
		passwordpolicy.SetupWebhookWithManager,
		policybinding.SetupWebhookWithManager,
		reputationpolicy.SetupWebhookWithManager,
		uniquepasswordpolicy.SetupWebhookWithManager,
		notificationpropertymapping.SetupWebhookWithManager,
		providergoogleworkspacepropertymapping.SetupWebhookWithManager,
		providermicrosoftentrapropertymapping.SetupWebhookWithManager,
		providerracpropertymapping.SetupWebhookWithManager,
		providerradiuspropertymapping.SetupWebhookWithManager,
		providersamlpropertymapping.SetupWebhookWithManager,
		providerscimpropertymapping.SetupWebhookWithManager,
		providerscopepropertymapping.SetupWebhookWithManager,
		sourcekerberospropertymapping.SetupWebhookWithManager,
		sourceldappropertymapping.SetupWebhookWithManager,
		sourceoauthpropertymapping.SetupWebhookWithManager,
		sourceplexpropertymapping.SetupWebhookWithManager,
		sourcesamlpropertymapping.SetupWebhookWithManager,
		sourcescimpropertymapping.SetupWebhookWithManager,
		googleworkspaceprovider.SetupWebhookWithManager,
		ldapprovider.SetupWebhookWithManager,
		microsoftentraprovider.SetupWebhookWithManager,
		oauth2dcrprovider.SetupWebhookWithManager,
		oauth2provider.SetupWebhookWithManager,
		proxyprovider.SetupWebhookWithManager,
		racendpoint.SetupWebhookWithManager,
		racprovider.SetupWebhookWithManager,
		radiusprovider.SetupWebhookWithManager,
		samlprovider.SetupWebhookWithManager,
		scimprovider.SetupWebhookWithManager,
		ssfprovider.SetupWebhookWithManager,
		wsfederationprovider.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		rule.SetupWebhookWithManager,
		rulebinding.SetupWebhookWithManager,
		rulechildbinding.SetupWebhookWithManager,
		kerberossource.SetupWebhookWithManager,
		ldapsource.SetupWebhookWithManager,
		oauthsource.SetupWebhookWithManager,
		plexsource.SetupWebhookWithManager,
		samlsource.SetupWebhookWithManager,
		scimsource.SetupWebhookWithManager,
		telegramsource.SetupWebhookWithManager,
		accountlockdownstage.SetupWebhookWithManager,
		authenticatorduostage.SetupWebhookWithManager,
		authenticatoremailstage.SetupWebhookWithManager,
		authenticatorendpointgdtcstage.SetupWebhookWithManager,
		authenticatorsmsstage.SetupWebhookWithManager,
		authenticatorstaticstage.SetupWebhookWithManager,
		authenticatortotpstage.SetupWebhookWithManager,
		authenticatorvalidatestage.SetupWebhookWithManager,
		authenticatorwebauthnstage.SetupWebhookWithManager,
		captchastage.SetupWebhookWithManager,
		consentstage.SetupWebhookWithManager,
		denystage.SetupWebhookWithManager,
		dummystage.SetupWebhookWithManager,
		emailstage.SetupWebhookWithManager,
		endpointsstage.SetupWebhookWithManager,
		identificationstage.SetupWebhookWithManager,
		invitationstage.SetupWebhookWithManager,
		mutualtlsstage.SetupWebhookWithManager,
		passwordstage.SetupWebhookWithManager,
		promptfield.SetupWebhookWithManager,
		promptstage.SetupWebhookWithManager,
		redirectstage.SetupWebhookWithManager,
		sourcestage.SetupWebhookWithManager,
		userdeletestage.SetupWebhookWithManager,
		userloginstage.SetupWebhookWithManager,
		userlogoutstage.SetupWebhookWithManager,
		userwritestage.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
