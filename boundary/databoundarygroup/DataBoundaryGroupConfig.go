// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databoundarygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataBoundaryGroupConfig struct {
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
	// The name of the group to retrieve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/data-sources/group#name DataBoundaryGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scope ID in which the resource is created. Defaults `global` if unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/data-sources/group#scope_id DataBoundaryGroup#scope_id}
	ScopeId *string `field:"optional" json:"scopeId" yaml:"scopeId"`
}

