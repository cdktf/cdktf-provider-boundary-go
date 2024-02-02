// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credentiallibraryvault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CredentialLibraryVaultConfig struct {
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
	// The ID of the credential store that this library belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/credential_library_vault#credential_store_id CredentialLibraryVault#credential_store_id}
	CredentialStoreId *string `field:"required" json:"credentialStoreId" yaml:"credentialStoreId"`
	// The path in Vault to request credentials from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/credential_library_vault#path CredentialLibraryVault#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The credential mapping override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/credential_library_vault#credential_mapping_overrides CredentialLibraryVault#credential_mapping_overrides}
	CredentialMappingOverrides *map[string]*string `field:"optional" json:"credentialMappingOverrides" yaml:"credentialMappingOverrides"`
	// The type of credential the library generates. Cannot be updated on an existing resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/credential_library_vault#credential_type CredentialLibraryVault#credential_type}
	CredentialType *string `field:"optional" json:"credentialType" yaml:"credentialType"`
	// The Vault credential library description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/credential_library_vault#description CredentialLibraryVault#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/credential_library_vault#http_method CredentialLibraryVault#http_method}
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The body of the HTTP request the library sends to Vault when requesting credentials.
	//
	// Only valid if `http_method` is set to `POST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/credential_library_vault#http_request_body CredentialLibraryVault#http_request_body}
	HttpRequestBody *string `field:"optional" json:"httpRequestBody" yaml:"httpRequestBody"`
	// The Vault credential library name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/credential_library_vault#name CredentialLibraryVault#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

