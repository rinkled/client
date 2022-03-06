/*
Amber Electric Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Range When prices are particularly volatile, the API may return a range of prices that are possible.
type Range struct {
	// Estimated minimum price (c/kWh)
	Min float32 `json:"min"`
	// Estimated maximum price (c/kWh)
	Max float32 `json:"max"`
}

// NewRange instantiates a new Range object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRange(min float32, max float32) *Range {
	this := Range{}
	this.Min = min
	this.Max = max
	return &this
}

// NewRangeWithDefaults instantiates a new Range object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeWithDefaults() *Range {
	this := Range{}
	return &this
}

// GetMin returns the Min field value
func (o *Range) GetMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *Range) GetMinOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value
func (o *Range) SetMin(v float32) {
	o.Min = v
}

// GetMax returns the Max field value
func (o *Range) GetMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *Range) GetMaxOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *Range) SetMax(v float32) {
	o.Max = v
}

func (o Range) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["min"] = o.Min
	}
	if true {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableRange struct {
	value *Range
	isSet bool
}

func (v NullableRange) Get() *Range {
	return v.value
}

func (v *NullableRange) Set(val *Range) {
	v.value = val
	v.isSet = true
}

func (v NullableRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRange(val *Range) *NullableRange {
	return &NullableRange{value: val, isSet: true}
}

func (v NullableRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

