/*
 * Snippets API
 *
 * Test description
 *
 * API version: v1
 * Contact: contact@snippets.local
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type ResourceAccessListElement struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	SummaryFields *interface{} `json:"summary_fields,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Email string `json:"email,omitempty"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser bool `json:"is_superuser,omitempty"`
	IsSystemAuditor bool `json:"is_system_auditor,omitempty"`
	// Write-only field used to change the password.
	Password string `json:"password,omitempty"`
	LdapDn string `json:"ldap_dn,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty"`
	// Set if the account is managed by an external service
	ExternalAccount string `json:"external_account,omitempty"`
}