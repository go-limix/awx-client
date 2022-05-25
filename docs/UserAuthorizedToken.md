# UserAuthorizedToken

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
**Description** | **string** |  | [optional] 
**User** | **int32** | The user representing the token owner | [optional] [default to null]
**Token** | **string** |  | [optional] [default to null]
**RefreshToken** | **string** |  | [optional] [default to null]
**Application** | **int32** |  | [default to null]
**Expires** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Scope** | **string** | Allowed scopes, further restricts user&#x27;s permissions. Must be a simple space-separated string with allowed scopes [&#x27;read&#x27;, &#x27;write&#x27;]. | [optional] [default to write]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

