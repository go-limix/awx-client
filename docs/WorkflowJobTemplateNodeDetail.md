# WorkflowJobTemplateNodeDetail

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
**WorkflowJobTemplate** | **string** |  | [default to null]
**UnifiedJobTemplate** | **int32** |  | [optional] [default to null]
**SuccessNodes** | **[]int32** |  | [optional] [default to null]
**FailureNodes** | **[]int32** |  | [optional] [default to null]
**AlwaysNodes** | **[]int32** |  | [optional] [default to null]
**AllParentsMustConverge** | **bool** | If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node | [optional] [default to false]
**Identifier** | **string** | An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. | [optional] [default to f1e30d14-6795-4941-ac3c-a0ef001ee026]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

