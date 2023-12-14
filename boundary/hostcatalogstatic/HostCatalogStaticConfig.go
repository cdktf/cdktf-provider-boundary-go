// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostcatalogstatic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostCatalogStaticConfig struct {
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
	// The scope ID in which the resource is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/host_catalog_static#scope_id HostCatalogStatic#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The host catalog description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/host_catalog_static#description HostCatalogStatic#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The host catalog name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.11/docs/resources/host_catalog_static#name HostCatalogStatic#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

