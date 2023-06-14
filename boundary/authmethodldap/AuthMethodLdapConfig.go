package authmethodldap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthMethodLdapConfig struct {
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
	// The scope ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#scope_id AuthMethodLdap#scope_id}
	ScopeId *string `field:"required" json:"scopeId" yaml:"scopeId"`
	// Account attribute maps fullname and email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#account_attribute_maps AuthMethodLdap#account_attribute_maps}
	AccountAttributeMaps *[]*string `field:"optional" json:"accountAttributeMaps" yaml:"accountAttributeMaps"`
	// Use anon bind when performing LDAP group searches (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#anon_group_search AuthMethodLdap#anon_group_search}
	AnonGroupSearch interface{} `field:"optional" json:"anonGroupSearch" yaml:"anonGroupSearch"`
	// The distinguished name of entry to bind when performing user and group searches (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#bind_dn AuthMethodLdap#bind_dn}
	BindDn *string `field:"optional" json:"bindDn" yaml:"bindDn"`
	// The password to use along with bind-dn performing user and group searches (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#bind_password AuthMethodLdap#bind_password}
	BindPassword *string `field:"optional" json:"bindPassword" yaml:"bindPassword"`
	// The HMAC of the bind password returned by the Boundary controller, which is used for comparison after initial setting of the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#bind_password_hmac AuthMethodLdap#bind_password_hmac}
	BindPasswordHmac *string `field:"optional" json:"bindPasswordHmac" yaml:"bindPasswordHmac"`
	// PEM-encoded X.509 CA certificate in ASN.1 DER form that can be used as a trust anchor when connecting to an LDAP server(optional).  This may be specified multiple times.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#certificates AuthMethodLdap#certificates}
	Certificates *[]*string `field:"optional" json:"certificates" yaml:"certificates"`
	// PEM-encoded X.509 client certificate in ASN.1 DER form that can be used to authenticate against an LDAP server(optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#client_certificate AuthMethodLdap#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// PEM-encoded X.509 client certificate key in PKCS #8, ASN.1 DER form used with the client certificate (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#client_certificate_key AuthMethodLdap#client_certificate_key}
	ClientCertificateKey *string `field:"optional" json:"clientCertificateKey" yaml:"clientCertificateKey"`
	// The HMAC of the client certificate key returned by the Boundary controller, which is used for comparison after initial setting of the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#client_certificate_key_hmac AuthMethodLdap#client_certificate_key_hmac}
	ClientCertificateKeyHmac *string `field:"optional" json:"clientCertificateKeyHmac" yaml:"clientCertificateKeyHmac"`
	// The auth method description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#description AuthMethodLdap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Use anon bind to discover the bind DN of a user (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#discover_dn AuthMethodLdap#discover_dn}
	DiscoverDn interface{} `field:"optional" json:"discoverDn" yaml:"discoverDn"`
	// Find the authenticated user's groups during authentication (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#enable_groups AuthMethodLdap#enable_groups}
	EnableGroups interface{} `field:"optional" json:"enableGroups" yaml:"enableGroups"`
	// The attribute that enumerates a user's group membership from entries returned by a group search (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#group_attr AuthMethodLdap#group_attr}
	GroupAttr *string `field:"optional" json:"groupAttr" yaml:"groupAttr"`
	// The base DN under which to perform group search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#group_dn AuthMethodLdap#group_dn}
	GroupDn *string `field:"optional" json:"groupDn" yaml:"groupDn"`
	// A go template used to construct a LDAP group search filter (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#group_filter AuthMethodLdap#group_filter}
	GroupFilter *string `field:"optional" json:"groupFilter" yaml:"groupFilter"`
	// Skip the LDAP server SSL certificate validation (optional) - insecure and use with caution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#insecure_tls AuthMethodLdap#insecure_tls}
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// When true, makes this auth method the primary auth method for the scope in which it resides.
	//
	// The primary auth method for a scope means the the user will be automatically created when they login using an LDAP account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#is_primary_for_scope AuthMethodLdap#is_primary_for_scope}
	IsPrimaryForScope interface{} `field:"optional" json:"isPrimaryForScope" yaml:"isPrimaryForScope"`
	// The auth method name. Defaults to the resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#name AuthMethodLdap#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Issue StartTLS command after connecting (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#start_tls AuthMethodLdap#start_tls}
	StartTls interface{} `field:"optional" json:"startTls" yaml:"startTls"`
	// Can be one of 'inactive', 'active-private', or 'active-public'. Defaults to active-public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#state AuthMethodLdap#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The type of auth method; hardcoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#type AuthMethodLdap#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The userPrincipalDomain used to construct the UPN string for the authenticating user (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#upn_domain AuthMethodLdap#upn_domain}
	UpnDomain *string `field:"optional" json:"upnDomain" yaml:"upnDomain"`
	// The LDAP URLs that specify LDAP servers to connect to (required).  May be specified multiple times.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#urls AuthMethodLdap#urls}
	Urls *[]*string `field:"optional" json:"urls" yaml:"urls"`
	// The attribute on user entry matching the username passed when authenticating (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#user_attr AuthMethodLdap#user_attr}
	UserAttr *string `field:"optional" json:"userAttr" yaml:"userAttr"`
	// The base DN under which to perform user search (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#user_dn AuthMethodLdap#user_dn}
	UserDn *string `field:"optional" json:"userDn" yaml:"userDn"`
	// A go template used to construct a LDAP user search filter (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#user_filter AuthMethodLdap#user_filter}
	UserFilter *string `field:"optional" json:"userFilter" yaml:"userFilter"`
	// Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/boundary/1.1.8/docs/resources/auth_method_ldap#use_token_groups AuthMethodLdap#use_token_groups}
	UseTokenGroups interface{} `field:"optional" json:"useTokenGroups" yaml:"useTokenGroups"`
}

