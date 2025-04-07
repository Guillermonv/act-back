# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Actividad** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewActivity

`func NewActivity(date string, actividad string, status string, ) *Activity`

NewActivity instantiates a new Activity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityWithDefaults

`func NewActivityWithDefaults() *Activity`

NewActivityWithDefaults instantiates a new Activity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *Activity) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Activity) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Activity) SetDate(v string)`

SetDate sets Date field to given value.


### GetActividad

`func (o *Activity) GetActividad() string`

GetActividad returns the Actividad field if non-nil, zero value otherwise.

### GetActividadOk

`func (o *Activity) GetActividadOk() (*string, bool)`

GetActividadOk returns a tuple with the Actividad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActividad

`func (o *Activity) SetActividad(v string)`

SetActividad sets Actividad field to given value.


### GetStatus

`func (o *Activity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Activity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Activity) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


