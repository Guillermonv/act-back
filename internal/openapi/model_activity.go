/*
Example API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Activity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity{}

// Activity struct for Activity
type Activity struct {
	Date string `json:"date"`
	Actividad string `json:"actividad"`
	Status string `json:"status"`
}

type _Activity Activity

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity(date string, actividad string, status string) *Activity {
	this := Activity{}
	this.Date = date
	this.Actividad = actividad
	this.Status = status
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetDate returns the Date field value
func (o *Activity) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Activity) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Activity) SetDate(v string) {
	o.Date = v
}

// GetActividad returns the Actividad field value
func (o *Activity) GetActividad() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Actividad
}

// GetActividadOk returns a tuple with the Actividad field value
// and a boolean to check if the value has been set.
func (o *Activity) GetActividadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actividad, true
}

// SetActividad sets field value
func (o *Activity) SetActividad(v string) {
	o.Actividad = v
}

// GetStatus returns the Status field value
func (o *Activity) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Activity) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Activity) SetStatus(v string) {
	o.Status = v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["actividad"] = o.Actividad
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *Activity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"actividad",
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varActivity := _Activity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivity)

	if err != nil {
		return err
	}

	*o = Activity(varActivity)

	return err
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


