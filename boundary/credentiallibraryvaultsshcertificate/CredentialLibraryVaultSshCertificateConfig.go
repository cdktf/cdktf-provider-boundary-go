// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credentiallibraryvaultsshcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CredentialLibraryVaultSshCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#credential_store_id CredentialLibraryVaultSshCertificate#credential_store_id}
	CredentialStoreId *string `field:"required" json:"credentialStoreId" yaml:"credentialStoreId"`
	// The path in Vault to request credentials from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#path CredentialLibraryVaultSshCertificate#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The username to use with the certificate returned by the library.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#username CredentialLibraryVaultSshCertificate#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Principals to be signed as "valid_principles" in addition to username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#additional_valid_principals CredentialLibraryVaultSshCertificate#additional_valid_principals}
	AdditionalValidPrincipals *[]*string `field:"optional" json:"additionalValidPrincipals" yaml:"additionalValidPrincipals"`
	// Specifies a map of the critical options that the certificate should be signed for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#critical_options CredentialLibraryVaultSshCertificate#critical_options}
	CriticalOptions *map[string]*string `field:"optional" json:"criticalOptions" yaml:"criticalOptions"`
	// The Vault credential library description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#description CredentialLibraryVaultSshCertificate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies a map of the extensions that the certificate should be signed for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#extensions CredentialLibraryVaultSshCertificate#extensions}
	Extensions *map[string]*string `field:"optional" json:"extensions" yaml:"extensions"`
	// Specifies the number of bits to use for the generated keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#key_bits CredentialLibraryVaultSshCertificate#key_bits}
	KeyBits *float64 `field:"optional" json:"keyBits" yaml:"keyBits"`
	// Specifies the key id a certificate should have.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#key_id CredentialLibraryVaultSshCertificate#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#key_type CredentialLibraryVaultSshCertificate#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The Vault credential library name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#name CredentialLibraryVaultSshCertificate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the requested time to live for a certificate returned from the library.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/credential_library_vault_ssh_certificate#ttl CredentialLibraryVaultSshCertificate#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

