package hostcatalogplugin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-boundary-go/boundary/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-boundary-go/boundary/hostcatalogplugin/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/boundary/r/host_catalog_plugin boundary_host_catalog_plugin}.
type HostCatalogPlugin interface {
	cdktf.TerraformResource
	AttributesJson() *string
	SetAttributesJson(val *string)
	AttributesJsonInput() *string
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
	Id() *string
	InternalForceUpdate() *string
	SetInternalForceUpdate(val *string)
	InternalForceUpdateInput() *string
	InternalHmacUsedForSecretsConfigHmac() *string
	SetInternalHmacUsedForSecretsConfigHmac(val *string)
	InternalHmacUsedForSecretsConfigHmacInput() *string
	InternalSecretsConfigHmac() *string
	SetInternalSecretsConfigHmac(val *string)
	InternalSecretsConfigHmacInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PluginId() *string
	SetPluginId(val *string)
	PluginIdInput() *string
	PluginName() *string
	SetPluginName(val *string)
	PluginNameInput() *string
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
	SecretsHmac() *string
	SetSecretsHmac(val *string)
	SecretsHmacInput() *string
	SecretsJson() *string
	SetSecretsJson(val *string)
	SecretsJsonInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAttributesJson()
	ResetDescription()
	ResetInternalForceUpdate()
	ResetInternalHmacUsedForSecretsConfigHmac()
	ResetInternalSecretsConfigHmac()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPluginId()
	ResetPluginName()
	ResetSecretsHmac()
	ResetSecretsJson()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for HostCatalogPlugin
type jsiiProxy_HostCatalogPlugin struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HostCatalogPlugin) AttributesJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) AttributesJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) InternalForceUpdate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalForceUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) InternalForceUpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalForceUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) InternalHmacUsedForSecretsConfigHmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalHmacUsedForSecretsConfigHmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) InternalHmacUsedForSecretsConfigHmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalHmacUsedForSecretsConfigHmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) InternalSecretsConfigHmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalSecretsConfigHmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) InternalSecretsConfigHmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalSecretsConfigHmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) PluginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) PluginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) ScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) ScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) SecretsHmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsHmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) SecretsHmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsHmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) SecretsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) SecretsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostCatalogPlugin) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/boundary/r/host_catalog_plugin boundary_host_catalog_plugin} Resource.
func NewHostCatalogPlugin(scope constructs.Construct, id *string, config *HostCatalogPluginConfig) HostCatalogPlugin {
	_init_.Initialize()

	if err := validateNewHostCatalogPluginParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HostCatalogPlugin{}

	_jsii_.Create(
		"@cdktf/provider-boundary.hostCatalogPlugin.HostCatalogPlugin",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/boundary/r/host_catalog_plugin boundary_host_catalog_plugin} Resource.
func NewHostCatalogPlugin_Override(h HostCatalogPlugin, scope constructs.Construct, id *string, config *HostCatalogPluginConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-boundary.hostCatalogPlugin.HostCatalogPlugin",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetAttributesJson(val *string) {
	if err := j.validateSetAttributesJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributesJson",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetInternalForceUpdate(val *string) {
	if err := j.validateSetInternalForceUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalForceUpdate",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetInternalHmacUsedForSecretsConfigHmac(val *string) {
	if err := j.validateSetInternalHmacUsedForSecretsConfigHmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalHmacUsedForSecretsConfigHmac",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetInternalSecretsConfigHmac(val *string) {
	if err := j.validateSetInternalSecretsConfigHmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalSecretsConfigHmac",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetPluginId(val *string) {
	if err := j.validateSetPluginIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginId",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetScopeId(val *string) {
	if err := j.validateSetScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeId",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetSecretsHmac(val *string) {
	if err := j.validateSetSecretsHmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsHmac",
		val,
	)
}

func (j *jsiiProxy_HostCatalogPlugin)SetSecretsJson(val *string) {
	if err := j.validateSetSecretsJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsJson",
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
func HostCatalogPlugin_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHostCatalogPlugin_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-boundary.hostCatalogPlugin.HostCatalogPlugin",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HostCatalogPlugin_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-boundary.hostCatalogPlugin.HostCatalogPlugin",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HostCatalogPlugin) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetAttributesJson() {
	_jsii_.InvokeVoid(
		h,
		"resetAttributesJson",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetDescription() {
	_jsii_.InvokeVoid(
		h,
		"resetDescription",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetInternalForceUpdate() {
	_jsii_.InvokeVoid(
		h,
		"resetInternalForceUpdate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetInternalHmacUsedForSecretsConfigHmac() {
	_jsii_.InvokeVoid(
		h,
		"resetInternalHmacUsedForSecretsConfigHmac",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetInternalSecretsConfigHmac() {
	_jsii_.InvokeVoid(
		h,
		"resetInternalSecretsConfigHmac",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetName() {
	_jsii_.InvokeVoid(
		h,
		"resetName",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetPluginId() {
	_jsii_.InvokeVoid(
		h,
		"resetPluginId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetPluginName() {
	_jsii_.InvokeVoid(
		h,
		"resetPluginName",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetSecretsHmac() {
	_jsii_.InvokeVoid(
		h,
		"resetSecretsHmac",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) ResetSecretsJson() {
	_jsii_.InvokeVoid(
		h,
		"resetSecretsJson",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HostCatalogPlugin) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostCatalogPlugin) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

