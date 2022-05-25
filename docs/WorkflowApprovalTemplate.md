# WorkflowApprovalTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Related** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**SummaryFields** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Created** | **string** |  | [optional] [default to null]
**Modified** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Description** | **string** |  | [optional] 
**LastJobRun** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**LastJobFailed** | **bool** |  | [optional] [default to null]
**NextJobRun** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**ExecutionEnvironment** | **int32** | The container image to be used for execution. | [optional] [default to null]
**Timeout** | **int32** | The amount of time (in seconds) before the approval node expires and fails. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

