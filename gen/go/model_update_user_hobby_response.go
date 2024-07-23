/*
CoWorkerMatch API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateUserHobbyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserHobbyResponse{}

// UpdateUserHobbyResponse struct for UpdateUserHobbyResponse
type UpdateUserHobbyResponse struct {
	HobbyIds []string `json:"hobby_ids,omitempty"`
}

// NewUpdateUserHobbyResponse instantiates a new UpdateUserHobbyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserHobbyResponse() *UpdateUserHobbyResponse {
	this := UpdateUserHobbyResponse{}
	return &this
}

// NewUpdateUserHobbyResponseWithDefaults instantiates a new UpdateUserHobbyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserHobbyResponseWithDefaults() *UpdateUserHobbyResponse {
	this := UpdateUserHobbyResponse{}
	return &this
}

// GetHobbyIds returns the HobbyIds field value if set, zero value otherwise.
func (o *UpdateUserHobbyResponse) GetHobbyIds() []string {
	if o == nil || IsNil(o.HobbyIds) {
		var ret []string
		return ret
	}
	return o.HobbyIds
}

// GetHobbyIdsOk returns a tuple with the HobbyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserHobbyResponse) GetHobbyIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HobbyIds) {
		return nil, false
	}
	return o.HobbyIds, true
}

// HasHobbyIds returns a boolean if a field has been set.
func (o *UpdateUserHobbyResponse) HasHobbyIds() bool {
	if o != nil && !IsNil(o.HobbyIds) {
		return true
	}

	return false
}

// SetHobbyIds gets a reference to the given []string and assigns it to the HobbyIds field.
func (o *UpdateUserHobbyResponse) SetHobbyIds(v []string) {
	o.HobbyIds = v
}

func (o UpdateUserHobbyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserHobbyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HobbyIds) {
		toSerialize["hobby_ids"] = o.HobbyIds
	}
	return toSerialize, nil
}

type NullableUpdateUserHobbyResponse struct {
	value *UpdateUserHobbyResponse
	isSet bool
}

func (v NullableUpdateUserHobbyResponse) Get() *UpdateUserHobbyResponse {
	return v.value
}

func (v *NullableUpdateUserHobbyResponse) Set(val *UpdateUserHobbyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserHobbyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserHobbyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserHobbyResponse(val *UpdateUserHobbyResponse) *NullableUpdateUserHobbyResponse {
	return &NullableUpdateUserHobbyResponse{value: val, isSet: true}
}

func (v NullableUpdateUserHobbyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserHobbyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


