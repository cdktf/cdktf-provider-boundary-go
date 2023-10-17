// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package target

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-boundary.target.Target",
		reflect.TypeOf((*Target)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "address", GoGetter: "Address"},
			_jsii_.MemberProperty{JsiiProperty: "addressInput", GoGetter: "AddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "brokeredCredentialSourceIds", GoGetter: "BrokeredCredentialSourceIds"},
			_jsii_.MemberProperty{JsiiProperty: "brokeredCredentialSourceIdsInput", GoGetter: "BrokeredCredentialSourceIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultClientPort", GoGetter: "DefaultClientPort"},
			_jsii_.MemberProperty{JsiiProperty: "defaultClientPortInput", GoGetter: "DefaultClientPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPort", GoGetter: "DefaultPort"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPortInput", GoGetter: "DefaultPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "egressWorkerFilter", GoGetter: "EgressWorkerFilter"},
			_jsii_.MemberProperty{JsiiProperty: "egressWorkerFilterInput", GoGetter: "EgressWorkerFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableSessionRecording", GoGetter: "EnableSessionRecording"},
			_jsii_.MemberProperty{JsiiProperty: "enableSessionRecordingInput", GoGetter: "EnableSessionRecordingInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostSourceIds", GoGetter: "HostSourceIds"},
			_jsii_.MemberProperty{JsiiProperty: "hostSourceIdsInput", GoGetter: "HostSourceIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "ingressWorkerFilter", GoGetter: "IngressWorkerFilter"},
			_jsii_.MemberProperty{JsiiProperty: "ingressWorkerFilterInput", GoGetter: "IngressWorkerFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "injectedApplicationCredentialSourceIds", GoGetter: "InjectedApplicationCredentialSourceIds"},
			_jsii_.MemberProperty{JsiiProperty: "injectedApplicationCredentialSourceIdsInput", GoGetter: "InjectedApplicationCredentialSourceIdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAddress", GoMethod: "ResetAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetBrokeredCredentialSourceIds", GoMethod: "ResetBrokeredCredentialSourceIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultClientPort", GoMethod: "ResetDefaultClientPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultPort", GoMethod: "ResetDefaultPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEgressWorkerFilter", GoMethod: "ResetEgressWorkerFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableSessionRecording", GoMethod: "ResetEnableSessionRecording"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostSourceIds", GoMethod: "ResetHostSourceIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetIngressWorkerFilter", GoMethod: "ResetIngressWorkerFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetInjectedApplicationCredentialSourceIds", GoMethod: "ResetInjectedApplicationCredentialSourceIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionConnectionLimit", GoMethod: "ResetSessionConnectionLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionMaxSeconds", GoMethod: "ResetSessionMaxSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageBucketId", GoMethod: "ResetStorageBucketId"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkerFilter", GoMethod: "ResetWorkerFilter"},
			_jsii_.MemberProperty{JsiiProperty: "scopeId", GoGetter: "ScopeId"},
			_jsii_.MemberProperty{JsiiProperty: "scopeIdInput", GoGetter: "ScopeIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionConnectionLimit", GoGetter: "SessionConnectionLimit"},
			_jsii_.MemberProperty{JsiiProperty: "sessionConnectionLimitInput", GoGetter: "SessionConnectionLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionMaxSeconds", GoGetter: "SessionMaxSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "sessionMaxSecondsInput", GoGetter: "SessionMaxSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageBucketId", GoGetter: "StorageBucketId"},
			_jsii_.MemberProperty{JsiiProperty: "storageBucketIdInput", GoGetter: "StorageBucketIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "workerFilter", GoGetter: "WorkerFilter"},
			_jsii_.MemberProperty{JsiiProperty: "workerFilterInput", GoGetter: "WorkerFilterInput"},
		},
		func() interface{} {
			j := jsiiProxy_Target{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-boundary.target.TargetConfig",
		reflect.TypeOf((*TargetConfig)(nil)).Elem(),
	)
}
