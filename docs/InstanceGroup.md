# InstanceGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Related** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Created** | **string** |  | [optional] [default to null]
**Modified** | **string** |  | [optional] [default to null]
**Capacity** | **string** |  | [optional] [default to null]
**CommittedCapacity** | **string** |  | [optional] [default to null]
**ConsumedCapacity** | **string** |  | [optional] [default to null]
**PercentCapacityRemaining** | **string** |  | [optional] [default to null]
**JobsRunning** | **int32** | Count of jobs in the running or waiting state that are targeted for this instance group | [optional] [default to null]
**JobsTotal** | **int32** | Count of all jobs that target this instance group | [optional] [default to null]
**Instances** | **string** |  | [optional] [default to null]
**IsContainerGroup** | **bool** | Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster. | [optional] [default to null]
**Credential** | **int32** |  | [optional] [default to null]
**PolicyInstancePercentage** | **int32** | Minimum percentage of all instances that will be automatically assigned to this group when new instances come online. | [optional] [default to null]
**PolicyInstanceMinimum** | **int32** | Static minimum number of Instances that will be automatically assign to this group when new instances come online. | [optional] [default to null]
**PolicyInstanceList** | **[]string** | List of exact-match Instances that will be assigned to this group | [optional] [default to null]
**PodSpecOverride** | **string** |  | [optional] 
**SummaryFields** | [***interface{}**](interface{}.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

