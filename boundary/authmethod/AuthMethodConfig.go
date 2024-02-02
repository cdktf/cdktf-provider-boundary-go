// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authmethod

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthMethodConfig struct {
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
	// The scope ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/auth_method#scope_id AuthMethod#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/auth_method#type AuthMethod#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The auth method description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/auth_method#description AuthMethod#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The minimum login name length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/auth_method#min_login_name_length AuthMethod#min_login_name_length}
	MinLoginNameLength *float64 `field:"optional" json:"minLoginNameLength" yaml:"minLoginNameLength"`
	// The minimum password length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/auth_method#min_password_length AuthMethod#min_password_length}
	MinPasswordLength *float64 `field:"optional" json:"minPasswordLength" yaml:"minPasswordLength"`
	// The auth method name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.13/docs/resources/auth_method#name AuthMethod#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

