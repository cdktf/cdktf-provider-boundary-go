package target

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TargetConfig struct {
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
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#scope_id Target#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The target resource type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#type Target#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A list of brokered credential source ID's.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#brokered_credential_source_ids Target#brokered_credential_source_ids}
	BrokeredCredentialSourceIds *[]*string `field:"optional" json:"brokeredCredentialSourceIds" yaml:"brokeredCredentialSourceIds"`
	// The default port for this target.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#default_port Target#default_port}
	DefaultPort *float64 `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// The target description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#description Target#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of host source ID's.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#host_source_ids Target#host_source_ids}
	HostSourceIds *[]*string `field:"optional" json:"hostSourceIds" yaml:"hostSourceIds"`
	// A list of injected application credential source ID's.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#injected_application_credential_source_ids Target#injected_application_credential_source_ids}
	InjectedApplicationCredentialSourceIds *[]*string `field:"optional" json:"injectedApplicationCredentialSourceIds" yaml:"injectedApplicationCredentialSourceIds"`
	// The target name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#name Target#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#session_connection_limit Target#session_connection_limit}.
	SessionConnectionLimit *float64 `field:"optional" json:"sessionConnectionLimit" yaml:"sessionConnectionLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#session_max_seconds Target#session_max_seconds}.
	SessionMaxSeconds *float64 `field:"optional" json:"sessionMaxSeconds" yaml:"sessionMaxSeconds"`
	// Boolean expression to filter the workers for this target.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/target#worker_filter Target#worker_filter}
	WorkerFilter *string `field:"optional" json:"workerFilter" yaml:"workerFilter"`
}

