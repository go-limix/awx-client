# Project

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
**LocalPath** | **string** | Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. | [optional] [default to null]
**ScmType** | **string** | Specifies the source control system used to store the project. | [optional] [default to SCM_TYPE.EMPTY]
**ScmUrl** | **string** | The location where the project is stored. | [optional] 
**ScmBranch** | **string** | Specific branch, tag or commit to checkout. | [optional] 
**ScmRefspec** | **string** | For git projects, an additional refspec to fetch. | [optional] 
**ScmClean** | **bool** | Discard any local changes before syncing the project. | [optional] [default to false]
**ScmTrackSubmodules** | **bool** | Track submodules latest commits on defined branch. | [optional] [default to false]
**ScmDeleteOnUpdate** | **bool** | Delete the project before syncing. | [optional] [default to false]
**Credential** | **int32** |  | [optional] [default to null]
**Timeout** | **int32** | The amount of time (in seconds) to run before the task is canceled. | [optional] [default to null]
**ScmRevision** | **string** | The last revision fetched by a project update | [optional] [default to null]
**LastJobRun** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**LastJobFailed** | **bool** |  | [optional] [default to null]
**NextJobRun** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Organization** | **int32** | The organization used to determine access to this template. | [optional] [default to null]
**ScmUpdateOnLaunch** | **bool** | Update the project when a job is launched that uses the project. | [optional] [default to false]
**ScmUpdateCacheTimeout** | **int32** | The number of seconds after the last project update ran that a new project update will be launched as a job dependency. | [optional] [default to null]
**AllowOverride** | **bool** | Allow changing the SCM branch or revision in a job template that uses this project. | [optional] [default to false]
**CustomVirtualenv** | **string** | Local absolute file path containing a custom Python virtualenv to use | [optional] [default to null]
**DefaultEnvironment** | **int32** | The default execution environment for jobs run using this project. | [optional] [default to null]
**LastUpdateFailed** | **bool** |  | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

