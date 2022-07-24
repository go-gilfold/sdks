/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PatchAccountRequest struct for PatchAccountRequest
type PatchAccountRequest struct {
	AccountName *string `json:"accountName,omitempty"`
}

// NewPatchAccountRequest instantiates a new PatchAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchAccountRequest() *PatchAccountRequest {
	this := PatchAccountRequest{}
	return &this
}

// NewPatchAccountRequestWithDefaults instantiates a new PatchAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchAccountRequestWithDefaults() *PatchAccountRequest {
	this := PatchAccountRequest{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *PatchAccountRequest) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountRequest) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *PatchAccountRequest) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *PatchAccountRequest) SetAccountName(v string) {
	o.AccountName = &v
}

func (o PatchAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountName != nil {
		toSerialize["accountName"] = o.AccountName
	}
	return json.Marshal(toSerialize)
}

type NullablePatchAccountRequest struct {
	value *PatchAccountRequest
	isSet bool
}

func (v NullablePatchAccountRequest) Get() *PatchAccountRequest {
	return v.value
}

func (v *NullablePatchAccountRequest) Set(val *PatchAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchAccountRequest(val *PatchAccountRequest) *NullablePatchAccountRequest {
	return &NullablePatchAccountRequest{value: val, isSet: true}
}

func (v NullablePatchAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


