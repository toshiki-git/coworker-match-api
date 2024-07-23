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

// checks if the CreateUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserRequest{}

// CreateUserRequest struct for CreateUserRequest
type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	Age NullableInt32 `json:"age,omitempty"`
	Gender NullableString `json:"gender,omitempty"`
	Birthplace NullableString `json:"birthplace,omitempty"`
	JobType NullableString `json:"job_type,omitempty"`
	LineAccount NullableString `json:"line_account,omitempty"`
	DiscordAccount NullableString `json:"discord_account,omitempty"`
	Biography NullableString `json:"biography,omitempty"`
}

type _CreateUserRequest CreateUserRequest

// NewCreateUserRequest instantiates a new CreateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserRequest(name string, email string, avatarUrl string) *CreateUserRequest {
	this := CreateUserRequest{}
	this.Name = name
	this.Email = email
	this.AvatarUrl = avatarUrl
	return &this
}

// NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserRequestWithDefaults() *CreateUserRequest {
	this := CreateUserRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateUserRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateUserRequest) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *CreateUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *CreateUserRequest) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *CreateUserRequest) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserRequest) GetAge() int32 {
	if o == nil || IsNil(o.Age.Get()) {
		var ret int32
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserRequest) GetAgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *CreateUserRequest) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableInt32 and assigns it to the Age field.
func (o *CreateUserRequest) SetAge(v int32) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *CreateUserRequest) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *CreateUserRequest) UnsetAge() {
	o.Age.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserRequest) GetGender() string {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret string
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserRequest) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *CreateUserRequest) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableString and assigns it to the Gender field.
func (o *CreateUserRequest) SetGender(v string) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *CreateUserRequest) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *CreateUserRequest) UnsetGender() {
	o.Gender.Unset()
}

// GetBirthplace returns the Birthplace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserRequest) GetBirthplace() string {
	if o == nil || IsNil(o.Birthplace.Get()) {
		var ret string
		return ret
	}
	return *o.Birthplace.Get()
}

// GetBirthplaceOk returns a tuple with the Birthplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserRequest) GetBirthplaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Birthplace.Get(), o.Birthplace.IsSet()
}

// HasBirthplace returns a boolean if a field has been set.
func (o *CreateUserRequest) HasBirthplace() bool {
	if o != nil && o.Birthplace.IsSet() {
		return true
	}

	return false
}

// SetBirthplace gets a reference to the given NullableString and assigns it to the Birthplace field.
func (o *CreateUserRequest) SetBirthplace(v string) {
	o.Birthplace.Set(&v)
}
// SetBirthplaceNil sets the value for Birthplace to be an explicit nil
func (o *CreateUserRequest) SetBirthplaceNil() {
	o.Birthplace.Set(nil)
}

// UnsetBirthplace ensures that no value is present for Birthplace, not even an explicit nil
func (o *CreateUserRequest) UnsetBirthplace() {
	o.Birthplace.Unset()
}

// GetJobType returns the JobType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserRequest) GetJobType() string {
	if o == nil || IsNil(o.JobType.Get()) {
		var ret string
		return ret
	}
	return *o.JobType.Get()
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserRequest) GetJobTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobType.Get(), o.JobType.IsSet()
}

// HasJobType returns a boolean if a field has been set.
func (o *CreateUserRequest) HasJobType() bool {
	if o != nil && o.JobType.IsSet() {
		return true
	}

	return false
}

// SetJobType gets a reference to the given NullableString and assigns it to the JobType field.
func (o *CreateUserRequest) SetJobType(v string) {
	o.JobType.Set(&v)
}
// SetJobTypeNil sets the value for JobType to be an explicit nil
func (o *CreateUserRequest) SetJobTypeNil() {
	o.JobType.Set(nil)
}

// UnsetJobType ensures that no value is present for JobType, not even an explicit nil
func (o *CreateUserRequest) UnsetJobType() {
	o.JobType.Unset()
}

// GetLineAccount returns the LineAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserRequest) GetLineAccount() string {
	if o == nil || IsNil(o.LineAccount.Get()) {
		var ret string
		return ret
	}
	return *o.LineAccount.Get()
}

// GetLineAccountOk returns a tuple with the LineAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserRequest) GetLineAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineAccount.Get(), o.LineAccount.IsSet()
}

