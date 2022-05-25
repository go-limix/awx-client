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

type AdHocCommandList struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	SummaryFields *interface{} `json:"summary_fields,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name,omitempty"`
	LaunchType string `json:"launch_type,omitempty"`
	Status string `json:"status,omitempty"`
	// The container image to be used for execution.
	ExecutionEnvironment int32 `json:"execution_environment,omitempty"`
	Failed bool `json:"failed,omitempty"`
	// The date and time the job was queued for starting.
	Started time.Time `json:"started,omitempty"`
	// The date and time the job finished execution.
	Finished time.Time `json:"finished,omitempty"`
	// The date and time when the cancel request was sent.
	CanceledOn time.Time `json:"canceled_on,omitempty"`
	// Elapsed time in seconds that the job ran.
	Elapsed string `json:"elapsed,omitempty"`
	// A status field to indicate the state of the job if it wasn't able to run and capture stdout
	JobExplanation string `json:"job_explanation,omitempty"`
	// The node the job executed on.
	ExecutionNode string `json:"execution_node,omitempty"`
	// The instance that managed the execution environment.
	ControllerNode string `json:"controller_node,omitempty"`
	LaunchedBy *LaunchedBy `json:"launched_by,omitempty"`
	// The Receptor work unit ID associated with this job.
	WorkUnitId string `json:"work_unit_id,omitempty"`
	JobType string `json:"job_type,omitempty"`
	Inventory int32 `json:"inventory,omitempty"`
	Limit string `json:"limit,omitempty"`
	Credential int32 `json:"credential,omitempty"`
	ModuleName string `json:"module_name,omitempty"`
	ModuleArgs string `json:"module_args,omitempty"`
	Forks int32 `json:"forks,omitempty"`
	Verbosity int32 `json:"verbosity,omitempty"`
	ExtraVars string `json:"extra_vars,omitempty"`
	BecomeEnabled bool `json:"become_enabled,omitempty"`
	DiffMode bool `json:"diff_mode,omitempty"`
}
