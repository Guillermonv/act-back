# GetTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | Pointer to [**[]Task**](Task.md) |  | [optional] 

## Methods

### NewGetTasksResponse

`func NewGetTasksResponse() *GetTasksResponse`

NewGetTasksResponse instantiates a new GetTasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTasksResponseWithDefaults

`func NewGetTasksResponseWithDefaults() *GetTasksResponse`

NewGetTasksResponseWithDefaults instantiates a new GetTasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *GetTasksResponse) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *GetTasksResponse) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *GetTasksResponse) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *GetTasksResponse) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


