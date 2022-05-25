# OAuth2Application

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
**ClientId** | **string** |  | [optional] [default to null]
**ClientSecret** | **string** | Used for more stringent verification of access to an application when creating a token. | [optional] [default to null]
**ClientType** | **string** | Set to Public or Confidential depending on how secure the client device is. | [default to null]
**RedirectUris** | **string** | Allowed URIs list, space separated | [optional] [default to null]
**AuthorizationGrantType** | **string** | The Grant type the user must use for acquire tokens for this application. | [default to null]
**SkipAuthorization** | **bool** | Set True to skip authorization step for completely trusted applications. | [optional] [default to false]
**Organization** | **int32** | Organization containing this application. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

