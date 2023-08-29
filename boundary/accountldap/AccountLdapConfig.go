// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountldap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountLdapConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/account_ldap#auth_method_id AccountLdap#auth_method_id}
	AuthMethodId *string `field:"required" json:"authMethodId" yaml:"authMethodId"`
	// The account description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/account_ldap#description AccountLdap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The login name for this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/account_ldap#login_name AccountLdap#login_name}
	LoginName *string `field:"optional" json:"loginName" yaml:"loginName"`
	// The account name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/account_ldap#name AccountLdap#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.9/docs/resources/account_ldap#type AccountLdap#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

