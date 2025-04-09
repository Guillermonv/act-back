# GetActivitiesGroupedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activities** | Pointer to [**map[string][]ActivityDateStatus**](array.md) |  | [optional] 

## Methods

### NewGetActivitiesGroupedResponse

`func NewGetActivitiesGroupedResponse() *GetActivitiesGroupedResponse`

NewGetActivitiesGroupedResponse instantiates a new GetActivitiesGroupedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivitiesGroupedResponseWithDefaults

`func NewGetActivitiesGroupedResponseWithDefaults() *GetActivitiesGroupedResponse`

NewGetActivitiesGroupedResponseWithDefaults instantiates a new GetActivitiesGroupedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivities

`func (o *GetActivitiesGroupedResponse) GetActivities() map[string][]ActivityDateStatus`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *GetActivitiesGroupedResponse) GetActivitiesOk() (*map[string][]ActivityDateStatus, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *GetActivitiesGroupedResponse) SetActivities(v map[string][]ActivityDateStatus)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *GetActivitiesGroupedResponse) HasActivities() bool`

HasActivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


