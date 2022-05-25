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

type WorkflowApprovalTemplate struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	SummaryFields *interface{} `json:"summary_fields,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	LastJobRun time.Time `json:"last_job_run,omitempty"`
	LastJobFailed bool `json:"last_job_failed,omitempty"`
	NextJobRun time.Time `json:"next_job_run,omitempty"`
	Status string `json:"status,omitempty"`
	// The container image to be used for execution.
	ExecutionEnvironment int32 `json:"execution_environment,omitempty"`
	// The amount of time (in seconds) before the approval node expires and fails.
	Timeout int32 `json:"timeout,omitempty"`
}