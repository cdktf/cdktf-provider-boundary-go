package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-boundary.provider.BoundaryProvider",
		reflect.TypeOf((*BoundaryProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "addr", GoGetter: "Addr"},
			_jsii_.MemberProperty{JsiiProperty: "addrInput", GoGetter: "AddrInput"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "authMethodId", GoGetter: "AuthMethodId"},
			_jsii_.MemberProperty{JsiiProperty: "authMethodIdInput", GoGetter: "AuthMethodIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordAuthMethodLoginName", GoGetter: "PasswordAuthMethodLoginName"},
			_jsii_.MemberProperty{JsiiProperty: "passwordAuthMethodLoginNameInput", GoGetter: "PasswordAuthMethodLoginNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordAuthMethodPassword", GoGetter: "PasswordAuthMethodPassword"},
			_jsii_.MemberProperty{JsiiProperty: "passwordAuthMethodPasswordInput", GoGetter: "PasswordAuthMethodPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "pluginExecutionDir", GoGetter: "PluginExecutionDir"},
			_jsii_.MemberProperty{JsiiProperty: "pluginExecutionDirInput", GoGetter: "PluginExecutionDirInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryKmsHcl", GoGetter: "RecoveryKmsHcl"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryKmsHclInput", GoGetter: "RecoveryKmsHclInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthMethodId", GoMethod: "ResetAuthMethodId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordAuthMethodLoginName", GoMethod: "ResetPasswordAuthMethodLoginName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordAuthMethodPassword", GoMethod: "ResetPasswordAuthMethodPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetPluginExecutionDir", GoMethod: "ResetPluginExecutionDir"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecoveryKmsHcl", GoMethod: "ResetRecoveryKmsHcl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsInsecure", GoMethod: "ResetTlsInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsInsecure", GoGetter: "TlsInsecure"},
			_jsii_.MemberProperty{JsiiProperty: "tlsInsecureInput", GoGetter: "TlsInsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_BoundaryProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-boundary.provider.BoundaryProviderConfig",
		reflect.TypeOf((*BoundaryProviderConfig)(nil)).Elem(),
	)
}
