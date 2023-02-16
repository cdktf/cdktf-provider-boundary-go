package accountpassword

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountPasswordConfig struct {
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
	// The resource ID for the auth method.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/account_password#auth_method_id AccountPassword#auth_method_id}
	AuthMethodId *string `field:"required" json:"authMethodId" yaml:"authMethodId"`
	// The resource type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/account_password#type AccountPassword#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The account description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/account_password#description AccountPassword#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The login name for this account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/account_password#login_name AccountPassword#login_name}
	LoginName *string `field:"optional" json:"loginName" yaml:"loginName"`
	// The account name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/account_password#name AccountPassword#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The account password. Only set on create, changes will not be reflected when updating account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/boundary/r/account_password#password AccountPassword#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
}

