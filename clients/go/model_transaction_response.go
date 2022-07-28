/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TransactionResponse struct for TransactionResponse
type TransactionResponse struct {
	TransactionId string `json:"transactionId"`
	AccountId string `json:"accountId"`
	BusinessId string `json:"businessId"`
	Amount float32 `json:"amount"`
	Note *string `json:"note,omitempty"`
	PaymentProvider *string `json:"paymentProvider,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
	CreatedAt float32 `json:"createdAt"`
	ModifiedAt float32 `json:"modifiedAt"`
}

// NewTransactionResponse instantiates a new TransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionResponse(transactionId string, accountId string, businessId string, amount float32, createdAt float32, modifiedAt float32) *TransactionResponse {
	this := TransactionResponse{}
	this.TransactionId = transactionId
	this.AccountId = accountId
	this.BusinessId = businessId
	this.Amount = amount
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionResponseWithDefaults() *TransactionResponse {
	this := TransactionResponse{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *TransactionResponse) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TransactionResponse) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetAccountId returns the AccountId field value
func (o *TransactionResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransactionResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetBusinessId returns the BusinessId field value
func (o *TransactionResponse) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *TransactionResponse) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetAmount returns the Amount field value
func (o *TransactionResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *TransactionResponse) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *TransactionResponse) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *TransactionResponse) SetNote(v string) {
	o.Note = &v
}

// GetPaymentProvider returns the PaymentProvider field value if set, zero value otherwise.
func (o *TransactionResponse) GetPaymentProvider() string {
	if o == nil || o.PaymentProvider == nil {
		var ret string
		return ret
	}
	return *o.PaymentProvider
}

// GetPaymentProviderOk returns a tuple with the PaymentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetPaymentProviderOk() (*string, bool) {
	if o == nil || o.PaymentProvider == nil {
		return nil, false
	}
	return o.PaymentProvider, true
}

// HasPaymentProvider returns a boolean if a field has been set.
func (o *TransactionResponse) HasPaymentProvider() bool {
	if o != nil && o.PaymentProvider != nil {
		return true
	}

	return false
}

// SetPaymentProvider gets a reference to the given string and assigns it to the PaymentProvider field.
func (o *TransactionResponse) SetPaymentProvider(v string) {
	o.PaymentProvider = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *TransactionResponse) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *TransactionResponse) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *TransactionResponse) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TransactionResponse) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TransactionResponse) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *TransactionResponse) GetModifiedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *TransactionResponse) GetModifiedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *TransactionResponse) SetModifiedAt(v float32) {
	o.ModifiedAt = v
}

func (o TransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["businessId"] = o.BusinessId
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
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionResponse struct {
	value *TransactionResponse
	isSet bool
}

func (v NullableTransactionResponse) Get() *TransactionResponse {
	return v.value
}

func (v *NullableTransactionResponse) Set(val *TransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionResponse(val *TransactionResponse) *NullableTransactionResponse {
	return &NullableTransactionResponse{value: val, isSet: true}
}

func (v NullableTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


