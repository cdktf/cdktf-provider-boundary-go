package credentialstorevault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CredentialStoreVaultConfig struct {
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
	// The address to Vault server. This should be a complete URL such as 'https://127.0.0.1:8200'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#address CredentialStoreVault#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// The scope for this credential store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#scope_id CredentialStoreVault#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// A token used for accessing Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#token CredentialStoreVault#token}
	Token *string `field:"required" json:"token" yaml:"token"`
	// A PEM-encoded CA certificate to verify the Vault server's TLS certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#ca_cert CredentialStoreVault#ca_cert}
	CaCert *string `field:"optional" json:"caCert" yaml:"caCert"`
	// A PEM-encoded client certificate to use for TLS authentication to the Vault server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#client_certificate CredentialStoreVault#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// A PEM-encoded private key matching the client certificate from 'client_certificate'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#client_certificate_key CredentialStoreVault#client_certificate_key}
	ClientCertificateKey *string `field:"optional" json:"clientCertificateKey" yaml:"clientCertificateKey"`
	// The Vault credential store description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#description CredentialStoreVault#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Vault credential store name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#name CredentialStoreVault#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace within Vault to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#namespace CredentialStoreVault#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Name to use as the SNI host when connecting to Vault via TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#tls_server_name CredentialStoreVault#tls_server_name}
	TlsServerName *string `field:"optional" json:"tlsServerName" yaml:"tlsServerName"`
	// Whether or not to skip TLS verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#tls_skip_verify CredentialStoreVault#tls_skip_verify}
	TlsSkipVerify interface{} `field:"optional" json:"tlsSkipVerify" yaml:"tlsSkipVerify"`
	// HCP Only.
	//
	// A filter used to control which PKI workers can handle Vault requests. This allows the use of private Vault instances with Boundary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.7/docs/resources/credential_store_vault#worker_filter CredentialStoreVault#worker_filter}
	WorkerFilter *string `field:"optional" json:"workerFilter" yaml:"workerFilter"`
}

