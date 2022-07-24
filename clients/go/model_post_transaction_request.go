/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PostTransactionRequest struct for PostTransactionRequest
type PostTransactionRequest struct {
	AccountId string `json:"accountId"`
	Amount float32 `json:"amount"`
	Note *string `json:"note,omitempty"`
	PaymentProvider *PaymentProvider `json:"paymentProvider,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
}

// NewPostTransactionRequest instantiates a new PostTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTransactionRequest(accountId string, amount float32) *PostTransactionRequest {
	this := PostTransactionRequest{}
	this.AccountId = accountId
	this.Amount = amount
	return &this
}

// NewPostTransactionRequestWithDefaults instantiates a new PostTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTransactionRequestWithDefaults() *PostTransactionRequest {
	this := PostTransactionRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *PostTransactionRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *PostTransactionRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *PostTransactionRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *PostTransactionRequest) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PostTransactionRequest) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PostTransactionRequest) SetAmount(v float32) {
	o.Amount = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *PostTransactionRequest) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionRequest) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *PostTransactionRequest) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *PostTransactionRequest) SetNote(v string) {
	o.Note = &v
}

// GetPaymentProvider returns the PaymentProvider field value if set, zero value otherwise.
func (o *PostTransactionRequest) GetPaymentProvider() PaymentProvider {
	if o == nil || o.PaymentProvider == nil {
		var ret PaymentProvider
		return ret
	}
	return *o.PaymentProvider
}

// GetPaymentProviderOk returns a tuple with the PaymentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionRequest) GetPaymentProviderOk() (*PaymentProvider, bool) {
	if o == nil || o.PaymentProvider == nil {
		return nil, false
	}
	return o.PaymentProvider, true
}

// HasPaymentProvider returns a boolean if a field has been set.
func (o *PostTransactionRequest) HasPaymentProvider() bool {
	if o != nil && o.PaymentProvider != nil {
		return true
	}

	return false
}

// SetPaymentProvider gets a reference to the given PaymentProvider and assigns it to the PaymentProvider field.
func (o *PostTransactionRequest) SetPaymentProvider(v PaymentProvider) {
	o.PaymentProvider = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *PostTransactionRequest) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *PostTransactionRequest) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *PostTransactionRequest) SetPaymentId(v string) {
	o.PaymentId = &v
}

func (o PostTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.PaymentProvider != nil {
		toSerialize["paymentProvider"] = o.PaymentProvider
	}
	if o.PaymentId != nil {
		toSerialize["paymentId"] = o.PaymentId
	}
	return json.Marshal(toSerialize)
}

type NullablePostTransactionRequest struct {
	value *PostTransactionRequest
	isSet bool
}

func (v NullablePostTransactionRequest) Get() *PostTransactionRequest {
	return v.value
}

func (v *NullablePostTransactionRequest) Set(val *PostTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTransactionRequest(val *PostTransactionRequest) *NullablePostTransactionRequest {
	return &NullablePostTransactionRequest{value: val, isSet: true}
}

func (v NullablePostTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


