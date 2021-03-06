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

// Channel Describes a power meter channel.  The General channel provides continuous power - it's the channel all of your appliances and lights are attached to.  Controlled loads are only on for a limited time during the day (usually when the load on the network is low, or generation is high) - you may have your hot water system attached to this channel.  The feed in channel sends power back to the grid - you will have these types of channels if you have solar or batteries.
type Channel struct {
	// Identifier of the channel
	Identifier string `json:"identifier"`
	// Channel type.
	Type string `json:"type"`
	// The tariff code of the channel
	Tariff string `json:"tariff"`
}

// NewChannel instantiates a new Channel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannel(identifier string, type_ string, tariff string) *Channel {
	this := Channel{}
	this.Identifier = identifier
	this.Type = type_
	this.Tariff = tariff
	return &this
}

// NewChannelWithDefaults instantiates a new Channel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelWithDefaults() *Channel {
	this := Channel{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *Channel) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *Channel) GetIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *Channel) SetIdentifier(v string) {
	o.Identifier = v
}

// GetType returns the Type field value
func (o *Channel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Channel) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Channel) SetType(v string) {
	o.Type = v
}

// GetTariff returns the Tariff field value
func (o *Channel) GetTariff() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tariff
}

// GetTariffOk returns a tuple with the Tariff field value
// and a boolean to check if the value has been set.
func (o *Channel) GetTariffOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tariff, true
}

// SetTariff sets field value
func (o *Channel) SetTariff(v string) {
	o.Tariff = v
}

func (o Channel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["tariff"] = o.Tariff
	}
	return json.Marshal(toSerialize)
}

type NullableChannel struct {
	value *Channel
	isSet bool
}

func (v NullableChannel) Get() *Channel {
	return v.value
}

func (v *NullableChannel) Set(val *Channel) {
	v.value = val
	v.isSet = true
}

func (v NullableChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannel(val *Channel) *NullableChannel {
	return &NullableChannel{value: val, isSet: true}
}

func (v NullableChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


