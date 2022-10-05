package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-boundary-go/boundary/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-boundary-go/boundary/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/boundary boundary}.
type BoundaryProvider interface {
	cdktf.TerraformProvider
	Addr() *string
	SetAddr(val *string)
	AddrInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AuthMethodId() *string
	SetAuthMethodId(val *string)
	AuthMethodIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	PasswordAuthMethodLoginName() *string
	SetPasswordAuthMethodLoginName(val *string)
	PasswordAuthMethodLoginNameInput() *string
	PasswordAuthMethodPassword() *string
	SetPasswordAuthMethodPassword(val *string)
	PasswordAuthMethodPasswordInput() *string
	PluginExecutionDir() *string
	SetPluginExecutionDir(val *string)
	PluginExecutionDirInput() *string
	// Experimental.
	RawOverrides() interface{}
	RecoveryKmsHcl() *string
	SetRecoveryKmsHcl(val *string)
	RecoveryKmsHclInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TlsInsecure() interface{}
	SetTlsInsecure(val interface{})
	TlsInsecureInput() interface{}
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetAuthMethodId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordAuthMethodLoginName()
	ResetPasswordAuthMethodPassword()
	ResetPluginExecutionDir()
	ResetRecoveryKmsHcl()
	ResetTlsInsecure()
	ResetToken()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BoundaryProvider
type jsiiProxy_BoundaryProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_BoundaryProvider) Addr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) AddrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) AuthMethodId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) AuthMethodIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) PasswordAuthMethodLoginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordAuthMethodLoginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) PasswordAuthMethodLoginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordAuthMethodLoginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) PasswordAuthMethodPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordAuthMethodPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) PasswordAuthMethodPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordAuthMethodPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) PluginExecutionDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginExecutionDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) PluginExecutionDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginExecutionDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) RecoveryKmsHcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryKmsHcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) RecoveryKmsHclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryKmsHclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) TlsInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) TlsInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BoundaryProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/boundary boundary} Resource.
func NewBoundaryProvider(scope constructs.Construct, id *string, config *BoundaryProviderConfig) BoundaryProvider {
	_init_.Initialize()

	if err := validateNewBoundaryProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BoundaryProvider{}

	_jsii_.Create(
		"@cdktf/provider-boundary.provider.BoundaryProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/boundary boundary} Resource.
func NewBoundaryProvider_Override(b BoundaryProvider, scope constructs.Construct, id *string, config *BoundaryProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-boundary.provider.BoundaryProvider",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetAddr(val *string) {
	_jsii_.Set(
		j,
		"addr",
		val,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetAuthMethodId(val *string) {
	_jsii_.Set(
		j,
		"authMethodId",
		val,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetPasswordAuthMethodLoginName(val *string) {
	_jsii_.Set(
		j,
		"passwordAuthMethodLoginName",
		val,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetPasswordAuthMethodPassword(val *string) {
	_jsii_.Set(
		j,
		"passwordAuthMethodPassword",
		val,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetPluginExecutionDir(val *string) {
	_jsii_.Set(
		j,
		"pluginExecutionDir",
		val,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetRecoveryKmsHcl(val *string) {
	_jsii_.Set(
		j,
		"recoveryKmsHcl",
		val,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetTlsInsecure(val interface{}) {
	if err := j.validateSetTlsInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsInsecure",
		val,
	)
}

func (j *jsiiProxy_BoundaryProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
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
func BoundaryProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBoundaryProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.provider.BoundaryProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BoundaryProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-boundary.provider.BoundaryProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BoundaryProvider) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BoundaryProvider) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		b,
		"resetAlias",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetAuthMethodId() {
	_jsii_.InvokeVoid(
		b,
		"resetAuthMethodId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetPasswordAuthMethodLoginName() {
	_jsii_.InvokeVoid(
		b,
		"resetPasswordAuthMethodLoginName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetPasswordAuthMethodPassword() {
	_jsii_.InvokeVoid(
		b,
		"resetPasswordAuthMethodPassword",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetPluginExecutionDir() {
	_jsii_.InvokeVoid(
		b,
		"resetPluginExecutionDir",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetRecoveryKmsHcl() {
	_jsii_.InvokeVoid(
		b,
		"resetRecoveryKmsHcl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetTlsInsecure() {
	_jsii_.InvokeVoid(
		b,
		"resetTlsInsecure",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) ResetToken() {
	_jsii_.InvokeVoid(
		b,
		"resetToken",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BoundaryProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BoundaryProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BoundaryProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BoundaryProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

