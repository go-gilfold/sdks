/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AccountStatsResponse struct for AccountStatsResponse
type AccountStatsResponse struct {
	Total float32 `json:"total"`
	Daily []AccountStatsResponseDailyInner `json:"daily"`
	Monthly []AccountStatsResponseDailyInner `json:"monthly"`
}

// NewAccountStatsResponse instantiates a new AccountStatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatsResponse(total float32, daily []AccountStatsResponseDailyInner, monthly []AccountStatsResponseDailyInner) *AccountStatsResponse {
	this := AccountStatsResponse{}
	this.Total = total
	this.Daily = daily
	this.Monthly = monthly
	return &this
}

// NewAccountStatsResponseWithDefaults instantiates a new AccountStatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatsResponseWithDefaults() *AccountStatsResponse {
	this := AccountStatsResponse{}
	return &this
}

// GetTotal returns the Total field value
func (o *AccountStatsResponse) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AccountStatsResponse) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AccountStatsResponse) SetTotal(v float32) {
	o.Total = v
}

// GetDaily returns the Daily field value
func (o *AccountStatsResponse) GetDaily() []AccountStatsResponseDailyInner {
	if o == nil {
		var ret []AccountStatsResponseDailyInner
		return ret
	}

	return o.Daily
}

// GetDailyOk returns a tuple with the Daily field value
// and a boolean to check if the value has been set.
func (o *AccountStatsResponse) GetDailyOk() ([]AccountStatsResponseDailyInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Daily, true
}

// SetDaily sets field value
func (o *AccountStatsResponse) SetDaily(v []AccountStatsResponseDailyInner) {
	o.Daily = v
}

// GetMonthly returns the Monthly field value
func (o *AccountStatsResponse) GetMonthly() []AccountStatsResponseDailyInner {
	if o == nil {
		var ret []AccountStatsResponseDailyInner
		return ret
	}

	return o.Monthly
}

// GetMonthlyOk returns a tuple with the Monthly field value
// and a boolean to check if the value has been set.
func (o *AccountStatsResponse) GetMonthlyOk() ([]AccountStatsResponseDailyInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Monthly, true
}

// SetMonthly sets field value
func (o *AccountStatsResponse) SetMonthly(v []AccountStatsResponseDailyInner) {
	o.Monthly = v
}

func (o AccountStatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["daily"] = o.Daily
	}
	if true {
		toSerialize["monthly"] = o.Monthly
	}
	return json.Marshal(toSerialize)
}

type NullableAccountStatsResponse struct {
	value *AccountStatsResponse
	isSet bool
}

func (v NullableAccountStatsResponse) Get() *AccountStatsResponse {
	return v.value
}

func (v *NullableAccountStatsResponse) Set(val *AccountStatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatsResponse(val *AccountStatsResponse) *NullableAccountStatsResponse {
	return &NullableAccountStatsResponse{value: val, isSet: true}
}

func (v NullableAccountStatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

