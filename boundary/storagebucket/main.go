// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-boundary.storageBucket.StorageBucket",
		reflect.TypeOf((*StorageBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "attributesJson", GoGetter: "AttributesJson"},
			_jsii_.MemberProperty{JsiiProperty: "attributesJsonInput", GoGetter: "AttributesJsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketNameInput", GoGetter: "BucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefix", GoGetter: "BucketPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "bucketPrefixInput", GoGetter: "BucketPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "internalForceUpdate", GoGetter: "InternalForceUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "internalHmacUsedForSecretsConfigHmac", GoGetter: "InternalHmacUsedForSecretsConfigHmac"},
			_jsii_.MemberProperty{JsiiProperty: "internalSecretsConfigHmac", GoGetter: "InternalSecretsConfigHmac"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pluginId", GoGetter: "PluginId"},
			_jsii_.MemberProperty{JsiiProperty: "pluginIdInput", GoGetter: "PluginIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "pluginName", GoGetter: "PluginName"},
			_jsii_.MemberProperty{JsiiProperty: "pluginNameInput", GoGetter: "PluginNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttributesJson", GoMethod: "ResetAttributesJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucketPrefix", GoMethod: "ResetBucketPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPluginId", GoMethod: "ResetPluginId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPluginName", GoMethod: "ResetPluginName"},
			_jsii_.MemberProperty{JsiiProperty: "scopeId", GoGetter: "ScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "scopeIdInput", GoGetter: "ScopeIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "secretsHmac", GoGetter: "SecretsHmac"},
			_jsii_.MemberProperty{JsiiProperty: "secretsJson", GoGetter: "SecretsJson"},
			_jsii_.MemberProperty{JsiiProperty: "secretsJsonInput", GoGetter: "SecretsJsonInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "workerFilter", GoGetter: "WorkerFilter"},
			_jsii_.MemberProperty{JsiiProperty: "workerFilterInput", GoGetter: "WorkerFilterInput"},
		},
		func() interface{} {
			j := jsiiProxy_StorageBucket{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-boundary.storageBucket.StorageBucketConfig",
		reflect.TypeOf((*StorageBucketConfig)(nil)).Elem(),
	)
}
