// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package role

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleConfig struct {
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
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/role#scope_id Role#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The role description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/role#description Role#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// For Boundary 0.15+, use `grant_scope_ids` instead. The scope for which the grants in the role should apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/role#grant_scope_id Role#grant_scope_id}
	GrantScopeId *string `field:"optional" json:"grantScopeId" yaml:"grantScopeId"`
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/role#grant_scope_ids Role#grant_scope_ids}
	GrantScopeIds *[]*string `field:"optional" json:"grantScopeIds" yaml:"grantScopeIds"`
	// A list of stringified grants for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/role#grant_strings Role#grant_strings}
	GrantStrings *[]*string `field:"optional" json:"grantStrings" yaml:"grantStrings"`
	// The role name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/role#name Role#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of principal (user or group) IDs to add as principals on the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/role#principal_ids Role#principal_ids}
	PrincipalIds *[]*string `field:"optional" json:"principalIds" yaml:"principalIds"`
}

