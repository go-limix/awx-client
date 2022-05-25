# ActivityStream

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Related** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**SummaryFields** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Operation** | **string** | The action taken with respect to the given object(s). | [default to null]
**Changes** | **string** | A summary of the new and changed values when an object is created, updated, or deleted | [optional] [default to null]
**Object1** | **string** | For create, update, and delete events this is the object type that was affected. For associate and disassociate events this is the object type associated or disassociated with object2. | [default to null]
**Object2** | **string** | Unpopulated for create, update, and delete events. For associate and disassociate events this is the object type that object1 is being associated with. | [default to null]
**ObjectAssociation** | **string** | When present, shows the field name of the role or relationship that changed. | [optional] [default to null]
**ActionNode** | **string** | The cluster node the activity took place on. | [optional] [default to null]
**ObjectType** | **string** | When present, shows the model on which the role or relationship was defined. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

