/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AllAccountsResponseAccounts struct for AllAccountsResponseAccounts
type AllAccountsResponseAccounts struct {
	Items AccountResponse `json:"Items"`
	Count float32 `json:"Count"`
	ScannedCount float32 `json:"ScannedCount"`
	LastEvaluatedKey map[string]string `json:"LastEvaluatedKey"`
}

// NewAllAccountsResponseAccounts instantiates a new AllAccountsResponseAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllAccountsResponseAccounts(items AccountResponse, count float32, scannedCount float32, lastEvaluatedKey map[string]string) *AllAccountsResponseAccounts {
	this := AllAccountsResponseAccounts{}
	this.Items = items
	this.Count = count
	this.ScannedCount = scannedCount
	this.LastEvaluatedKey = lastEvaluatedKey
	return &this
}

// NewAllAccountsResponseAccountsWithDefaults instantiates a new AllAccountsResponseAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllAccountsResponseAccountsWithDefaults() *AllAccountsResponseAccounts {
	this := AllAccountsResponseAccounts{}
	return &this
}

// GetItems returns the Items field value
func (o *AllAccountsResponseAccounts) GetItems() AccountResponse {
	if o == nil {
		var ret AccountResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AllAccountsResponseAccounts) GetItemsOk() (*AccountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *AllAccountsResponseAccounts) SetItems(v AccountResponse) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *AllAccountsResponseAccounts) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *AllAccountsResponseAccounts) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *AllAccountsResponseAccounts) SetCount(v float32) {
	o.Count = v
}

// GetScannedCount returns the ScannedCount field value
func (o *AllAccountsResponseAccounts) GetScannedCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ScannedCount
}

// GetScannedCountOk returns a tuple with the ScannedCount field value
// and a boolean to check if the value has been set.
func (o *AllAccountsResponseAccounts) GetScannedCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScannedCount, true
}

// SetScannedCount sets field value
func (o *AllAccountsResponseAccounts) SetScannedCount(v float32) {
	o.ScannedCount = v
}

// GetLastEvaluatedKey returns the LastEvaluatedKey field value
func (o *AllAccountsResponseAccounts) GetLastEvaluatedKey() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.LastEvaluatedKey
}

// GetLastEvaluatedKeyOk returns a tuple with the LastEvaluatedKey field value
// and a boolean to check if the value has been set.
func (o *AllAccountsResponseAccounts) GetLastEvaluatedKeyOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEvaluatedKey, true
}

// SetLastEvaluatedKey sets field value
func (o *AllAccountsResponseAccounts) SetLastEvaluatedKey(v map[string]string) {
	o.LastEvaluatedKey = v
}

func (o AllAccountsResponseAccounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Items"] = o.Items
	}
	if true {
		toSerialize["Count"] = o.Count
	}
	if true {
		toSerialize["ScannedCount"] = o.ScannedCount
	}
	if true {
		toSerialize["LastEvaluatedKey"] = o.LastEvaluatedKey
	}
	return json.Marshal(toSerialize)
}

type NullableAllAccountsResponseAccounts struct {
	value *AllAccountsResponseAccounts
	isSet bool
}

func (v NullableAllAccountsResponseAccounts) Get() *AllAccountsResponseAccounts {
	return v.value
}

func (v *NullableAllAccountsResponseAccounts) Set(val *AllAccountsResponseAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableAllAccountsResponseAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableAllAccountsResponseAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllAccountsResponseAccounts(val *AllAccountsResponseAccounts) *NullableAllAccountsResponseAccounts {
	return &NullableAllAccountsResponseAccounts{value: val, isSet: true}
}

func (v NullableAllAccountsResponseAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllAccountsResponseAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


