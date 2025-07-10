// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hoststatic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostStaticConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/host_static#host_catalog_id HostStatic#host_catalog_id}.
	HostCatalogId *string `field:"required" json:"hostCatalogId" yaml:"hostCatalogId"`
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/host_static#address HostStatic#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The host description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/host_static#description HostStatic#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The host name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/host_static#name HostStatic#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.0/docs/resources/host_static#type HostStatic#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

