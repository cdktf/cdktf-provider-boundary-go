// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package credentialsshprivatekey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CredentialSshPrivateKeyConfig struct {
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
	// ID of the credential store this credential belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/credential_ssh_private_key#credential_store_id CredentialSshPrivateKey#credential_store_id}
	CredentialStoreId *string `field:"required" json:"credentialStoreId" yaml:"credentialStoreId"`
	// The private key associated with the credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/credential_ssh_private_key#private_key CredentialSshPrivateKey#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// The username associated with the credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/credential_ssh_private_key#username CredentialSshPrivateKey#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The description of the credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/credential_ssh_private_key#description CredentialSshPrivateKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the credential. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/credential_ssh_private_key#name CredentialSshPrivateKey#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The passphrase of the private key associated with the credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/credential_ssh_private_key#private_key_passphrase CredentialSshPrivateKey#private_key_passphrase}
	PrivateKeyPassphrase *string `field:"optional" json:"privateKeyPassphrase" yaml:"privateKeyPassphrase"`
}

