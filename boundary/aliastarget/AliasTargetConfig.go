// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aliastarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AliasTargetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/alias_target#scope_id AliasTarget#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The value of the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/alias_target#value AliasTarget#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The host id to pass to Boundary when performing an authorize session action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/alias_target#authorize_session_host_id AliasTarget#authorize_session_host_id}
	AuthorizeSessionHostId *string `field:"optional" json:"authorizeSessionHostId" yaml:"authorizeSessionHostId"`
	// The alias description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/alias_target#description AliasTarget#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination of the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/alias_target#destination_id AliasTarget#destination_id}
	DestinationId *string `field:"optional" json:"destinationId" yaml:"destinationId"`
	// The alias name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/alias_target#name AliasTarget#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of alias; hardcoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/alias_target#type AliasTarget#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

