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

type OAuth2TokenResult struct {
	Count int32 `json:"count"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []OAuth2Token `json:"results"`
}
