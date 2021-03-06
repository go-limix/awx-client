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

type InstanceGroup struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	Name string `json:"name"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Capacity string `json:"capacity,omitempty"`
	CommittedCapacity string `json:"committed_capacity,omitempty"`
	ConsumedCapacity string `json:"consumed_capacity,omitempty"`
	PercentCapacityRemaining string `json:"percent_capacity_remaining,omitempty"`
	// Count of jobs in the running or waiting state that are targeted for this instance group
	JobsRunning int32 `json:"jobs_running,omitempty"`
	// Count of all jobs that target this instance group
	JobsTotal int32 `json:"jobs_total,omitempty"`
	Instances string `json:"instances,omitempty"`
	// Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster.
	IsContainerGroup bool `json:"is_container_group,omitempty"`
	Credential int32 `json:"credential,omitempty"`
	// Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.
	PolicyInstancePercentage int32 `json:"policy_instance_percentage,omitempty"`
	// Static minimum number of Instances that will be automatically assign to this group when new instances come online.
	PolicyInstanceMinimum int32 `json:"policy_instance_minimum,omitempty"`
	// List of exact-match Instances that will be assigned to this group
	PolicyInstanceList []string `json:"policy_instance_list,omitempty"`
	PodSpecOverride string `json:"pod_spec_override,omitempty"`
	SummaryFields *interface{} `json:"summary_fields,omitempty"`
}
