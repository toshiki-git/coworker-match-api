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

// checks if the CreateHobbyRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateHobbyRes{}

// CreateHobbyRes struct for CreateHobbyRes
type CreateHobbyRes struct {
	HobbyId string `json:"hobbyId"`
	HobbyName string `json:"hobbyName"`
	CreatorId string `json:"creatorId"`
	CategoryId string `json:"categoryId"`
}

type _CreateHobbyRes CreateHobbyRes

// NewCreateHobbyRes instantiates a new CreateHobbyRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateHobbyRes(hobbyId string, hobbyName string, creatorId string, categoryId string) *CreateHobbyRes {
	this := CreateHobbyRes{}
	this.HobbyId = hobbyId
	this.HobbyName = hobbyName
	this.CreatorId = creatorId
	this.CategoryId = categoryId
	return &this
}

// NewCreateHobbyResWithDefaults instantiates a new CreateHobbyRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateHobbyResWithDefaults() *CreateHobbyRes {
	this := CreateHobbyRes{}
	return &this
}

// GetHobbyId returns the HobbyId field value
func (o *CreateHobbyRes) GetHobbyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HobbyId
}

// GetHobbyIdOk returns a tuple with the HobbyId field value
// and a boolean to check if the value has been set.
func (o *CreateHobbyRes) GetHobbyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HobbyId, true
}

// SetHobbyId sets field value
func (o *CreateHobbyRes) SetHobbyId(v string) {
	o.HobbyId = v
}

// GetHobbyName returns the HobbyName field value
func (o *CreateHobbyRes) GetHobbyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HobbyName
}

// GetHobbyNameOk returns a tuple with the HobbyName field value
// and a boolean to check if the value has been set.
func (o *CreateHobbyRes) GetHobbyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HobbyName, true
}

// SetHobbyName sets field value
func (o *CreateHobbyRes) SetHobbyName(v string) {
	o.HobbyName = v
}

// GetCreatorId returns the CreatorId field value
func (o *CreateHobbyRes) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *CreateHobbyRes) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *CreateHobbyRes) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetCategoryId returns the CategoryId field value
func (o *CreateHobbyRes) GetCategoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *CreateHobbyRes) GetCategoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *CreateHobbyRes) SetCategoryId(v string) {
	o.CategoryId = v
}

func (o CreateHobbyRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateHobbyRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hobbyId"] = o.HobbyId
	toSerialize["hobbyName"] = o.HobbyName
	toSerialize["creatorId"] = o.CreatorId
	toSerialize["categoryId"] = o.CategoryId
	return toSerialize, nil
}

func (o *CreateHobbyRes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hobbyId",
		"hobbyName",
		"creatorId",
		"categoryId",
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

	varCreateHobbyRes := _CreateHobbyRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateHobbyRes)

	if err != nil {
		return err
	}

	*o = CreateHobbyRes(varCreateHobbyRes)

	return err
}

type NullableCreateHobbyRes struct {
	value *CreateHobbyRes
	isSet bool
}

func (v NullableCreateHobbyRes) Get() *CreateHobbyRes {
	return v.value
}

func (v *NullableCreateHobbyRes) Set(val *CreateHobbyRes) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateHobbyRes) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateHobbyRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateHobbyRes(val *CreateHobbyRes) *NullableCreateHobbyRes {
	return &NullableCreateHobbyRes{value: val, isSet: true}
}

func (v NullableCreateHobbyRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateHobbyRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


