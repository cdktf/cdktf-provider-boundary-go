package group

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/group#scope_id Group#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The group description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/group#description Group#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Resource IDs for group members, these are most likely boundary users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/group#member_ids Group#member_ids}
	MemberIds *[]*string `field:"optional" json:"memberIds" yaml:"memberIds"`
	// The group name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/group#name Group#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

