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

type InventoryUpdateList struct {
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
	// The node the job executed on.
	ExecutionNode string `json:"execution_node,omitempty"`
	LaunchedBy *LaunchedBy `json:"launched_by,omitempty"`
	// The Receptor work unit ID associated with this job.
	WorkUnitId string `json:"work_unit_id,omitempty"`
	Source string `json:"source,omitempty"`
	SourcePath string `json:"source_path,omitempty"`
	// Inventory source variables in YAML or JSON format.
	SourceVars string `json:"source_vars,omitempty"`
	// Cloud credential to use for inventory updates.
	Credential int32 `json:"credential,omitempty"`
	// Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)
	EnabledVar string `json:"enabled_var,omitempty"`
	// Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled
	EnabledValue string `json:"enabled_value,omitempty"`
	// Regex where only matching hosts will be imported.
	HostFilter string `json:"host_filter,omitempty"`
	// Overwrite local groups and hosts from remote inventory source.
	Overwrite bool `json:"overwrite,omitempty"`
	// Overwrite local variables from remote inventory source.
	OverwriteVars bool `json:"overwrite_vars,omitempty"`
	CustomVirtualenv string `json:"custom_virtualenv,omitempty"`
	// The amount of time (in seconds) to run before the task is canceled.
	Timeout int32 `json:"timeout,omitempty"`
	Verbosity int32 `json:"verbosity,omitempty"`
	Inventory int32 `json:"inventory,omitempty"`
	InventorySource string `json:"inventory_source,omitempty"`
	LicenseError bool `json:"license_error,omitempty"`
	OrgHostLimitError bool `json:"org_host_limit_error,omitempty"`
	// Inventory files from this Project Update were used for the inventory update.
	SourceProjectUpdate string `json:"source_project_update,omitempty"`
	// The Instance group the job was run under
	InstanceGroup int32 `json:"instance_group,omitempty"`
}
