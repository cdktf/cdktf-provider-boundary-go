package worker

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkerConfig struct {
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
	// The scope for the worker.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/worker#scope_id Worker#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// The description for the worker.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/worker#description Worker#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name for the worker.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/worker#name Worker#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The worker authentication token required to register the worker for the worker-led authentication flow.
	//
	// Leaving this blank will result in a controller generated token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/worker#worker_generated_auth_token Worker#worker_generated_auth_token}
	WorkerGeneratedAuthToken *string `field:"optional" json:"workerGeneratedAuthToken" yaml:"workerGeneratedAuthToken"`
}

