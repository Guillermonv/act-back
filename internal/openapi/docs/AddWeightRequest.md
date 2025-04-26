# AddWeightRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Weight** | **float64** |  | 

## Methods

### NewAddWeightRequest

`func NewAddWeightRequest(date string, weight float64, ) *AddWeightRequest`

NewAddWeightRequest instantiates a new AddWeightRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWeightRequestWithDefaults

`func NewAddWeightRequestWithDefaults() *AddWeightRequest`

NewAddWeightRequestWithDefaults instantiates a new AddWeightRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AddWeightRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AddWeightRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AddWeightRequest) SetDate(v string)`

SetDate sets Date field to given value.


### GetWeight

`func (o *AddWeightRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *AddWeightRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *AddWeightRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


