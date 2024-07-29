/*
CoWorkerMatch API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateUserHobbyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserHobbyRequest{}

// UpdateUserHobbyRequest struct for UpdateUserHobbyRequest
type UpdateUserHobbyRequest struct {
	HobbyIds []string `json:"hobby_ids"`
}

type _UpdateUserHobbyRequest UpdateUserHobbyRequest

// NewUpdateUserHobbyRequest instantiates a new UpdateUserHobbyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserHobbyRequest(hobbyIds []string) *UpdateUserHobbyRequest {
	this := UpdateUserHobbyRequest{}
	this.HobbyIds = hobbyIds
	return &this
}

// NewUpdateUserHobbyRequestWithDefaults instantiates a new UpdateUserHobbyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserHobbyRequestWithDefaults() *UpdateUserHobbyRequest {
	this := UpdateUserHobbyRequest{}
	return &this
}

// GetHobbyIds returns the HobbyIds field value
func (o *UpdateUserHobbyRequest) GetHobbyIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HobbyIds
}

// GetHobbyIdsOk returns a tuple with the HobbyIds field value
// and a boolean to check if the value has been set.
func (o *UpdateUserHobbyRequest) GetHobbyIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HobbyIds, true
}

// SetHobbyIds sets field value
func (o *UpdateUserHobbyRequest) SetHobbyIds(v []string) {
	o.HobbyIds = v
}

func (o UpdateUserHobbyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserHobbyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hobby_ids"] = o.HobbyIds
	return toSerialize, nil
}

func (o *UpdateUserHobbyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hobby_ids",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateUserHobbyRequest := _UpdateUserHobbyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateUserHobbyRequest)

	if err != nil {
		return err
	}

	*o = UpdateUserHobbyRequest(varUpdateUserHobbyRequest)

	return err
}

type NullableUpdateUserHobbyRequest struct {
	value *UpdateUserHobbyRequest
	isSet bool
}

func (v NullableUpdateUserHobbyRequest) Get() *UpdateUserHobbyRequest {
	return v.value
}

func (v *NullableUpdateUserHobbyRequest) Set(val *UpdateUserHobbyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserHobbyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserHobbyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserHobbyRequest(val *UpdateUserHobbyRequest) *NullableUpdateUserHobbyRequest {
	return &NullableUpdateUserHobbyRequest{value: val, isSet: true}
}

func (v NullableUpdateUserHobbyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserHobbyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


