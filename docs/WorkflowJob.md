# WorkflowJob

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
**UnifiedJobTemplate** | **string** |  | [optional] [default to null]
**LaunchType** | **string** |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Failed** | **bool** |  | [optional] [default to null]
**Started** | [**time.Time**](time.Time.md) | The date and time the job was queued for starting. | [optional] [default to null]
**Finished** | [**time.Time**](time.Time.md) | The date and time the job finished execution. | [optional] [default to null]
**CanceledOn** | [**time.Time**](time.Time.md) | The date and time when the cancel request was sent. | [optional] [default to null]
**Elapsed** | **string** | Elapsed time in seconds that the job ran. | [optional] [default to null]
**JobArgs** | **string** |  | [optional] [default to null]
**JobCwd** | **string** |  | [optional] [default to null]
**JobEnv** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**JobExplanation** | **string** | A status field to indicate the state of the job if it wasn&#x27;t able to run and capture stdout | [optional] [default to null]
**ResultTraceback** | **string** |  | [optional] [default to null]
**LaunchedBy** | [***LaunchedBy**](Launched by.md) |  | [optional] [default to null]
**WorkUnitId** | **string** | The Receptor work unit ID associated with this job. | [optional] [default to null]
**WorkflowJobTemplate** | **string** |  | [optional] [default to null]
**ExtraVars** | **string** |  | [optional] 
**AllowSimultaneous** | **bool** |  | [optional] [default to false]
**JobTemplate** | **string** | If automatically created for a sliced job run, the job template the workflow job was created from. | [optional] [default to null]
**IsSlicedJob** | **bool** |  | [optional] [default to false]
**Inventory** | **int32** | Inventory applied as a prompt, assuming job template prompts for inventory | [optional] [default to null]
**Limit** | **string** |  | [optional] [default to null]
**ScmBranch** | **string** |  | [optional] [default to null]
**WebhookService** | **string** | Service that webhook requests will be accepted from | [optional] [default to null]
**WebhookCredential** | **int32** | Personal Access Token for posting back the status to the service API | [optional] [default to null]
**WebhookGuid** | **string** | Unique identifier of the event that triggered this webhook | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

