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

type SystemJobEvent struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	SummaryFields *interface{} `json:"summary_fields,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Event string `json:"event,omitempty"`
	Counter int32 `json:"counter,omitempty"`
	EventDisplay string `json:"event_display,omitempty"`
	EventData *interface{} `json:"event_data,omitempty"`
	Failed string `json:"failed,omitempty"`
	Changed string `json:"changed,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Stdout string `json:"stdout,omitempty"`
	StartLine int32 `json:"start_line,omitempty"`
	EndLine int32 `json:"end_line,omitempty"`
	Verbosity int32 `json:"verbosity,omitempty"`
	SystemJob string `json:"system_job,omitempty"`
}
