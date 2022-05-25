# Schedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rrule** | **string** | A value representing the schedules iCal recurrence rule. | [default to null]
**Id** | **int32** |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Related** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**SummaryFields** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Created** | **string** |  | [optional] [default to null]
**Modified** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Description** | **string** |  | [optional] 
**ExtraData** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Inventory** | **int32** | Inventory applied as a prompt, assuming job template prompts for inventory | [optional] [default to null]
**ScmBranch** | **string** |  | [optional] [default to null]
**JobType** | **string** |  | [optional] [default to null]
**JobTags** | **string** |  | [optional] [default to null]
**SkipTags** | **string** |  | [optional] [default to null]
**Limit** | **string** |  | [optional] [default to null]
**DiffMode** | **bool** |  | [optional] [default to null]
**Verbosity** | **string** |  | [optional] [default to null]
**UnifiedJobTemplate** | **int32** |  | [default to null]
**Enabled** | **bool** | Enables processing of this schedule. | [optional] [default to true]
**Dtstart** | [**time.Time**](time.Time.md) | The first occurrence of the schedule occurs on or after this time. | [optional] [default to null]
**Dtend** | [**time.Time**](time.Time.md) | The last occurrence of the schedule occurs before this time, aftewards the schedule expires. | [optional] [default to null]
**NextRun** | [**time.Time**](time.Time.md) | The next time that the scheduled action will run. | [optional] [default to null]
**Timezone** | **string** |  | [optional] [default to null]
**Until** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

