/*
Amber Electric Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ActualInterval struct for ActualInterval
type ActualInterval struct {
	Type string `json:"type"`
	// Length of the interval in minutes.
	Duration float32 `json:"duration"`
	// NEM spot price (c/kWh). This is the price generators get paid to generate electricity, and what drives the variable component of your perKwh price - includes GST
	SpotPerKwh float32 `json:"spotPerKwh"`
	// Number of cents you will pay per kilowatt-hour (c/kWh) - includes GST
	PerKwh float32 `json:"perKwh"`
	// Date the interval belongs to (in NEM time). This may be different to the date component of nemTime, as the last interval of the day ends at 12:00 the following day. Formatted as a ISO 8601 date
	Date string `json:"date"`
	// The interval's NEM time. This represents the time at the end of the interval UTC+10. Formatted as a ISO 8601 time
	NemTime time.Time `json:"nemTime"`
	// Start time of the interval in UTC. Formatted as a ISO 8601 time
	StartTime time.Time `json:"startTime"`
	// End time of the interval in UTC. Formatted as a ISO 8601 time
	EndTime time.Time `json:"endTime"`
	// Percentage of renewables in the grid
	Renewables float32 `json:"renewables"`
	// Meter channel type
	ChannelType string `json:"channelType"`
	TariffInformation NullableIntervalTariffInformation `json:"tariffInformation,omitempty"`
	// Indicates whether this interval will potentially spike, or is currently in a spike state
	SpikeStatus string `json:"spikeStatus"`
	// Describes the current price. Gives you an indication of how cheap the price is in relation to the average VMO and DMO. Note: Negative is no longer used. It has been replaced with extremelyLow.
	Descriptor *string `json:"descriptor,omitempty"`
}

// NewActualInterval instantiates a new ActualInterval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActualInterval(type_ string, duration float32, spotPerKwh float32, perKwh float32, date string, nemTime time.Time, startTime time.Time, endTime time.Time, renewables float32, channelType string, spikeStatus string) *ActualInterval {
	this := ActualInterval{}
	this.Type = type_
	this.Duration = duration
	this.SpotPerKwh = spotPerKwh
	this.PerKwh = perKwh
	this.Date = date
	this.NemTime = nemTime
	this.StartTime = startTime
	this.EndTime = endTime
	this.Renewables = renewables
	this.ChannelType = channelType
	this.SpikeStatus = spikeStatus
	return &this
}

// NewActualIntervalWithDefaults instantiates a new ActualInterval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActualIntervalWithDefaults() *ActualInterval {
	this := ActualInterval{}
	return &this
}

// GetType returns the Type field value
func (o *ActualInterval) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActualInterval) SetType(v string) {
	o.Type = v
}

// GetDuration returns the Duration field value
func (o *ActualInterval) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetDurationOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *ActualInterval) SetDuration(v float32) {
	o.Duration = v
}

// GetSpotPerKwh returns the SpotPerKwh field value
func (o *ActualInterval) GetSpotPerKwh() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpotPerKwh
}

// GetSpotPerKwhOk returns a tuple with the SpotPerKwh field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetSpotPerKwhOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpotPerKwh, true
}

// SetSpotPerKwh sets field value
func (o *ActualInterval) SetSpotPerKwh(v float32) {
	o.SpotPerKwh = v
}

// GetPerKwh returns the PerKwh field value
func (o *ActualInterval) GetPerKwh() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PerKwh
}

// GetPerKwhOk returns a tuple with the PerKwh field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetPerKwhOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PerKwh, true
}

// SetPerKwh sets field value
func (o *ActualInterval) SetPerKwh(v float32) {
	o.PerKwh = v
}

// GetDate returns the Date field value
func (o *ActualInterval) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ActualInterval) SetDate(v string) {
	o.Date = v
}

// GetNemTime returns the NemTime field value
func (o *ActualInterval) GetNemTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.NemTime
}

// GetNemTimeOk returns a tuple with the NemTime field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetNemTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NemTime, true
}

// SetNemTime sets field value
func (o *ActualInterval) SetNemTime(v time.Time) {
	o.NemTime = v
}

// GetStartTime returns the StartTime field value
func (o *ActualInterval) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetStartTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ActualInterval) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *ActualInterval) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetEndTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *ActualInterval) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetRenewables returns the Renewables field value
func (o *ActualInterval) GetRenewables() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Renewables
}

// GetRenewablesOk returns a tuple with the Renewables field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetRenewablesOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Renewables, true
}

// SetRenewables sets field value
func (o *ActualInterval) SetRenewables(v float32) {
	o.Renewables = v
}

// GetChannelType returns the ChannelType field value
func (o *ActualInterval) GetChannelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelType
}

// GetChannelTypeOk returns a tuple with the ChannelType field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetChannelTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelType, true
}

// SetChannelType sets field value
func (o *ActualInterval) SetChannelType(v string) {
	o.ChannelType = v
}

// GetTariffInformation returns the TariffInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActualInterval) GetTariffInformation() IntervalTariffInformation {
	if o == nil || o.TariffInformation.Get() == nil {
		var ret IntervalTariffInformation
		return ret
	}
	return *o.TariffInformation.Get()
}

// GetTariffInformationOk returns a tuple with the TariffInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActualInterval) GetTariffInformationOk() (*IntervalTariffInformation, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TariffInformation.Get(), o.TariffInformation.IsSet()
}

// HasTariffInformation returns a boolean if a field has been set.
func (o *ActualInterval) HasTariffInformation() bool {
	if o != nil && o.TariffInformation.IsSet() {
		return true
	}

	return false
}

// SetTariffInformation gets a reference to the given NullableIntervalTariffInformation and assigns it to the TariffInformation field.
func (o *ActualInterval) SetTariffInformation(v IntervalTariffInformation) {
	o.TariffInformation.Set(&v)
}
// SetTariffInformationNil sets the value for TariffInformation to be an explicit nil
func (o *ActualInterval) SetTariffInformationNil() {
	o.TariffInformation.Set(nil)
}

// UnsetTariffInformation ensures that no value is present for TariffInformation, not even an explicit nil
func (o *ActualInterval) UnsetTariffInformation() {
	o.TariffInformation.Unset()
}

// GetSpikeStatus returns the SpikeStatus field value
func (o *ActualInterval) GetSpikeStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpikeStatus
}

// GetSpikeStatusOk returns a tuple with the SpikeStatus field value
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetSpikeStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpikeStatus, true
}

// SetSpikeStatus sets field value
func (o *ActualInterval) SetSpikeStatus(v string) {
	o.SpikeStatus = v
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *ActualInterval) GetDescriptor() string {
	if o == nil || o.Descriptor == nil {
		var ret string
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActualInterval) GetDescriptorOk() (*string, bool) {
	if o == nil || o.Descriptor == nil {
		return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *ActualInterval) HasDescriptor() bool {
	if o != nil && o.Descriptor != nil {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given string and assigns it to the Descriptor field.
func (o *ActualInterval) SetDescriptor(v string) {
	o.Descriptor = &v
}

func (o ActualInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["spotPerKwh"] = o.SpotPerKwh
	}
	if true {
		toSerialize["perKwh"] = o.PerKwh
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["nemTime"] = o.NemTime
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	if true {
		toSerialize["renewables"] = o.Renewables
	}
	if true {
		toSerialize["channelType"] = o.ChannelType
	}
	if o.TariffInformation.IsSet() {
		toSerialize["tariffInformation"] = o.TariffInformation.Get()
	}
	if true {
		toSerialize["spikeStatus"] = o.SpikeStatus
	}
	if o.Descriptor != nil {
		toSerialize["descriptor"] = o.Descriptor
	}
	return json.Marshal(toSerialize)
}

type NullableActualInterval struct {
	value *ActualInterval
	isSet bool
}

func (v NullableActualInterval) Get() *ActualInterval {
	return v.value
}

func (v *NullableActualInterval) Set(val *ActualInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableActualInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableActualInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActualInterval(val *ActualInterval) *NullableActualInterval {
	return &NullableActualInterval{value: val, isSet: true}
}

func (v NullableActualInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActualInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

