package credentialstorestatic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CredentialStoreStaticConfig struct {
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
	// The scope for this credential store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_static#scope_id CredentialStoreStatic#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The static credential store description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_static#description CredentialStoreStatic#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The static credential store name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_static#name CredentialStoreStatic#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

