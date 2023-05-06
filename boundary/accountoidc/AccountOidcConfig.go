package accountoidc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountOidcConfig struct {
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
	// The resource ID for the auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.6/docs/resources/account_oidc#auth_method_id AccountOidc#auth_method_id}
	AuthMethodId *string `field:"required" json:"authMethodId" yaml:"authMethodId"`
	// The account description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.6/docs/resources/account_oidc#description AccountOidc#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The OIDC issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.6/docs/resources/account_oidc#issuer AccountOidc#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The account name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.6/docs/resources/account_oidc#name AccountOidc#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The OIDC subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.6/docs/resources/account_oidc#subject AccountOidc#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}

