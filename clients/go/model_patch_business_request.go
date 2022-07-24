/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PatchBusinessRequest struct for PatchBusinessRequest
type PatchBusinessRequest struct {
	// I am curious to see if this works with markdown like this link: [to google](https://www.google.com)
	Name *string `json:"name,omitempty"`
}

// NewPatchBusinessRequest instantiates a new PatchBusinessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBusinessRequest() *PatchBusinessRequest {
	this := PatchBusinessRequest{}
	return &this
}

// NewPatchBusinessRequestWithDefaults instantiates a new PatchBusinessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBusinessRequestWithDefaults() *PatchBusinessRequest {
	this := PatchBusinessRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchBusinessRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchBusinessRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchBusinessRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchBusinessRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePatchBusinessRequest struct {
	value *PatchBusinessRequest
	isSet bool
}

func (v NullablePatchBusinessRequest) Get() *PatchBusinessRequest {
	return v.value
}

func (v *NullablePatchBusinessRequest) Set(val *PatchBusinessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBusinessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBusinessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBusinessRequest(val *PatchBusinessRequest) *NullablePatchBusinessRequest {
	return &NullablePatchBusinessRequest{value: val, isSet: true}
}

func (v NullablePatchBusinessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBusinessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


