# User

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
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | [default to null]
**FirstName** | **string** |  | [optional] [default to null]
**LastName** | **string** |  | [optional] [default to null]
**Email** | **string** |  | [optional] [default to null]
**IsSuperuser** | **bool** | Designates that this user has all permissions without explicitly assigning them. | [optional] [default to false]
**IsSystemAuditor** | **bool** |  | [optional] [default to false]
**Password** | **string** | Write-only field used to change the password. | [optional] 
**LdapDn** | **string** |  | [optional] [default to null]
**LastLogin** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ExternalAccount** | **string** | Set if the account is managed by an external service | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

