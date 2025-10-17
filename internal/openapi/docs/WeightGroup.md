# WeightGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Weights** | Pointer to [**[]Weight**](Weight.md) |  | [optional] 

## Methods

### NewWeightGroup

`func NewWeightGroup() *WeightGroup`

NewWeightGroup instantiates a new WeightGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeightGroupWithDefaults

`func NewWeightGroupWithDefaults() *WeightGroup`

NewWeightGroupWithDefaults instantiates a new WeightGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WeightGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WeightGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WeightGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WeightGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWeights

`func (o *WeightGroup) GetWeights() []Weight`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *WeightGroup) GetWeightsOk() (*[]Weight, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *WeightGroup) SetWeights(v []Weight)`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *WeightGroup) HasWeights() bool`

HasWeights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


