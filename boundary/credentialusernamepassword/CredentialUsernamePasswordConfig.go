package credentialusernamepassword

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CredentialUsernamePasswordConfig struct {
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
	// The credential store in which to save this username/password credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/credential_username_password#credential_store_id CredentialUsernamePassword#credential_store_id}
	CredentialStoreId *string `field:"required" json:"credentialStoreId" yaml:"credentialStoreId"`
	// The password of this username/password credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/credential_username_password#password CredentialUsernamePassword#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username of this username/password credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/credential_username_password#username CredentialUsernamePassword#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The description of this username/password credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/credential_username_password#description CredentialUsernamePassword#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of this username/password credential. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/credential_username_password#name CredentialUsernamePassword#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

