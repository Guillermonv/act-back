# GetWeightsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ideal** | Pointer to [**[]WeightEntry**](WeightEntry.md) |  | [optional] 
**Current** | Pointer to [**[]WeightEntry**](WeightEntry.md) |  | [optional] 

## Methods

### NewGetWeightsListResponse

`func NewGetWeightsListResponse() *GetWeightsListResponse`

NewGetWeightsListResponse instantiates a new GetWeightsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWeightsListResponseWithDefaults

`func NewGetWeightsListResponseWithDefaults() *GetWeightsListResponse`

NewGetWeightsListResponseWithDefaults instantiates a new GetWeightsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdeal

`func (o *GetWeightsListResponse) GetIdeal() []WeightEntry`

GetIdeal returns the Ideal field if non-nil, zero value otherwise.

### GetIdealOk

`func (o *GetWeightsListResponse) GetIdealOk() (*[]WeightEntry, bool)`

GetIdealOk returns a tuple with the Ideal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeal

`func (o *GetWeightsListResponse) SetIdeal(v []WeightEntry)`

SetIdeal sets Ideal field to given value.

### HasIdeal

`func (o *GetWeightsListResponse) HasIdeal() bool`

HasIdeal returns a boolean if a field has been set.

### GetCurrent

`func (o *GetWeightsListResponse) GetCurrent() []WeightEntry`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetWeightsListResponse) GetCurrentOk() (*[]WeightEntry, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetWeightsListResponse) SetCurrent(v []WeightEntry)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *GetWeightsListResponse) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


