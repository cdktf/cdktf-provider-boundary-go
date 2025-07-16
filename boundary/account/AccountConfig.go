// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package account

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountConfig struct {
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
	// The resource ID for the auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/account#auth_method_id Account#auth_method_id}
	AuthMethodId *string `field:"required" json:"authMethodId" yaml:"authMethodId"`
	// The resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/account#type Account#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The account description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/account#description Account#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The login name for this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/account#login_name Account#login_name}
	LoginName *string `field:"optional" json:"loginName" yaml:"loginName"`
	// The account name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/account#name Account#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The account password. Only set on create, changes will not be reflected when updating account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/account#password Account#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
}

