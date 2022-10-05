package host

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/host#host_catalog_id Host#host_catalog_id}.
	HostCatalogId *string `field:"required" json:"hostCatalogId" yaml:"hostCatalogId"`
	// The type of host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/host#type Host#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/host#address Host#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The host description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/host#description Host#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The host name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/host#name Host#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

