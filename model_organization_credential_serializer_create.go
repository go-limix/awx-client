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

type OrganizationCredentialSerializerCreate struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	SummaryFields *interface{} `json:"summary_fields,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	// Inherit permissions from organization roles. If provided on creation, do not give either user or team.
	Organization int32 `json:"organization,omitempty"`
	// Specify the type of credential you want to create. Refer to the documentation for details on each type.
	CredentialType int32 `json:"credential_type"`
	Managed string `json:"managed,omitempty"`
	// Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.
	Inputs *interface{} `json:"inputs,omitempty"`
	Kind string `json:"kind,omitempty"`
	Cloud string `json:"cloud,omitempty"`
	Kubernetes string `json:"kubernetes,omitempty"`
}
