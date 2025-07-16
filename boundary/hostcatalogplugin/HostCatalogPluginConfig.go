// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hostcatalogplugin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostCatalogPluginConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#scope_id HostCatalogPlugin#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The attributes for the host catalog.
	//
	// Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#attributes_json HostCatalogPlugin#attributes_json}
	AttributesJson *string `field:"optional" json:"attributesJson" yaml:"attributesJson"`
	// The host catalog description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#description HostCatalogPlugin#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Internal only. Used to force update so that we can always check the value of secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#internal_force_update HostCatalogPlugin#internal_force_update}
	InternalForceUpdate *string `field:"optional" json:"internalForceUpdate" yaml:"internalForceUpdate"`
	// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#internal_hmac_used_for_secrets_config_hmac HostCatalogPlugin#internal_hmac_used_for_secrets_config_hmac}
	InternalHmacUsedForSecretsConfigHmac *string `field:"optional" json:"internalHmacUsedForSecretsConfigHmac" yaml:"internalHmacUsedForSecretsConfigHmac"`
	// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#internal_secrets_config_hmac HostCatalogPlugin#internal_secrets_config_hmac}
	InternalSecretsConfigHmac *string `field:"optional" json:"internalSecretsConfigHmac" yaml:"internalSecretsConfigHmac"`
	// The host catalog name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#name HostCatalogPlugin#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the plugin that should back the resource. This or plugin_name must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#plugin_id HostCatalogPlugin#plugin_id}
	PluginId *string `field:"optional" json:"pluginId" yaml:"pluginId"`
	// The name of the plugin that should back the resource. This or plugin_id must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#plugin_name HostCatalogPlugin#plugin_name}
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// The HMAC'd secrets value returned from the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#secrets_hmac HostCatalogPlugin#secrets_hmac}
	SecretsHmac *string `field:"optional" json:"secretsHmac" yaml:"secretsHmac"`
	// The secrets for the host catalog.
	//
	// Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#secrets_json HostCatalogPlugin#secrets_json}
	SecretsJson *string `field:"optional" json:"secretsJson" yaml:"secretsJson"`
	// HCP Only. A filter used to control which PKI workers can handle dynamic host catalog requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.3.1/docs/resources/host_catalog_plugin#worker_filter HostCatalogPlugin#worker_filter}
	WorkerFilter *string `field:"optional" json:"workerFilter" yaml:"workerFilter"`
}

