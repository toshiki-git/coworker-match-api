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

// checks if the GetQuestionCardResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQuestionCardResponseInner{}

// GetQuestionCardResponseInner struct for GetQuestionCardResponseInner
type GetQuestionCardResponseInner struct {
	QuestionCardId *string `json:"question_card_id,omitempty"`
	QuestionCardText *string `json:"question_card_text,omitempty"`
	IsUsed *bool `json:"is_used,omitempty"`
}

// NewGetQuestionCardResponseInner instantiates a new GetQuestionCardResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuestionCardResponseInner() *GetQuestionCardResponseInner {
	this := GetQuestionCardResponseInner{}
	return &this
}

// NewGetQuestionCardResponseInnerWithDefaults instantiates a new GetQuestionCardResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuestionCardResponseInnerWithDefaults() *GetQuestionCardResponseInner {
	this := GetQuestionCardResponseInner{}
	return &this
}

// GetQuestionCardId returns the QuestionCardId field value if set, zero value otherwise.
func (o *GetQuestionCardResponseInner) GetQuestionCardId() string {
	if o == nil || IsNil(o.QuestionCardId) {
		var ret string
		return ret
	}
	return *o.QuestionCardId
}

// GetQuestionCardIdOk returns a tuple with the QuestionCardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuestionCardResponseInner) GetQuestionCardIdOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionCardId) {
		return nil, false
	}
	return o.QuestionCardId, true
}

// HasQuestionCardId returns a boolean if a field has been set.
func (o *GetQuestionCardResponseInner) HasQuestionCardId() bool {
	if o != nil && !IsNil(o.QuestionCardId) {
		return true
	}

	return false
}

// SetQuestionCardId gets a reference to the given string and assigns it to the QuestionCardId field.
func (o *GetQuestionCardResponseInner) SetQuestionCardId(v string) {
	o.QuestionCardId = &v
}

// GetQuestionCardText returns the QuestionCardText field value if set, zero value otherwise.
func (o *GetQuestionCardResponseInner) GetQuestionCardText() string {
	if o == nil || IsNil(o.QuestionCardText) {
		var ret string
		return ret
	}
	return *o.QuestionCardText
}

// GetQuestionCardTextOk returns a tuple with the QuestionCardText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuestionCardResponseInner) GetQuestionCardTextOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionCardText) {
		return nil, false
	}
	return o.QuestionCardText, true
}

// HasQuestionCardText returns a boolean if a field has been set.
func (o *GetQuestionCardResponseInner) HasQuestionCardText() bool {
	if o != nil && !IsNil(o.QuestionCardText) {
		return true
	}

	return false
}

// SetQuestionCardText gets a reference to the given string and assigns it to the QuestionCardText field.
func (o *GetQuestionCardResponseInner) SetQuestionCardText(v string) {
	o.QuestionCardText = &v
}

// GetIsUsed returns the IsUsed field value if set, zero value otherwise.
func (o *GetQuestionCardResponseInner) GetIsUsed() bool {
	if o == nil || IsNil(o.IsUsed) {
		var ret bool
		return ret
	}
	return *o.IsUsed
}

// GetIsUsedOk returns a tuple with the IsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuestionCardResponseInner) GetIsUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUsed) {
		return nil, false
	}
	return o.IsUsed, true
}

// HasIsUsed returns a boolean if a field has been set.
func (o *GetQuestionCardResponseInner) HasIsUsed() bool {
	if o != nil && !IsNil(o.IsUsed) {
		return true
	}

	return false
}

// SetIsUsed gets a reference to the given bool and assigns it to the IsUsed field.
func (o *GetQuestionCardResponseInner) SetIsUsed(v bool) {
	o.IsUsed = &v
}

func (o GetQuestionCardResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuestionCardResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuestionCardId) {
		toSerialize["question_card_id"] = o.QuestionCardId
	}
	if !IsNil(o.QuestionCardText) {
		toSerialize["question_card_text"] = o.QuestionCardText
	}
	if !IsNil(o.IsUsed) {
		toSerialize["is_used"] = o.IsUsed
	}
	return toSerialize, nil
}

type NullableGetQuestionCardResponseInner struct {
	value *GetQuestionCardResponseInner
	isSet bool
}

func (v NullableGetQuestionCardResponseInner) Get() *GetQuestionCardResponseInner {
	return v.value
}

func (v *NullableGetQuestionCardResponseInner) Set(val *GetQuestionCardResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuestionCardResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuestionCardResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuestionCardResponseInner(val *GetQuestionCardResponseInner) *NullableGetQuestionCardResponseInner {
	return &NullableGetQuestionCardResponseInner{value: val, isSet: true}
}

func (v NullableGetQuestionCardResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuestionCardResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


