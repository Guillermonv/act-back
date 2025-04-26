# GetWeightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ideal** | Pointer to [**[]GetWeightsResponseIdealInner**](GetWeightsResponseIdealInner.md) |  | [optional] 
**Current** | Pointer to [**[]GetWeightsResponseIdealInner**](GetWeightsResponseIdealInner.md) |  | [optional] 

## Methods

### NewGetWeightsResponse

`func NewGetWeightsResponse() *GetWeightsResponse`

NewGetWeightsResponse instantiates a new GetWeightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWeightsResponseWithDefaults

`func NewGetWeightsResponseWithDefaults() *GetWeightsResponse`

NewGetWeightsResponseWithDefaults instantiates a new GetWeightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdeal

`func (o *GetWeightsResponse) GetIdeal() []GetWeightsResponseIdealInner`

GetIdeal returns the Ideal field if non-nil, zero value otherwise.

### GetIdealOk

`func (o *GetWeightsResponse) GetIdealOk() (*[]GetWeightsResponseIdealInner, bool)`

GetIdealOk returns a tuple with the Ideal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeal

`func (o *GetWeightsResponse) SetIdeal(v []GetWeightsResponseIdealInner)`

SetIdeal sets Ideal field to given value.

### HasIdeal

`func (o *GetWeightsResponse) HasIdeal() bool`

HasIdeal returns a boolean if a field has been set.

### GetCurrent

`func (o *GetWeightsResponse) GetCurrent() []GetWeightsResponseIdealInner`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetWeightsResponse) GetCurrentOk() (*[]GetWeightsResponseIdealInner, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetWeightsResponse) SetCurrent(v []GetWeightsResponseIdealInner)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *GetWeightsResponse) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


