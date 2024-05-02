// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBucketConfig struct {
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
	// The name of the bucket within the external object store service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#bucket_name StorageBucket#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The scope for this storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#scope_id StorageBucket#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// Filters to the worker(s) that can handle requests for this storage bucket.
	//
	// The filter must match an existing worker in order to create a storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#worker_filter StorageBucket#worker_filter}
	WorkerFilter *string `field:"required" json:"workerFilter" yaml:"workerFilter"`
	// The attributes for the storage bucket.
	//
	// The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#attributes_json StorageBucket#attributes_json}
	AttributesJson *string `field:"optional" json:"attributesJson" yaml:"attributesJson"`
	// The prefix used to organize the data held within the external object store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#bucket_prefix StorageBucket#bucket_prefix}
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The storage bucket description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#description StorageBucket#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The storage bucket name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#name StorageBucket#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the plugin that should back the resource. This or plugin_name must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#plugin_id StorageBucket#plugin_id}
	PluginId *string `field:"optional" json:"pluginId" yaml:"pluginId"`
	// The name of the plugin that should back the resource. This or plugin_id must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#plugin_name StorageBucket#plugin_name}
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// The secrets for the storage bucket.
	//
	// Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.15/docs/resources/storage_bucket#secrets_json StorageBucket#secrets_json}
	SecretsJson *string `field:"optional" json:"secretsJson" yaml:"secretsJson"`
}

