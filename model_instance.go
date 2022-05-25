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

type Instance struct {
	Id int32 `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Related *interface{} `json:"related,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	// Last time instance ran its heartbeat task for main cluster nodes. Last known connection to receptor mesh for execution nodes.
	LastSeen time.Time `json:"last_seen,omitempty"`
	// Last time a health check was ran on this instance to refresh cpu, memory, and capacity.
	LastHealthCheck time.Time `json:"last_health_check,omitempty"`
	// Any error details from the last health check.
	Errors string `json:"errors,omitempty"`
	CapacityAdjustment string `json:"capacity_adjustment,omitempty"`
	Version string `json:"version,omitempty"`
	Capacity int32 `json:"capacity,omitempty"`
	ConsumedCapacity string `json:"consumed_capacity,omitempty"`
	PercentCapacityRemaining string `json:"percent_capacity_remaining,omitempty"`
	// Count of jobs in the running or waiting state that are targeted for this instance
	JobsRunning int32 `json:"jobs_running,omitempty"`
	// Count of all jobs that target this instance
	JobsTotal int32 `json:"jobs_total,omitempty"`
	Cpu string `json:"cpu,omitempty"`
	// Total system memory of this instance in bytes.
	Memory int32 `json:"memory,omitempty"`
	CpuCapacity int32 `json:"cpu_capacity,omitempty"`
	MemCapacity int32 `json:"mem_capacity,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	ManagedByPolicy bool `json:"managed_by_policy,omitempty"`
	NodeType string `json:"node_type,omitempty"`
}
