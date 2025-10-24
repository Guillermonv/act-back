# ActivityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Activities** | Pointer to [**[]Activity**](Activity.md) |  | [optional] 

## Methods

### NewActivityGroup

`func NewActivityGroup() *ActivityGroup`

NewActivityGroup instantiates a new ActivityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityGroupWithDefaults

`func NewActivityGroupWithDefaults() *ActivityGroup`

NewActivityGroupWithDefaults instantiates a new ActivityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ActivityGroup) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityGroup) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityGroup) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ActivityGroup) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetActivities

`func (o *ActivityGroup) GetActivities() []Activity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *ActivityGroup) GetActivitiesOk() (*[]Activity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *ActivityGroup) SetActivities(v []Activity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *ActivityGroup) HasActivities() bool`

HasActivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


