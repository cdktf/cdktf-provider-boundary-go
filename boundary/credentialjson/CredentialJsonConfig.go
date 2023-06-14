package credentialjson

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CredentialJsonConfig struct {
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
	// The credential store in which to save this json credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/credential_json#credential_store_id CredentialJson#credential_store_id}
	CredentialStoreId *string `field:"required" json:"credentialStoreId" yaml:"credentialStoreId"`
	// The object for the this json credential.
	//
	// Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/credential_json#object CredentialJson#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// The description of this json credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/credential_json#description CredentialJson#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of this json credential. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/credential_json#name CredentialJson#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

