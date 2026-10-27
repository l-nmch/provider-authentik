// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	agent "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/agent"
	application "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/application"
	applicationentitlement "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/applicationentitlement"
	blueprint "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/blueprint"
	brand "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/brand"
	certificatekeypair "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/certificatekeypair"
	group "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/group"
	initialpermissions "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/initialpermissions"
	objectattribute "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/objectattribute"
	role "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/role"
	rolepermission "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/rolepermission"
	systemsettings "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/systemsettings"
	taskschedule "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/taskschedule"
	token "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/token"
	user "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/user"
	useroffboarding "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/useroffboarding"
	userpermission "github.com/l-nmch/provider-authentik/internal/controller/cluster/core/userpermission"
	connectoragent "github.com/l-nmch/provider-authentik/internal/controller/cluster/endpoints/connectoragent"
	connectoragentenrollmenttoken "github.com/l-nmch/provider-authentik/internal/controller/cluster/endpoints/connectoragentenrollmenttoken"
	deviceaccessgroup "github.com/l-nmch/provider-authentik/internal/controller/cluster/endpoints/deviceaccessgroup"
	googlechromeconnector "github.com/l-nmch/provider-authentik/internal/controller/cluster/endpoints/googlechromeconnector"
	license "github.com/l-nmch/provider-authentik/internal/controller/cluster/enterprise/license"
	notificationrule "github.com/l-nmch/provider-authentik/internal/controller/cluster/events/notificationrule"
	notificationtransport "github.com/l-nmch/provider-authentik/internal/controller/cluster/events/notificationtransport"
	flow "github.com/l-nmch/provider-authentik/internal/controller/cluster/flow/flow"
	stagebinding "github.com/l-nmch/provider-authentik/internal/controller/cluster/flow/stagebinding"
	dockerserviceconnection "github.com/l-nmch/provider-authentik/internal/controller/cluster/outpost/dockerserviceconnection"
	kubernetesserviceconnection "github.com/l-nmch/provider-authentik/internal/controller/cluster/outpost/kubernetesserviceconnection"
	outpost "github.com/l-nmch/provider-authentik/internal/controller/cluster/outpost/outpost"
	providerattachment "github.com/l-nmch/provider-authentik/internal/controller/cluster/outpost/providerattachment"
	dummypolicy "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/dummypolicy"
	eventmatcherpolicy "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/eventmatcherpolicy"
	expirypolicy "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/expirypolicy"
	expressionpolicy "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/expressionpolicy"
	geoippolicy "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/geoippolicy"
	passwordpolicy "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/passwordpolicy"
	policybinding "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/policybinding"
	reputationpolicy "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/reputationpolicy"
	uniquepasswordpolicy "github.com/l-nmch/provider-authentik/internal/controller/cluster/policy/uniquepasswordpolicy"
	notificationpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/notificationpropertymapping"
	providergoogleworkspacepropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/providergoogleworkspacepropertymapping"
	providermicrosoftentrapropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/providermicrosoftentrapropertymapping"
	providerracpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/providerracpropertymapping"
	providerradiuspropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/providerradiuspropertymapping"
	providersamlpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/providersamlpropertymapping"
	providerscimpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/providerscimpropertymapping"
	providerscopepropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/providerscopepropertymapping"
	sourcekerberospropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/sourcekerberospropertymapping"
	sourceldappropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/sourceldappropertymapping"
	sourceoauthpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/sourceoauthpropertymapping"
	sourceplexpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/sourceplexpropertymapping"
	sourcesamlpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/sourcesamlpropertymapping"
	sourcescimpropertymapping "github.com/l-nmch/provider-authentik/internal/controller/cluster/propertymapping/sourcescimpropertymapping"
	googleworkspaceprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/googleworkspaceprovider"
	ldapprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/ldapprovider"
	microsoftentraprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/microsoftentraprovider"
	oauth2dcrprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/oauth2dcrprovider"
	oauth2provider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/oauth2provider"
	proxyprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/proxyprovider"
	racendpoint "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/racendpoint"
	racprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/racprovider"
	radiusprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/radiusprovider"
	samlprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/samlprovider"
	scimprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/scimprovider"
	ssfprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/ssfprovider"
	wsfederationprovider "github.com/l-nmch/provider-authentik/internal/controller/cluster/provider/wsfederationprovider"
	providerconfig "github.com/l-nmch/provider-authentik/internal/controller/cluster/providerconfig"
	rule "github.com/l-nmch/provider-authentik/internal/controller/cluster/request/rule"
	rulebinding "github.com/l-nmch/provider-authentik/internal/controller/cluster/request/rulebinding"
	rulechildbinding "github.com/l-nmch/provider-authentik/internal/controller/cluster/request/rulechildbinding"
	kerberossource "github.com/l-nmch/provider-authentik/internal/controller/cluster/source/kerberossource"
	ldapsource "github.com/l-nmch/provider-authentik/internal/controller/cluster/source/ldapsource"
	oauthsource "github.com/l-nmch/provider-authentik/internal/controller/cluster/source/oauthsource"
	plexsource "github.com/l-nmch/provider-authentik/internal/controller/cluster/source/plexsource"
	samlsource "github.com/l-nmch/provider-authentik/internal/controller/cluster/source/samlsource"
	scimsource "github.com/l-nmch/provider-authentik/internal/controller/cluster/source/scimsource"
	telegramsource "github.com/l-nmch/provider-authentik/internal/controller/cluster/source/telegramsource"
	accountlockdownstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/accountlockdownstage"
	authenticatorduostage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/authenticatorduostage"
	authenticatoremailstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/authenticatoremailstage"
	authenticatorendpointgdtcstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/authenticatorendpointgdtcstage"
	authenticatorsmsstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/authenticatorsmsstage"
	authenticatorstaticstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/authenticatorstaticstage"
	authenticatortotpstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/authenticatortotpstage"
	authenticatorvalidatestage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/authenticatorvalidatestage"
	authenticatorwebauthnstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/authenticatorwebauthnstage"
	captchastage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/captchastage"
	consentstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/consentstage"
	denystage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/denystage"
	dummystage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/dummystage"
	emailstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/emailstage"
	endpointsstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/endpointsstage"
	identificationstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/identificationstage"
	invitationstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/invitationstage"
	mutualtlsstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/mutualtlsstage"
	passwordstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/passwordstage"
	promptfield "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/promptfield"
	promptstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/promptstage"
	redirectstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/redirectstage"
	sourcestage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/sourcestage"
	userdeletestage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/userdeletestage"
	userloginstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/userloginstage"
	userlogoutstage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/userlogoutstage"
	userwritestage "github.com/l-nmch/provider-authentik/internal/controller/cluster/stage/userwritestage"
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