// HasLineAccount returns a boolean if a field has been set.
func (o *CreateUserRequest) HasLineAccount() bool {
	if o != nil && o.LineAccount.IsSet() {
		return true
	}

	return false
}

// SetLineAccount gets a reference to the given NullableString and assigns it to the LineAccount field.
func (o *CreateUserRequest) SetLineAccount(v string) {
	o.LineAccount.Set(&v)
}
// SetLineAccountNil sets the value for LineAccount to be an explicit nil
func (o *CreateUserRequest) SetLineAccountNil() {
	o.LineAccount.Set(nil)
}

// UnsetLineAccount ensures that no value is present for LineAccount, not even an explicit nil
func (o *CreateUserRequest) UnsetLineAccount() {
	o.LineAccount.Unset()
}

// GetDiscordAccount returns the DiscordAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserRequest) GetDiscordAccount() string {
	if o == nil || IsNil(o.DiscordAccount.Get()) {
		var ret string
		return ret
	}
	return *o.DiscordAccount.Get()
}

// GetDiscordAccountOk returns a tuple with the DiscordAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserRequest) GetDiscordAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscordAccount.Get(), o.DiscordAccount.IsSet()
}

// HasDiscordAccount returns a boolean if a field has been set.
func (o *CreateUserRequest) HasDiscordAccount() bool {
	if o != nil && o.DiscordAccount.IsSet() {
		return true
	}

	return false
}

// SetDiscordAccount gets a reference to the given NullableString and assigns it to the DiscordAccount field.
func (o *CreateUserRequest) SetDiscordAccount(v string) {
	o.DiscordAccount.Set(&v)
}
// SetDiscordAccountNil sets the value for DiscordAccount to be an explicit nil
func (o *CreateUserRequest) SetDiscordAccountNil() {
	o.DiscordAccount.Set(nil)
}

// UnsetDiscordAccount ensures that no value is present for DiscordAccount, not even an explicit nil
func (o *CreateUserRequest) UnsetDiscordAccount() {
	o.DiscordAccount.Unset()
}

// GetBiography returns the Biography field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserRequest) GetBiography() string {
	if o == nil || IsNil(o.Biography.Get()) {
		var ret string
		return ret
	}
	return *o.Biography.Get()
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserRequest) GetBiographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Biography.Get(), o.Biography.IsSet()
}

// HasBiography returns a boolean if a field has been set.
func (o *CreateUserRequest) HasBiography() bool {
	if o != nil && o.Biography.IsSet() {
		return true
	}

	return false
}

// SetBiography gets a reference to the given NullableString and assigns it to the Biography field.
func (o *CreateUserRequest) SetBiography(v string) {
	o.Biography.Set(&v)
}
// SetBiographyNil sets the value for Biography to be an explicit nil
func (o *CreateUserRequest) SetBiographyNil() {
	o.Biography.Set(nil)
}

// UnsetBiography ensures that no value is present for Biography, not even an explicit nil
func (o *CreateUserRequest) UnsetBiography() {
	o.Biography.Unset()
}

func (o CreateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["avatar_url"] = o.AvatarUrl
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.Birthplace.IsSet() {
		toSerialize["birthplace"] = o.Birthplace.Get()
	}
	if o.JobType.IsSet() {
		toSerialize["job_type"] = o.JobType.Get()
	}
	if o.LineAccount.IsSet() {
		toSerialize["line_account"] = o.LineAccount.Get()
	}
	if o.DiscordAccount.IsSet() {
		toSerialize["discord_account"] = o.DiscordAccount.Get()
	}
	if o.Biography.IsSet() {
		toSerialize["biography"] = o.Biography.Get()
	}
	return toSerialize, nil
}

func (o *CreateUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateUserRequest := _CreateUserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUserRequest)

	if err != nil {
		return err
	}

	*o = CreateUserRequest(varCreateUserRequest)

	return err
}

type NullableCreateUserRequest struct {
	value *CreateUserRequest
	isSet bool
}

func (v NullableCreateUserRequest) Get() *CreateUserRequest {
	return v.value
}

func (v *NullableCreateUserRequest) Set(val *CreateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserRequest(val *CreateUserRequest) *NullableCreateUserRequest {
	return &NullableCreateUserRequest{value: val, isSet: true}
}

func (v NullableCreateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


