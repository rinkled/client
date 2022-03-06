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

// ForecastIntervalAllOfRange When prices are particularly volatile, the API may return a range of prices that are possible.
type ForecastIntervalAllOfRange struct {
	// Estimated minimum price (c/kWh)
	Min float32 `json:"min"`
	// Estimated maximum price (c/kWh)
	Max float32 `json:"max"`
}

// NewForecastIntervalAllOfRange instantiates a new ForecastIntervalAllOfRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastIntervalAllOfRange(min float32, max float32) *ForecastIntervalAllOfRange {
	this := ForecastIntervalAllOfRange{}
	this.Min = min
	this.Max = max
	return &this
}

// NewForecastIntervalAllOfRangeWithDefaults instantiates a new ForecastIntervalAllOfRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastIntervalAllOfRangeWithDefaults() *ForecastIntervalAllOfRange {
	this := ForecastIntervalAllOfRange{}
	return &this
}

// GetMin returns the Min field value
func (o *ForecastIntervalAllOfRange) GetMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *ForecastIntervalAllOfRange) GetMinOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value
func (o *ForecastIntervalAllOfRange) SetMin(v float32) {
	o.Min = v
}

// GetMax returns the Max field value
func (o *ForecastIntervalAllOfRange) GetMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *ForecastIntervalAllOfRange) GetMaxOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *ForecastIntervalAllOfRange) SetMax(v float32) {
	o.Max = v
}

func (o ForecastIntervalAllOfRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["min"] = o.Min
	}
	if true {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableForecastIntervalAllOfRange struct {
	value *ForecastIntervalAllOfRange
	isSet bool
}

func (v NullableForecastIntervalAllOfRange) Get() *ForecastIntervalAllOfRange {
	return v.value
}

func (v *NullableForecastIntervalAllOfRange) Set(val *ForecastIntervalAllOfRange) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastIntervalAllOfRange) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastIntervalAllOfRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastIntervalAllOfRange(val *ForecastIntervalAllOfRange) *NullableForecastIntervalAllOfRange {
	return &NullableForecastIntervalAllOfRange{value: val, isSet: true}
}

func (v NullableForecastIntervalAllOfRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecastIntervalAllOfRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

