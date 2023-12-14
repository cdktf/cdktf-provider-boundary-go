// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authmethodoidc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthMethodOidcConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The scope ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#scope_id AuthMethodOidc#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// Account claim maps for the to_claim of sub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#account_claim_maps AuthMethodOidc#account_claim_maps}
	AccountClaimMaps *[]*string `field:"optional" json:"accountClaimMaps" yaml:"accountClaimMaps"`
	// Audiences for which the provider responses will be allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#allowed_audiences AuthMethodOidc#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// The API prefix to use when generating callback URLs for the provider.
	//
	// Should be set to an address at which the provider can reach back to the controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#api_url_prefix AuthMethodOidc#api_url_prefix}
	ApiUrlPrefix *string `field:"optional" json:"apiUrlPrefix" yaml:"apiUrlPrefix"`
	// The URL that should be provided to the IdP for callbacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#callback_url AuthMethodOidc#callback_url}
	CallbackUrl *string `field:"optional" json:"callbackUrl" yaml:"callbackUrl"`
	// Claims scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#claims_scopes AuthMethodOidc#claims_scopes}
	ClaimsScopes *[]*string `field:"optional" json:"claimsScopes" yaml:"claimsScopes"`
	// The client ID assigned to this auth method from the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#client_id AuthMethodOidc#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The secret key assigned to this auth method from the provider.
	//
	// Once set, only the hash will be kept and the original value can be removed from configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#client_secret AuthMethodOidc#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The HMAC of the client secret returned by the Boundary controller, which is used for comparison after initial setting of the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#client_secret_hmac AuthMethodOidc#client_secret_hmac}
	ClientSecretHmac *string `field:"optional" json:"clientSecretHmac" yaml:"clientSecretHmac"`
	// The auth method description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#description AuthMethodOidc#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disables validation logic ensuring that the OIDC provider's information from its discovery endpoint matches the information here.
	//
	// The validation is only performed at create or update time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#disable_discovered_config_validation AuthMethodOidc#disable_discovered_config_validation}
	DisableDiscoveredConfigValidation interface{} `field:"optional" json:"disableDiscoveredConfigValidation" yaml:"disableDiscoveredConfigValidation"`
	// A list of CA certificates to trust when validating the IdP's token signatures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#idp_ca_certs AuthMethodOidc#idp_ca_certs}
	IdpCaCerts *[]*string `field:"optional" json:"idpCaCerts" yaml:"idpCaCerts"`
	// When true, makes this auth method the primary auth method for the scope in which it resides.
	//
	// The primary auth method for a scope means the user will be automatically created when they login using an OIDC account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#is_primary_for_scope AuthMethodOidc#is_primary_for_scope}
	IsPrimaryForScope interface{} `field:"optional" json:"isPrimaryForScope" yaml:"isPrimaryForScope"`
	// The issuer corresponding to the provider, which must match the issuer field in generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#issuer AuthMethodOidc#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The max age to provide to the provider, indicating how much time is allowed to have passed since the last authentication before the user is challenged again.
	//
	// A value of 0 sets an immediate requirement for all users to reauthenticate, and an unset maxAge results in a Terraform value of -1 and the default TTL of the chosen OIDC will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#max_age AuthMethodOidc#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// The auth method name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#name AuthMethodOidc#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The prompts passed to the identity provider to determine whether to prompt the end-user for reauthentication, account selection or consent.
	//
	// Please note the values passed are case-sensitive. The valid values are: `none`, `login`, `consent` and `select_account`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#prompts AuthMethodOidc#prompts}
	Prompts *[]*string `field:"optional" json:"prompts" yaml:"prompts"`
	// Allowed signing algorithms for the provider's issued tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#signing_algorithms AuthMethodOidc#signing_algorithms}
	SigningAlgorithms *[]*string `field:"optional" json:"signingAlgorithms" yaml:"signingAlgorithms"`
	// Can be one of 'inactive', 'active-private', or 'active-public'. Currently automatically set to active-public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#state AuthMethodOidc#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The type of auth method; hardcoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_oidc#type AuthMethodOidc#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

