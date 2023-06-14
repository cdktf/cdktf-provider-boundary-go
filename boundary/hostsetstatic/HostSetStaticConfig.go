package hostsetstatic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostSetStaticConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/host_set_static#host_catalog_id HostSetStatic#host_catalog_id}
	HostCatalogId *string `field:"required" json:"hostCatalogId" yaml:"hostCatalogId"`
	// The host set description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/host_set_static#description HostSetStatic#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of host IDs contained in this set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/host_set_static#host_ids HostSetStatic#host_ids}
	HostIds *[]*string `field:"optional" json:"hostIds" yaml:"hostIds"`
	// The host set name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/host_set_static#name HostSetStatic#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of host set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/host_set_static#type HostSetStatic#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

