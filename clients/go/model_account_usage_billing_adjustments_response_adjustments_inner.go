/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AccountUsageBillingAdjustmentsResponseAdjustmentsInner struct for AccountUsageBillingAdjustmentsResponseAdjustmentsInner
type AccountUsageBillingAdjustmentsResponseAdjustmentsInner struct {
	BillingId string `json:"billingId"`
	AccountId string `json:"accountId"`
	HourlyCost float32 `json:"hourlyCost"`
	CreatedAt float32 `json:"createdAt"`
	ModifiedAt float32 `json:"modifiedAt"`
}

// NewAccountUsageBillingAdjustmentsResponseAdjustmentsInner instantiates a new AccountUsageBillingAdjustmentsResponseAdjustmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUsageBillingAdjustmentsResponseAdjustmentsInner(billingId string, accountId string, hourlyCost float32, createdAt float32, modifiedAt float32) *AccountUsageBillingAdjustmentsResponseAdjustmentsInner {
	this := AccountUsageBillingAdjustmentsResponseAdjustmentsInner{}
	this.BillingId = billingId
	this.AccountId = accountId
	this.HourlyCost = hourlyCost
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewAccountUsageBillingAdjustmentsResponseAdjustmentsInnerWithDefaults instantiates a new AccountUsageBillingAdjustmentsResponseAdjustmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUsageBillingAdjustmentsResponseAdjustmentsInnerWithDefaults() *AccountUsageBillingAdjustmentsResponseAdjustmentsInner {
	this := AccountUsageBillingAdjustmentsResponseAdjustmentsInner{}
	return &this
}

// GetBillingId returns the BillingId field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetBillingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingId
}

// GetBillingIdOk returns a tuple with the BillingId field value
// and a boolean to check if the value has been set.
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetBillingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingId, true
}

// SetBillingId sets field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) SetBillingId(v string) {
	o.BillingId = v
}

// GetAccountId returns the AccountId field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) SetAccountId(v string) {
	o.AccountId = v
}

// GetHourlyCost returns the HourlyCost field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetHourlyCost() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HourlyCost
}

// GetHourlyCostOk returns a tuple with the HourlyCost field value
// and a boolean to check if the value has been set.
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetHourlyCostOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HourlyCost, true
}

// SetHourlyCost sets field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) SetHourlyCost(v float32) {
	o.HourlyCost = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetModifiedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) GetModifiedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) SetModifiedAt(v float32) {
	o.ModifiedAt = v
}

func (o AccountUsageBillingAdjustmentsResponseAdjustmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["billingId"] = o.BillingId
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["hourlyCost"] = o.HourlyCost
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner struct {
	value *AccountUsageBillingAdjustmentsResponseAdjustmentsInner
	isSet bool
}

func (v NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner) Get() *AccountUsageBillingAdjustmentsResponseAdjustmentsInner {
	return v.value
}

func (v *NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner) Set(val *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner(val *AccountUsageBillingAdjustmentsResponseAdjustmentsInner) *NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner {
	return &NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner{value: val, isSet: true}
}

func (v NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUsageBillingAdjustmentsResponseAdjustmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

