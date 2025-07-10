// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policystorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyStorageConfig struct {
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
	// The scope for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/policy_storage#scope_id PolicyStorage#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The number of days after which a session recording will be automatically deleted.
	//
	// Defaults to 0: never automatically delete. However, delete_after_days and retain_for_days cannot both be 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/policy_storage#delete_after_days PolicyStorage#delete_after_days}
	DeleteAfterDays *float64 `field:"optional" json:"deleteAfterDays" yaml:"deleteAfterDays"`
	// Whether or not the associated delete_after_days value can be overridden by org scopes.
	//
	// Note: if the associated delete_after_days value is 0, overridable is ignored
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/policy_storage#delete_after_overridable PolicyStorage#delete_after_overridable}
	DeleteAfterOverridable interface{} `field:"optional" json:"deleteAfterOverridable" yaml:"deleteAfterOverridable"`
	// The policy description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/policy_storage#description PolicyStorage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The policy name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/policy_storage#name PolicyStorage#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of days a session recording is required to be stored.
	//
	// Defaults to 0: allow deletions at any time. However, retain_for_days and delete_after_days cannot both be 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/policy_storage#retain_for_days PolicyStorage#retain_for_days}
	RetainForDays *float64 `field:"optional" json:"retainForDays" yaml:"retainForDays"`
	// Whether or not the associated retain_for_days value can be overridden by org scopes.
	//
	// Note: if the associated retain_for_days value is 0, overridable is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/policy_storage#retain_for_overridable PolicyStorage#retain_for_overridable}
	RetainForOverridable interface{} `field:"optional" json:"retainForOverridable" yaml:"retainForOverridable"`
}

