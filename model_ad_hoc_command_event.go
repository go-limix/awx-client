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

type AdHocCommandEvent struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	SummaryFields *interface{} `json:"summary_fields,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	AdHocCommand string `json:"ad_hoc_command,omitempty"`
	Event string `json:"event"`
	Counter int32 `json:"counter,omitempty"`
	EventDisplay string `json:"event_display,omitempty"`
	EventData *interface{} `json:"event_data,omitempty"`
	Failed bool `json:"failed,omitempty"`
	Changed bool `json:"changed,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Host int32 `json:"host,omitempty"`
	HostName string `json:"host_name,omitempty"`
	Stdout string `json:"stdout,omitempty"`
	StartLine int32 `json:"start_line,omitempty"`
	EndLine int32 `json:"end_line,omitempty"`
	Verbosity int32 `json:"verbosity,omitempty"`
}