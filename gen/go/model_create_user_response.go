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

// checks if the CreateUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserResponse{}

// CreateUserResponse struct for CreateUserResponse
type CreateUserResponse struct {
	User *User `json:"user,omitempty"`
}

// NewCreateUserResponse instantiates a new CreateUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserResponse() *CreateUserResponse {
	this := CreateUserResponse{}
	return &this
}

// NewCreateUserResponseWithDefaults instantiates a new CreateUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserResponseWithDefaults() *CreateUserResponse {
	this := CreateUserResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateUserResponse) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserResponse) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateUserResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *CreateUserResponse) SetUser(v User) {
	o.User = &v
}

func (o CreateUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableCreateUserResponse struct {
	value *CreateUserResponse
	isSet bool
}

func (v NullableCreateUserResponse) Get() *CreateUserResponse {
	return v.value
}

func (v *NullableCreateUserResponse) Set(val *CreateUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserResponse(val *CreateUserResponse) *NullableCreateUserResponse {
	return &NullableCreateUserResponse{value: val, isSet: true}
}

func (v NullableCreateUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


