// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authmethodldap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-boundary-go/boundary/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-boundary-go/boundary/v8/authmethodldap/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_ldap boundary_auth_method_ldap}.
type AuthMethodLdap interface {
	cdktf.TerraformResource
	AccountAttributeMaps() *[]*string
	SetAccountAttributeMaps(val *[]*string)
	AccountAttributeMapsInput() *[]*string
	AnonGroupSearch() interface{}
	SetAnonGroupSearch(val interface{})
	AnonGroupSearchInput() interface{}
	BindDn() *string
	SetBindDn(val *string)
	BindDnInput() *string
	BindPassword() *string
	SetBindPassword(val *string)
	BindPasswordHmac() *string
	SetBindPasswordHmac(val *string)
	BindPasswordHmacInput() *string
	BindPasswordInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificates() *[]*string
	SetCertificates(val *[]*string)
	CertificatesInput() *[]*string
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientCertificateKey() *string
	SetClientCertificateKey(val *string)
	ClientCertificateKeyHmac() *string
	SetClientCertificateKeyHmac(val *string)
	ClientCertificateKeyHmacInput() *string
	ClientCertificateKeyInput() *string
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
	DereferenceAliases() *string
	SetDereferenceAliases(val *string)
	DereferenceAliasesInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiscoverDn() interface{}
	SetDiscoverDn(val interface{})
	DiscoverDnInput() interface{}
	EnableGroups() interface{}
	SetEnableGroups(val interface{})
	EnableGroupsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupAttr() *string
	SetGroupAttr(val *string)
	GroupAttrInput() *string
	GroupDn() *string
	SetGroupDn(val *string)
	GroupDnInput() *string
	GroupFilter() *string
	SetGroupFilter(val *string)
	GroupFilterInput() *string
	Id() *string
	InsecureTls() interface{}
	SetInsecureTls(val interface{})
	InsecureTlsInput() interface{}
	IsPrimaryForScope() interface{}
	SetIsPrimaryForScope(val interface{})
	IsPrimaryForScopeInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumPageSize() *float64
	SetMaximumPageSize(val *float64)
	MaximumPageSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	StartTls() interface{}
	SetStartTls(val interface{})
	StartTlsInput() interface{}
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
	UpnDomain() *string
	SetUpnDomain(val *string)
	UpnDomainInput() *string
	Urls() *[]*string
	SetUrls(val *[]*string)
	UrlsInput() *[]*string
	UserAttr() *string
	SetUserAttr(val *string)
	UserAttrInput() *string
	UserDn() *string
	SetUserDn(val *string)
	UserDnInput() *string
	UserFilter() *string
	SetUserFilter(val *string)
	UserFilterInput() *string
	UseTokenGroups() interface{}
	SetUseTokenGroups(val interface{})
	UseTokenGroupsInput() interface{}
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
	ResetAccountAttributeMaps()
	ResetAnonGroupSearch()
	ResetBindDn()
	ResetBindPassword()
	ResetBindPasswordHmac()
	ResetCertificates()
	ResetClientCertificate()
	ResetClientCertificateKey()
	ResetClientCertificateKeyHmac()
	ResetDereferenceAliases()
	ResetDescription()
	ResetDiscoverDn()
	ResetEnableGroups()
	ResetGroupAttr()
	ResetGroupDn()
	ResetGroupFilter()
	ResetInsecureTls()
	ResetIsPrimaryForScope()
	ResetMaximumPageSize()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartTls()
	ResetState()
	ResetType()
	ResetUpnDomain()
	ResetUrls()
	ResetUserAttr()
	ResetUserDn()
	ResetUserFilter()
	ResetUseTokenGroups()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AuthMethodLdap
type jsiiProxy_AuthMethodLdap struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AuthMethodLdap) AccountAttributeMaps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountAttributeMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) AccountAttributeMapsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountAttributeMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) AnonGroupSearch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonGroupSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) AnonGroupSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonGroupSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) BindDn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindDn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) BindDnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindDnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) BindPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) BindPasswordHmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindPasswordHmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) BindPasswordHmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindPasswordHmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) BindPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Certificates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) CertificatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ClientCertificateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ClientCertificateKeyHmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateKeyHmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ClientCertificateKeyHmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateKeyHmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ClientCertificateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) DereferenceAliases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dereferenceAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) DereferenceAliasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dereferenceAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) DiscoverDn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discoverDn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) DiscoverDnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discoverDnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) EnableGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) EnableGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) GroupAttr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupAttr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) GroupAttrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupAttrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) GroupDn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupDn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) GroupDnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupDnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) GroupFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) GroupFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) InsecureTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) InsecureTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) IsPrimaryForScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryForScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) IsPrimaryForScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryForScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) MaximumPageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) MaximumPageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) ScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) StartTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) StartTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UpnDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upnDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UpnDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upnDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) Urls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UserAttr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAttr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UserAttrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAttrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UserDn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UserDnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UserFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UserFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UseTokenGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTokenGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthMethodLdap) UseTokenGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTokenGroupsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_ldap boundary_auth_method_ldap} Resource.
func NewAuthMethodLdap(scope constructs.Construct, id *string, config *AuthMethodLdapConfig) AuthMethodLdap {
	_init_.Initialize()

	if err := validateNewAuthMethodLdapParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthMethodLdap{}

	_jsii_.Create(
		"@cdktf/provider-boundary.authMethodLdap.AuthMethodLdap",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/auth_method_ldap boundary_auth_method_ldap} Resource.
func NewAuthMethodLdap_Override(a AuthMethodLdap, scope constructs.Construct, id *string, config *AuthMethodLdapConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-boundary.authMethodLdap.AuthMethodLdap",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetAccountAttributeMaps(val *[]*string) {
	if err := j.validateSetAccountAttributeMapsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountAttributeMaps",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetAnonGroupSearch(val interface{}) {
	if err := j.validateSetAnonGroupSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonGroupSearch",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetBindDn(val *string) {
	if err := j.validateSetBindDnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindDn",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetBindPassword(val *string) {
	if err := j.validateSetBindPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindPassword",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetBindPasswordHmac(val *string) {
	if err := j.validateSetBindPasswordHmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindPasswordHmac",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetCertificates(val *[]*string) {
	if err := j.validateSetCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificates",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetClientCertificate(val *string) {
	if err := j.validateSetClientCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetClientCertificateKey(val *string) {
	if err := j.validateSetClientCertificateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateKey",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetClientCertificateKeyHmac(val *string) {
	if err := j.validateSetClientCertificateKeyHmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateKeyHmac",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetDereferenceAliases(val *string) {
	if err := j.validateSetDereferenceAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dereferenceAliases",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetDiscoverDn(val interface{}) {
	if err := j.validateSetDiscoverDnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoverDn",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetEnableGroups(val interface{}) {
	if err := j.validateSetEnableGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGroups",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetGroupAttr(val *string) {
	if err := j.validateSetGroupAttrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupAttr",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetGroupDn(val *string) {
	if err := j.validateSetGroupDnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupDn",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetGroupFilter(val *string) {
	if err := j.validateSetGroupFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupFilter",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetInsecureTls(val interface{}) {
	if err := j.validateSetInsecureTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureTls",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetIsPrimaryForScope(val interface{}) {
	if err := j.validateSetIsPrimaryForScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPrimaryForScope",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetMaximumPageSize(val *float64) {
	if err := j.validateSetMaximumPageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumPageSize",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetScopeId(val *string) {
	if err := j.validateSetScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeId",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetStartTls(val interface{}) {
	if err := j.validateSetStartTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTls",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetUpnDomain(val *string) {
	if err := j.validateSetUpnDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upnDomain",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetUrls(val *[]*string) {
	if err := j.validateSetUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urls",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetUserAttr(val *string) {
	if err := j.validateSetUserAttrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAttr",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetUserDn(val *string) {
	if err := j.validateSetUserDnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDn",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetUserFilter(val *string) {
	if err := j.validateSetUserFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userFilter",
		val,
	)
}

func (j *jsiiProxy_AuthMethodLdap)SetUseTokenGroups(val interface{}) {
	if err := j.validateSetUseTokenGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTokenGroups",
		val,
	)
}

