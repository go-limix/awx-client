# Host

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
**Inventory** | **int32** |  | [default to null]
**Enabled** | **bool** | Is this host online and available for running jobs? | [optional] [default to true]
**InstanceId** | **string** | The value used by the remote inventory source to uniquely identify the host | [optional] 
**Variables** | **string** | Host variables in JSON or YAML format. | [optional] 
**HasActiveFailures** | **string** |  | [optional] [default to null]
**HasInventorySources** | **string** |  | [optional] [default to null]
**LastJob** | **string** |  | [optional] [default to null]
**LastJobHostSummary** | **int32** |  | [optional] [default to null]
**AnsibleFactsModified** | [**time.Time**](time.Time.md) | The date and time ansible_facts was last modified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

