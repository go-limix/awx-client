# InstanceHealthCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | [optional] [default to null]
**Hostname** | **string** |  | [optional] [default to null]
**Version** | **string** |  | [optional] [default to null]
**LastHealthCheck** | [**time.Time**](time.Time.md) | Last time a health check was ran on this instance to refresh cpu, memory, and capacity. | [optional] [default to null]
**Errors** | **string** | Any error details from the last health check. | [optional] [default to null]
**Cpu** | **string** |  | [optional] [default to null]
**Memory** | **int32** | Total system memory of this instance in bytes. | [optional] [default to null]
**CpuCapacity** | **int32** |  | [optional] [default to null]
**MemCapacity** | **int32** |  | [optional] [default to null]
**Capacity** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

