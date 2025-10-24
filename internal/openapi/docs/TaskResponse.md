# TaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Task** | Pointer to [**Task**](Task.md) |  | [optional] 

## Methods

### NewTaskResponse

`func NewTaskResponse() *TaskResponse`

NewTaskResponse instantiates a new TaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseWithDefaults

`func NewTaskResponseWithDefaults() *TaskResponse`

NewTaskResponseWithDefaults instantiates a new TaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TaskResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TaskResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTask

`func (o *TaskResponse) GetTask() Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TaskResponse) GetTaskOk() (*Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TaskResponse) SetTask(v Task)`

SetTask sets Task field to given value.

### HasTask

`func (o *TaskResponse) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


