package worker

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkerConfig struct {
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
	// The description for the worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/worker#description Worker#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name for the worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/worker#name Worker#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The scope for the worker. Defaults to `global`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/worker#scope_id Worker#scope_id}
	ScopeId *string `field:"optional" json:"scopeId" yaml:"scopeId"`
	// The worker authentication token required to register the worker for the worker-led authentication flow.
	//
	// Leaving this blank will result in a controller generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/worker#worker_generated_auth_token Worker#worker_generated_auth_token}
	WorkerGeneratedAuthToken *string `field:"optional" json:"workerGeneratedAuthToken" yaml:"workerGeneratedAuthToken"`
}

