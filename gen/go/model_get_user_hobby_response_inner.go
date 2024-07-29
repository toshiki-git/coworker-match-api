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

// checks if the GetUserHobbyResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserHobbyResponseInner{}

// GetUserHobbyResponseInner struct for GetUserHobbyResponseInner
type GetUserHobbyResponseInner struct {
	HobbyId string `json:"hobbyId"`
	HobbyName string `json:"hobbyName"`
}

type _GetUserHobbyResponseInner GetUserHobbyResponseInner

// NewGetUserHobbyResponseInner instantiates a new GetUserHobbyResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserHobbyResponseInner(hobbyId string, hobbyName string) *GetUserHobbyResponseInner {
	this := GetUserHobbyResponseInner{}
	this.HobbyId = hobbyId
	this.HobbyName = hobbyName
	return &this
}

// NewGetUserHobbyResponseInnerWithDefaults instantiates a new GetUserHobbyResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserHobbyResponseInnerWithDefaults() *GetUserHobbyResponseInner {
	this := GetUserHobbyResponseInner{}
	return &this
}

// GetHobbyId returns the HobbyId field value
func (o *GetUserHobbyResponseInner) GetHobbyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HobbyId
}

// GetHobbyIdOk returns a tuple with the HobbyId field value
// and a boolean to check if the value has been set.
func (o *GetUserHobbyResponseInner) GetHobbyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HobbyId, true
}

// SetHobbyId sets field value
func (o *GetUserHobbyResponseInner) SetHobbyId(v string) {
	o.HobbyId = v
}

// GetHobbyName returns the HobbyName field value
func (o *GetUserHobbyResponseInner) GetHobbyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HobbyName
}

// GetHobbyNameOk returns a tuple with the HobbyName field value
// and a boolean to check if the value has been set.
func (o *GetUserHobbyResponseInner) GetHobbyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HobbyName, true
}

// SetHobbyName sets field value
func (o *GetUserHobbyResponseInner) SetHobbyName(v string) {
	o.HobbyName = v
}

func (o GetUserHobbyResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserHobbyResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hobbyId"] = o.HobbyId
	toSerialize["hobbyName"] = o.HobbyName
	return toSerialize, nil
}

func (o *GetUserHobbyResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hobbyId",
		"hobbyName",
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

	varGetUserHobbyResponseInner := _GetUserHobbyResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetUserHobbyResponseInner)

	if err != nil {
		return err
	}

	*o = GetUserHobbyResponseInner(varGetUserHobbyResponseInner)

	return err
}

type NullableGetUserHobbyResponseInner struct {
	value *GetUserHobbyResponseInner
	isSet bool
}

func (v NullableGetUserHobbyResponseInner) Get() *GetUserHobbyResponseInner {
	return v.value
}

func (v *NullableGetUserHobbyResponseInner) Set(val *GetUserHobbyResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserHobbyResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserHobbyResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserHobbyResponseInner(val *GetUserHobbyResponseInner) *NullableGetUserHobbyResponseInner {
	return &NullableGetUserHobbyResponseInner{value: val, isSet: true}
}

func (v NullableGetUserHobbyResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserHobbyResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


