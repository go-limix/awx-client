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

type JobRelaunch struct {
	PasswordsNeededToStart string `json:"passwords_needed_to_start,omitempty"`
	RetryCounts string `json:"retry_counts,omitempty"`
	Hosts string `json:"hosts,omitempty"`
	CredentialPasswords string `json:"credential_passwords"`
}