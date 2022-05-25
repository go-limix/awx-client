# Instance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Related** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Uuid** | **string** |  | [optional] [default to null]
**Hostname** | **string** |  | [optional] [default to null]
**Created** | **string** |  | [optional] [default to null]
**Modified** | **string** |  | [optional] [default to null]
**LastSeen** | [**time.Time**](time.Time.md) | Last time instance ran its heartbeat task for main cluster nodes. Last known connection to receptor mesh for execution nodes. | [optional] [default to null]
**LastHealthCheck** | [**time.Time**](time.Time.md) | Last time a health check was ran on this instance to refresh cpu, memory, and capacity. | [optional] [default to null]
**Errors** | **string** | Any error details from the last health check. | [optional] [default to null]
**CapacityAdjustment** | **string** |  | [optional] [default to 1.00]
**Version** | **string** |  | [optional] [default to null]
**Capacity** | **int32** |  | [optional] [default to null]
**ConsumedCapacity** | **string** |  | [optional] [default to null]
**PercentCapacityRemaining** | **string** |  | [optional] [default to null]
**JobsRunning** | **int32** | Count of jobs in the running or waiting state that are targeted for this instance | [optional] [default to null]
**JobsTotal** | **int32** | Count of all jobs that target this instance | [optional] [default to null]
**Cpu** | **string** |  | [optional] [default to null]
**Memory** | **int32** | Total system memory of this instance in bytes. | [optional] [default to null]
**CpuCapacity** | **int32** |  | [optional] [default to null]
**MemCapacity** | **int32** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to true]
**ManagedByPolicy** | **bool** |  | [optional] [default to true]
**NodeType** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

