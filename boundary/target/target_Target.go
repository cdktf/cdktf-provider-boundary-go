package target

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-boundary-go/boundary/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-boundary-go/boundary/target/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/boundary/r/target boundary_target}.
type Target interface {
	cdktf.TerraformResource
	BrokeredCredentialSourceIds() *[]*string
	SetBrokeredCredentialSourceIds(val *[]*string)
	BrokeredCredentialSourceIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultPort() *float64
	SetDefaultPort(val *float64)
	DefaultPortInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostSourceIds() *[]*string
	SetHostSourceIds(val *[]*string)
	HostSourceIdsInput() *[]*string
	Id() *string
	InjectedApplicationCredentialSourceIds() *[]*string
	SetInjectedApplicationCredentialSourceIds(val *[]*string)
	InjectedApplicationCredentialSourceIdsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	SessionConnectionLimit() *float64
	SetSessionConnectionLimit(val *float64)
	SessionConnectionLimitInput() *float64
	SessionMaxSeconds() *float64
	SetSessionMaxSeconds(val *float64)
	SessionMaxSecondsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WorkerFilter() *string
	SetWorkerFilter(val *string)
	WorkerFilterInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetBrokeredCredentialSourceIds()
	ResetDefaultPort()
	ResetDescription()
	ResetHostSourceIds()
	ResetInjectedApplicationCredentialSourceIds()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSessionConnectionLimit()
	ResetSessionMaxSeconds()
	ResetWorkerFilter()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Target
type jsiiProxy_Target struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Target) BrokeredCredentialSourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"brokeredCredentialSourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) BrokeredCredentialSourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"brokeredCredentialSourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) DefaultPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) DefaultPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) HostSourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostSourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) HostSourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostSourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) InjectedApplicationCredentialSourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"injectedApplicationCredentialSourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) InjectedApplicationCredentialSourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"injectedApplicationCredentialSourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) ScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) ScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) SessionConnectionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionConnectionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) SessionConnectionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionConnectionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) SessionMaxSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionMaxSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) SessionMaxSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionMaxSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) WorkerFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Target) WorkerFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerFilterInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/boundary/r/target boundary_target} Resource.
func NewTarget(scope constructs.Construct, id *string, config *TargetConfig) Target {
	_init_.Initialize()

	if err := validateNewTargetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Target{}

	_jsii_.Create(
		"@cdktf/provider-boundary.target.Target",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/boundary/r/target boundary_target} Resource.
func NewTarget_Override(t Target, scope constructs.Construct, id *string, config *TargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-boundary.target.Target",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_Target)SetBrokeredCredentialSourceIds(val *[]*string) {
	if err := j.validateSetBrokeredCredentialSourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brokeredCredentialSourceIds",
		val,
	)
}

func (j *jsiiProxy_Target)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Target)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Target)SetDefaultPort(val *float64) {
	if err := j.validateSetDefaultPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPort",
		val,
	)
}

func (j *jsiiProxy_Target)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Target)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Target)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Target)SetHostSourceIds(val *[]*string) {
	if err := j.validateSetHostSourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostSourceIds",
		val,
	)
}

func (j *jsiiProxy_Target)SetInjectedApplicationCredentialSourceIds(val *[]*string) {
	if err := j.validateSetInjectedApplicationCredentialSourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"injectedApplicationCredentialSourceIds",
		val,
	)
}

func (j *jsiiProxy_Target)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Target)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Target)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Target)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Target)SetScopeId(val *string) {
	if err := j.validateSetScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeId",
		val,
	)
}

func (j *jsiiProxy_Target)SetSessionConnectionLimit(val *float64) {
	if err := j.validateSetSessionConnectionLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionConnectionLimit",
		val,
	)
}

func (j *jsiiProxy_Target)SetSessionMaxSeconds(val *float64) {
	if err := j.validateSetSessionMaxSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionMaxSeconds",
		val,
	)
}

func (j *jsiiProxy_Target)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Target)SetWorkerFilter(val *string) {
	if err := j.validateSetWorkerFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerFilter",
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
func Target_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.target.Target",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Target_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-boundary.target.Target",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_Target) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_Target) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_Target) ResetBrokeredCredentialSourceIds() {
	_jsii_.InvokeVoid(
		t,
		"resetBrokeredCredentialSourceIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetDefaultPort() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetHostSourceIds() {
	_jsii_.InvokeVoid(
		t,
		"resetHostSourceIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetInjectedApplicationCredentialSourceIds() {
	_jsii_.InvokeVoid(
		t,
		"resetInjectedApplicationCredentialSourceIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetSessionConnectionLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetSessionConnectionLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetSessionMaxSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetSessionMaxSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) ResetWorkerFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkerFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Target) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Target) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

