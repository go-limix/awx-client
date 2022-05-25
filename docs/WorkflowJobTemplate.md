# WorkflowJobTemplate

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
**ExtraVars** | **string** |  | [optional] 
**Organization** | **int32** | The organization used to determine access to this template. | [optional] [default to null]
**SurveyEnabled** | **bool** |  | [optional] [default to false]
**AllowSimultaneous** | **bool** |  | [optional] [default to false]
**AskVariablesOnLaunch** | **bool** |  | [optional] [default to false]
**Inventory** | **int32** | Inventory applied as a prompt, assuming job template prompts for inventory | [optional] [default to null]
**Limit** | **string** |  | [optional] [default to null]
**ScmBranch** | **string** |  | [optional] [default to null]
**AskInventoryOnLaunch** | **bool** |  | [optional] [default to false]
**AskScmBranchOnLaunch** | **bool** |  | [optional] [default to false]
**AskLimitOnLaunch** | **bool** |  | [optional] [default to false]
**WebhookService** | **string** | Service that webhook requests will be accepted from | [optional] [default to null]
**WebhookCredential** | **int32** | Personal Access Token for posting back the status to the service API | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

