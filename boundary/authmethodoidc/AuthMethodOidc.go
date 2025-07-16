// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authmethodoidc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-boundary-go/boundary/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-boundary-go/boundary/v10/authmethodoidc/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/auth_method_oidc boundary_auth_method_oidc}.
type AuthMethodOidc interface {
	cdktf.TerraformResource
	AccountClaimMaps() *[]*string
	SetAccountClaimMaps(val *[]*string)
	AccountClaimMapsInput() *[]*string
	AllowedAudiences() *[]*string
	SetAllowedAudiences(val *[]*string)
	AllowedAudiencesInput() *[]*string
	ApiUrlPrefix() *string
	SetApiUrlPrefix(val *string)
	ApiUrlPrefixInput() *string
	CallbackUrl() *string
	SetCallbackUrl(val *string)
	CallbackUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClaimsScopes() *[]*string
	SetClaimsScopes(val *[]*string)
	ClaimsScopesInput() *[]*string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretHmac() *string
	SetClientSecretHmac(val *string)
	ClientSecretHmacInput() *string
	ClientSecretInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableDiscoveredConfigValidation() interface{}
	SetDisableDiscoveredConfigValidation(val interface{})
	DisableDiscoveredConfigValidationInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IdpCaCerts() *[]*string
	SetIdpCaCerts(val *[]*string)
	IdpCaCertsInput() *[]*string
	IsPrimaryForScope() interface{}
	SetIsPrimaryForScope(val interface{})
	IsPrimaryForScopeInput() interface{}
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Prompts() *[]*string
	SetPrompts(val *[]*string)
	PromptsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ScopeId() *string
	SetScopeId(val *string)
	ScopeIdInput() *string
	SigningAlgorithms() *[]*string
	SetSigningAlgorithms(val *[]*string)
	SigningAlgorithmsInput() *[]*string
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccountClaimMaps()
	ResetAllowedAudiences()
	ResetApiUrlPrefix()
	ResetCallbackUrl()
	ResetClaimsScopes()
	ResetClientId()
	ResetClientSecret()
	ResetClientSecretHmac()
	ResetDescription()
	ResetDisableDiscoveredConfigValidation()
	ResetIdpCaCerts()
	ResetIsPrimaryForScope()
	ResetIssuer()
	ResetMaxAge()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrompts()
	ResetSigningAlgorithms()
	ResetState()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AuthMethodOidc
type jsiiProxy_AuthMethodOidc struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AuthMethodOidc) AccountClaimMaps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountClaimMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) AccountClaimMapsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountClaimMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) AllowedAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) AllowedAudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAudiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ApiUrlPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ApiUrlPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) CallbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) CallbackUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ClaimsScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"claimsScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ClaimsScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"claimsScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ClientSecretHmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretHmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ClientSecretHmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretHmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) DisableDiscoveredConfigValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDiscoveredConfigValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) DisableDiscoveredConfigValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDiscoveredConfigValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) IdpCaCerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idpCaCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) IdpCaCertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idpCaCertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) IsPrimaryForScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryForScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) IsPrimaryForScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryForScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Prompts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prompts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) PromptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"promptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) ScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) SigningAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) SigningAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"signingAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodOidc) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/auth_method_oidc boundary_auth_method_oidc} Resource.
func NewAuthMethodOidc(scope constructs.Construct, id *string, config *AuthMethodOidcConfig) AuthMethodOidc {
	_init_.Initialize()

	if err := validateNewAuthMethodOidcParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthMethodOidc{}

	_jsii_.Create(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidc",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/auth_method_oidc boundary_auth_method_oidc} Resource.
func NewAuthMethodOidc_Override(a AuthMethodOidc, scope constructs.Construct, id *string, config *AuthMethodOidcConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidc",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetAccountClaimMaps(val *[]*string) {
	if err := j.validateSetAccountClaimMapsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountClaimMaps",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetAllowedAudiences(val *[]*string) {
	if err := j.validateSetAllowedAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedAudiences",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetApiUrlPrefix(val *string) {
	if err := j.validateSetApiUrlPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiUrlPrefix",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetCallbackUrl(val *string) {
	if err := j.validateSetCallbackUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callbackUrl",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetClaimsScopes(val *[]*string) {
	if err := j.validateSetClaimsScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claimsScopes",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetClientSecretHmac(val *string) {
	if err := j.validateSetClientSecretHmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretHmac",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetDisableDiscoveredConfigValidation(val interface{}) {
	if err := j.validateSetDisableDiscoveredConfigValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDiscoveredConfigValidation",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetIdpCaCerts(val *[]*string) {
	if err := j.validateSetIdpCaCertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpCaCerts",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetIsPrimaryForScope(val interface{}) {
	if err := j.validateSetIsPrimaryForScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPrimaryForScope",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetPrompts(val *[]*string) {
	if err := j.validateSetPromptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prompts",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetScopeId(val *string) {
	if err := j.validateSetScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeId",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetSigningAlgorithms(val *[]*string) {
	if err := j.validateSetSigningAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_AuthMethodOidc)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a AuthMethodOidc resource upon running "cdktf plan <stack-name>".
func AuthMethodOidc_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAuthMethodOidc_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidc",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AuthMethodOidc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethodOidc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidc",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuthMethodOidc_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethodOidc_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidc",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuthMethodOidc_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethodOidc_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidc",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AuthMethodOidc_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-boundary.authMethodOidc.AuthMethodOidc",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AuthMethodOidc) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AuthMethodOidc) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AuthMethodOidc) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AuthMethodOidc) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AuthMethodOidc) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AuthMethodOidc) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AuthMethodOidc) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetAccountClaimMaps() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountClaimMaps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetAllowedAudiences() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedAudiences",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetApiUrlPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetApiUrlPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetCallbackUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetCallbackUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetClaimsScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetClaimsScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetClientSecretHmac() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretHmac",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetDisableDiscoveredConfigValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableDiscoveredConfigValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetIdpCaCerts() {
	_jsii_.InvokeVoid(
		a,
		"resetIdpCaCerts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetIsPrimaryForScope() {
	_jsii_.InvokeVoid(
		a,
		"resetIsPrimaryForScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetMaxAge() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetPrompts() {
	_jsii_.InvokeVoid(
		a,
		"resetPrompts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetSigningAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetSigningAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetState() {
	_jsii_.InvokeVoid(
		a,
		"resetState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodOidc) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodOidc) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

