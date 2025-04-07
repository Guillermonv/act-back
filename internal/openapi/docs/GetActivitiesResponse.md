# GetActivitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activities** | Pointer to [**[]Activity**](Activity.md) |  | [optional] 

## Methods

### NewGetActivitiesResponse

`func NewGetActivitiesResponse() *GetActivitiesResponse`

NewGetActivitiesResponse instantiates a new GetActivitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActivitiesResponseWithDefaults

`func NewGetActivitiesResponseWithDefaults() *GetActivitiesResponse`

NewGetActivitiesResponseWithDefaults instantiates a new GetActivitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivities

`func (o *GetActivitiesResponse) GetActivities() []Activity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *GetActivitiesResponse) GetActivitiesOk() (*[]Activity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *GetActivitiesResponse) SetActivities(v []Activity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *GetActivitiesResponse) HasActivities() bool`

HasActivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


