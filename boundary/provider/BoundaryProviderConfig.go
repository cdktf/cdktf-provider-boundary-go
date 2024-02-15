// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type BoundaryProviderConfig struct {
	// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env var.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#addr BoundaryProvider#addr}
	Addr *string `field:"required" json:"addr" yaml:"addr"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#alias BoundaryProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#auth_method_id BoundaryProvider#auth_method_id}
	AuthMethodId *string `field:"optional" json:"authMethodId" yaml:"authMethodId"`
	// The auth method login name for password-style or ldap-style auth methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#auth_method_login_name BoundaryProvider#auth_method_login_name}
	AuthMethodLoginName *string `field:"optional" json:"authMethodLoginName" yaml:"authMethodLoginName"`
	// The auth method password for password-style or ldap-style auth methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#auth_method_password BoundaryProvider#auth_method_password}
	AuthMethodPassword *string `field:"optional" json:"authMethodPassword" yaml:"authMethodPassword"`
	// The auth method login name for password-style auth methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#password_auth_method_login_name BoundaryProvider#password_auth_method_login_name}
	PasswordAuthMethodLoginName *string `field:"optional" json:"passwordAuthMethodLoginName" yaml:"passwordAuthMethodLoginName"`
	// The auth method password for password-style auth methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#password_auth_method_password BoundaryProvider#password_auth_method_password}
	PasswordAuthMethodPassword *string `field:"optional" json:"passwordAuthMethodPassword" yaml:"passwordAuthMethodPassword"`
	// Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#plugin_execution_dir BoundaryProvider#plugin_execution_dir}
	PluginExecutionDir *string `field:"optional" json:"pluginExecutionDir" yaml:"pluginExecutionDir"`
	// Can be a heredoc string or a path on disk.
	//
	// If set, the string/file will be parsed as HCL and used with the recovery KMS mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#recovery_kms_hcl BoundaryProvider#recovery_kms_hcl}
	RecoveryKmsHcl *string `field:"optional" json:"recoveryKmsHcl" yaml:"recoveryKmsHcl"`
	// The scope ID for the default auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#scope_id BoundaryProvider#scope_id}
	ScopeId *string `field:"optional" json:"scopeId" yaml:"scopeId"`
	// When set to true, does not validate the Boundary API endpoint certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#tls_insecure BoundaryProvider#tls_insecure}
	TlsInsecure interface{} `field:"optional" json:"tlsInsecure" yaml:"tlsInsecure"`
	// The Boundary token to use, as a string or path on disk containing just the string.
	//
	// If set, the token read here will be used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.14/docs#token BoundaryProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

