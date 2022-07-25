/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AllTransactionsResponse struct for AllTransactionsResponse
type AllTransactionsResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}

// NewAllTransactionsResponse instantiates a new AllTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllTransactionsResponse(transactions []TransactionResponse) *AllTransactionsResponse {
	this := AllTransactionsResponse{}
	this.Transactions = transactions
	return &this
}

// NewAllTransactionsResponseWithDefaults instantiates a new AllTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllTransactionsResponseWithDefaults() *AllTransactionsResponse {
	this := AllTransactionsResponse{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *AllTransactionsResponse) GetTransactions() []TransactionResponse {
	if o == nil {
		var ret []TransactionResponse
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *AllTransactionsResponse) GetTransactionsOk() ([]TransactionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *AllTransactionsResponse) SetTransactions(v []TransactionResponse) {
	o.Transactions = v
}

func (o AllTransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableAllTransactionsResponse struct {
	value *AllTransactionsResponse
	isSet bool
}

func (v NullableAllTransactionsResponse) Get() *AllTransactionsResponse {
	return v.value
}

func (v *NullableAllTransactionsResponse) Set(val *AllTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllTransactionsResponse(val *AllTransactionsResponse) *NullableAllTransactionsResponse {
	return &NullableAllTransactionsResponse{value: val, isSet: true}
}

func (v NullableAllTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


