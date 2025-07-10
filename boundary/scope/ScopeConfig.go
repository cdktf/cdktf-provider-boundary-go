// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ScopeConfig struct {
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
	// The scope ID containing the sub scope resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/scope#scope_id Scope#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user.
	//
	// Marking this true makes for simpler HCL but results in role resources that are unmanaged by Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/scope#auto_create_admin_role Scope#auto_create_admin_role}
	AutoCreateAdminRole interface{} `field:"optional" json:"autoCreateAdminRole" yaml:"autoCreateAdminRole"`
	// Only relevant when creating an org scope.
	//
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources that are unmanaged by Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/scope#auto_create_default_role Scope#auto_create_default_role}
	AutoCreateDefaultRole interface{} `field:"optional" json:"autoCreateDefaultRole" yaml:"autoCreateDefaultRole"`
	// The scope description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/scope#description Scope#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/scope#global_scope Scope#global_scope}
	GlobalScope interface{} `field:"optional" json:"globalScope" yaml:"globalScope"`
	// The scope name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/scope#name Scope#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

