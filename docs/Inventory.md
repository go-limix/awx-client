# Inventory

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
**Organization** | **int32** | Organization containing this inventory. | [default to null]
**Kind** | **string** | Kind of inventory being represented. | [optional] [default to KIND.EMPTY]
**HostFilter** | **string** | Filter that will be applied to the hosts of this inventory. | [optional] [default to null]
**Variables** | **string** | Inventory variables in JSON or YAML format. | [optional] 
**HasActiveFailures** | **bool** | This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed. | [optional] [default to null]
**TotalHosts** | **int32** | This field is deprecated and will be removed in a future release. Total number of hosts in this inventory. | [optional] [default to null]
**HostsWithActiveFailures** | **int32** | This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures. | [optional] [default to null]
**TotalGroups** | **int32** | This field is deprecated and will be removed in a future release. Total number of groups in this inventory. | [optional] [default to null]
**HasInventorySources** | **bool** | This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources. | [optional] [default to null]
**TotalInventorySources** | **int32** | Total number of external inventory sources configured within this inventory. | [optional] [default to null]
**InventorySourcesWithFailures** | **int32** | Number of external inventory sources in this inventory with failures. | [optional] [default to null]
**PendingDeletion** | **bool** | Flag indicating the inventory is being deleted. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

