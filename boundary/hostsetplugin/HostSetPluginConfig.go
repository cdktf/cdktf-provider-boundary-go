// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostsetplugin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostSetPluginConfig struct {
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
	// The catalog for the host set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_set_plugin#host_catalog_id HostSetPlugin#host_catalog_id}
	HostCatalogId *string `field:"required" json:"hostCatalogId" yaml:"hostCatalogId"`
	// The attributes for the host set.
	//
	// Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_set_plugin#attributes_json HostSetPlugin#attributes_json}
	AttributesJson *string `field:"optional" json:"attributesJson" yaml:"attributesJson"`
	// The host set description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_set_plugin#description HostSetPlugin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The host set name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_set_plugin#name HostSetPlugin#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ordered list of preferred endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_set_plugin#preferred_endpoints HostSetPlugin#preferred_endpoints}
	PreferredEndpoints *[]*string `field:"optional" json:"preferredEndpoints" yaml:"preferredEndpoints"`
	// The value to set for the sync interval seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_set_plugin#sync_interval_seconds HostSetPlugin#sync_interval_seconds}
	SyncIntervalSeconds *float64 `field:"optional" json:"syncIntervalSeconds" yaml:"syncIntervalSeconds"`
	// The type of host set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_set_plugin#type HostSetPlugin#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

