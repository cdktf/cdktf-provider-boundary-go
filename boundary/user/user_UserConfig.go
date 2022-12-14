package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/user#scope_id User#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// Account ID's to associate with this user resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/user#account_ids User#account_ids}
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// The user description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/user#description User#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The username. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/user#name User#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

