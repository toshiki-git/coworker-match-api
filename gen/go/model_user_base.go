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

// checks if the UserBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserBase{}

// UserBase struct for UserBase
type UserBase struct {
	UserName string `json:"user_name"`
	Email string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

type _UserBase UserBase

// NewUserBase instantiates a new UserBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserBase(userName string, email string, avatarUrl string) *UserBase {
	this := UserBase{}
	this.UserName = userName
	this.Email = email
	this.AvatarUrl = avatarUrl
	return &this
}

// NewUserBaseWithDefaults instantiates a new UserBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserBaseWithDefaults() *UserBase {
	this := UserBase{}
	return &this
}

// GetUserName returns the UserName field value
func (o *UserBase) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *UserBase) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *UserBase) SetUserName(v string) {
	o.UserName = v
}

// GetEmail returns the Email field value
func (o *UserBase) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserBase) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserBase) SetEmail(v string) {
	o.Email = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *UserBase) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *UserBase) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *UserBase) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

func (o UserBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_name"] = o.UserName
	toSerialize["email"] = o.Email
	toSerialize["avatar_url"] = o.AvatarUrl
	return toSerialize, nil
}

func (o *UserBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_name",
		"email",
		"avatar_url",
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

	varUserBase := _UserBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserBase)

	if err != nil {
		return err
	}

	*o = UserBase(varUserBase)

	return err
}

type NullableUserBase struct {
	value *UserBase
	isSet bool
}

func (v NullableUserBase) Get() *UserBase {
	return v.value
}

func (v *NullableUserBase) Set(val *UserBase) {
	v.value = val
	v.isSet = true
}

func (v NullableUserBase) IsSet() bool {
	return v.isSet
}

func (v *NullableUserBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserBase(val *UserBase) *NullableUserBase {
	return &NullableUserBase{value: val, isSet: true}
}

func (v NullableUserBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


