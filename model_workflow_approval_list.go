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

type WorkflowApprovalList struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	SummaryFields *interface{} `json:"summary_fields,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	UnifiedJobTemplate string `json:"unified_job_template,omitempty"`
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
	LaunchedBy *LaunchedBy `json:"launched_by,omitempty"`
	// The Receptor work unit ID associated with this job.
	WorkUnitId string `json:"work_unit_id,omitempty"`
	CanApproveOrDeny string `json:"can_approve_or_deny,omitempty"`
	ApprovalExpiration string `json:"approval_expiration,omitempty"`
	TimedOut string `json:"timed_out,omitempty"`
}
