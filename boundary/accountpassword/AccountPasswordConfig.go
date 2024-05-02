// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountpassword

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountPasswordConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/account_password#auth_method_id AccountPassword#auth_method_id}
	AuthMethodId *string `field:"required" json:"authMethodId" yaml:"authMethodId"`
	// The account description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/account_password#description AccountPassword#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The login name for this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/account_password#login_name AccountPassword#login_name}
	LoginName *string `field:"optional" json:"loginName" yaml:"loginName"`
	// The account name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/account_password#name AccountPassword#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The account password. Only set on create, changes will not be reflected when updating account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/account_password#password AccountPassword#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/account_password#type AccountPassword#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