// Generates CDKTF code for importing a AuthMethodLdap resource upon running "cdktf plan <stack-name>".
func AuthMethodLdap_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAuthMethodLdap_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.authMethodLdap.AuthMethodLdap",
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
func AuthMethodLdap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethodLdap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.authMethodLdap.AuthMethodLdap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuthMethodLdap_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethodLdap_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.authMethodLdap.AuthMethodLdap",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuthMethodLdap_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthMethodLdap_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.authMethodLdap.AuthMethodLdap",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AuthMethodLdap_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-boundary.authMethodLdap.AuthMethodLdap",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AuthMethodLdap) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AuthMethodLdap) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AuthMethodLdap) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AuthMethodLdap) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AuthMethodLdap) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AuthMethodLdap) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AuthMethodLdap) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AuthMethodLdap) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AuthMethodLdap) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AuthMethodLdap) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AuthMethodLdap) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AuthMethodLdap) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodLdap) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AuthMethodLdap) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AuthMethodLdap) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AuthMethodLdap) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AuthMethodLdap) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AuthMethodLdap) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetAccountAttributeMaps() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountAttributeMaps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetAnonGroupSearch() {
	_jsii_.InvokeVoid(
		a,
		"resetAnonGroupSearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetBindDn() {
	_jsii_.InvokeVoid(
		a,
		"resetBindDn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetBindPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetBindPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetBindPasswordHmac() {
	_jsii_.InvokeVoid(
		a,
		"resetBindPasswordHmac",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetCertificates() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetClientCertificateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetClientCertificateKeyHmac() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateKeyHmac",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetDereferenceAliases() {
	_jsii_.InvokeVoid(
		a,
		"resetDereferenceAliases",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetDiscoverDn() {
	_jsii_.InvokeVoid(
		a,
		"resetDiscoverDn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetEnableGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetGroupAttr() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupAttr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetGroupDn() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupDn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetGroupFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetInsecureTls() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecureTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetIsPrimaryForScope() {
	_jsii_.InvokeVoid(
		a,
		"resetIsPrimaryForScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetMaximumPageSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumPageSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetStartTls() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetState() {
	_jsii_.InvokeVoid(
		a,
		"resetState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetUpnDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetUpnDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetUserAttr() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAttr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetUserDn() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetUserFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetUserFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) ResetUseTokenGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetUseTokenGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthMethodLdap) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodLdap) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodLdap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthMethodLdap) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

