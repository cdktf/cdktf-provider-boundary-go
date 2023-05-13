package target

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TargetConfig struct {
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
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#scope_id Target#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The target resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#type Target#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Optionally, a valid network address to connect to for this target. Cannot be used alongside host_source_ids.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#address Target#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// A list of brokered credential source ID's.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#brokered_credential_source_ids Target#brokered_credential_source_ids}
	BrokeredCredentialSourceIds *[]*string `field:"optional" json:"brokeredCredentialSourceIds" yaml:"brokeredCredentialSourceIds"`
	// The default port for this target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#default_port Target#default_port}
	DefaultPort *float64 `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// The target description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#description Target#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Boolean expression to filter the workers used to access this target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#egress_worker_filter Target#egress_worker_filter}
	EgressWorkerFilter *string `field:"optional" json:"egressWorkerFilter" yaml:"egressWorkerFilter"`
	// A list of host source ID's. Cannot be used alongside address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#host_source_ids Target#host_source_ids}
	HostSourceIds *[]*string `field:"optional" json:"hostSourceIds" yaml:"hostSourceIds"`
	// HCP Only.
	//
	// Boolean expression to filter the workers a user will connect to when initiating a session against this target
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#ingress_worker_filter Target#ingress_worker_filter}
	IngressWorkerFilter *string `field:"optional" json:"ingressWorkerFilter" yaml:"ingressWorkerFilter"`
	// A list of injected application credential source ID's.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#injected_application_credential_source_ids Target#injected_application_credential_source_ids}
	InjectedApplicationCredentialSourceIds *[]*string `field:"optional" json:"injectedApplicationCredentialSourceIds" yaml:"injectedApplicationCredentialSourceIds"`
	// The target name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#name Target#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#session_connection_limit Target#session_connection_limit}.
	SessionConnectionLimit *float64 `field:"optional" json:"sessionConnectionLimit" yaml:"sessionConnectionLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#session_max_seconds Target#session_max_seconds}.
	SessionMaxSeconds *float64 `field:"optional" json:"sessionMaxSeconds" yaml:"sessionMaxSeconds"`
	// Boolean expression to filter the workers for this target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/target#worker_filter Target#worker_filter}
	WorkerFilter *string `field:"optional" json:"workerFilter" yaml:"workerFilter"`
}

