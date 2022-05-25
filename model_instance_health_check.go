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

type InstanceHealthCheck struct {
	Uuid string `json:"uuid,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Version string `json:"version,omitempty"`
	// Last time a health check was ran on this instance to refresh cpu, memory, and capacity.
	LastHealthCheck time.Time `json:"last_health_check,omitempty"`
	// Any error details from the last health check.
	Errors string `json:"errors,omitempty"`
	Cpu string `json:"cpu,omitempty"`
	// Total system memory of this instance in bytes.
	Memory int32 `json:"memory,omitempty"`
	CpuCapacity int32 `json:"cpu_capacity,omitempty"`
	MemCapacity int32 `json:"mem_capacity,omitempty"`
	Capacity int32 `json:"capacity,omitempty"`
}
