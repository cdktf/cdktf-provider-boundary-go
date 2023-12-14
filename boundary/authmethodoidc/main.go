// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authmethodoidc

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidc",
		reflect.TypeOf((*AuthMethodOidc)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountClaimMaps", GoGetter: "AccountClaimMaps"},
			_jsii_.MemberProperty{JsiiProperty: "accountClaimMapsInput", GoGetter: "AccountClaimMapsInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedAudiences", GoGetter: "AllowedAudiences"},
			_jsii_.MemberProperty{JsiiProperty: "allowedAudiencesInput", GoGetter: "AllowedAudiencesInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrlPrefix", GoGetter: "ApiUrlPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrlPrefixInput", GoGetter: "ApiUrlPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrl", GoGetter: "CallbackUrl"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrlInput", GoGetter: "CallbackUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "claimsScopes", GoGetter: "ClaimsScopes"},
			_jsii_.MemberProperty{JsiiProperty: "claimsScopesInput", GoGetter: "ClaimsScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretHmac", GoGetter: "ClientSecretHmac"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretHmacInput", GoGetter: "ClientSecretHmacInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableDiscoveredConfigValidation", GoGetter: "DisableDiscoveredConfigValidation"},
			_jsii_.MemberProperty{JsiiProperty: "disableDiscoveredConfigValidationInput", GoGetter: "DisableDiscoveredConfigValidationInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idpCaCerts", GoGetter: "IdpCaCerts"},
			_jsii_.MemberProperty{JsiiProperty: "idpCaCertsInput", GoGetter: "IdpCaCertsInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isPrimaryForScope", GoGetter: "IsPrimaryForScope"},
			_jsii_.MemberProperty{JsiiProperty: "isPrimaryForScopeInput", GoGetter: "IsPrimaryForScopeInput"},
			_jsii_.MemberProperty{JsiiProperty: "issuer", GoGetter: "Issuer"},
			_jsii_.MemberProperty{JsiiProperty: "issuerInput", GoGetter: "IssuerInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxAge", GoGetter: "MaxAge"},
			_jsii_.MemberProperty{JsiiProperty: "maxAgeInput", GoGetter: "MaxAgeInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "prompts", GoGetter: "Prompts"},
			_jsii_.MemberProperty{JsiiProperty: "promptsInput", GoGetter: "PromptsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountClaimMaps", GoMethod: "ResetAccountClaimMaps"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedAudiences", GoMethod: "ResetAllowedAudiences"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiUrlPrefix", GoMethod: "ResetApiUrlPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetCallbackUrl", GoMethod: "ResetCallbackUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetClaimsScopes", GoMethod: "ResetClaimsScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientSecret", GoMethod: "ResetClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientSecretHmac", GoMethod: "ResetClientSecretHmac"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableDiscoveredConfigValidation", GoMethod: "ResetDisableDiscoveredConfigValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdpCaCerts", GoMethod: "ResetIdpCaCerts"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsPrimaryForScope", GoMethod: "ResetIsPrimaryForScope"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssuer", GoMethod: "ResetIssuer"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxAge", GoMethod: "ResetMaxAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrompts", GoMethod: "ResetPrompts"},
			_jsii_.MemberMethod{JsiiMethod: "resetSigningAlgorithms", GoMethod: "ResetSigningAlgorithms"},
			_jsii_.MemberMethod{JsiiMethod: "resetState", GoMethod: "ResetState"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberProperty{JsiiProperty: "scopeId", GoGetter: "ScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "scopeIdInput", GoGetter: "ScopeIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "signingAlgorithms", GoGetter: "SigningAlgorithms"},
			_jsii_.MemberProperty{JsiiProperty: "signingAlgorithmsInput", GoGetter: "SigningAlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "stateInput", GoGetter: "StateInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AuthMethodOidc{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidcConfig",
		reflect.TypeOf((*AuthMethodOidcConfig)(nil)).Elem(),
	)
}
