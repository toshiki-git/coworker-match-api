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

// checks if the UserHobby type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserHobby{}

// UserHobby struct for UserHobby
type UserHobby struct {
	HobbyIds []string `json:"hobby_ids"`
}

type _UserHobby UserHobby

// NewUserHobby instantiates a new UserHobby object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserHobby(hobbyIds []string) *UserHobby {
	this := UserHobby{}
	this.HobbyIds = hobbyIds
	return &this
}

// NewUserHobbyWithDefaults instantiates a new UserHobby object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserHobbyWithDefaults() *UserHobby {
	this := UserHobby{}
	return &this
}

// GetHobbyIds returns the HobbyIds field value
func (o *UserHobby) GetHobbyIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HobbyIds
}

// GetHobbyIdsOk returns a tuple with the HobbyIds field value
// and a boolean to check if the value has been set.
func (o *UserHobby) GetHobbyIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HobbyIds, true
}

// SetHobbyIds sets field value
func (o *UserHobby) SetHobbyIds(v []string) {
	o.HobbyIds = v
}

func (o UserHobby) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserHobby) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hobby_ids"] = o.HobbyIds
	return toSerialize, nil
}

func (o *UserHobby) UnmarshalJSON(data []byte) (err error) {
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

	varUserHobby := _UserHobby{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserHobby)

	if err != nil {
		return err
	}

	*o = UserHobby(varUserHobby)

	return err
}

type NullableUserHobby struct {
	value *UserHobby
	isSet bool
}

func (v NullableUserHobby) Get() *UserHobby {
	return v.value
}

func (v *NullableUserHobby) Set(val *UserHobby) {
	v.value = val
	v.isSet = true
}

func (v NullableUserHobby) IsSet() bool {
	return v.isSet
}

func (v *NullableUserHobby) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserHobby(val *UserHobby) *NullableUserHobby {
	return &NullableUserHobby{value: val, isSet: true}
}

func (v NullableUserHobby) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserHobby) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


