# WorkflowJobNodeList

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
**ExtraData** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Inventory** | **int32** | Inventory applied as a prompt, assuming job template prompts for inventory | [optional] [default to null]
**ScmBranch** | **string** |  | [optional] [default to null]
**JobType** | **string** |  | [optional] [default to null]
**JobTags** | **string** |  | [optional] [default to null]
**SkipTags** | **string** |  | [optional] [default to null]
**Limit** | **string** |  | [optional] [default to null]
**DiffMode** | **bool** |  | [optional] [default to null]
**Verbosity** | **string** |  | [optional] [default to null]
**Job** | **int32** |  | [optional] [default to null]
**WorkflowJob** | **string** |  | [optional] [default to null]
**UnifiedJobTemplate** | **int32** |  | [optional] [default to null]
**SuccessNodes** | **[]int32** |  | [optional] [default to null]
**FailureNodes** | **[]int32** |  | [optional] [default to null]
**AlwaysNodes** | **[]int32** |  | [optional] [default to null]
**AllParentsMustConverge** | **bool** | If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node | [optional] [default to false]
**DoNotRun** | **bool** | Indicates that a job will not be created when True. Workflow runtime semantics will mark this True if the node is in a path that will decidedly not be ran. A value of False means the node may not run. | [optional] [default to false]
**Identifier** | **string** | An identifier coresponding to the workflow job template node that this node was created from. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

