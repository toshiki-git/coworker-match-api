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

// checks if the CreateMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessageResponse{}

// CreateMessageResponse struct for CreateMessageResponse
type CreateMessageResponse struct {
	MessageId string `json:"message_id"`
}

type _CreateMessageResponse CreateMessageResponse

// NewCreateMessageResponse instantiates a new CreateMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessageResponse(messageId string) *CreateMessageResponse {
	this := CreateMessageResponse{}
	this.MessageId = messageId
	return &this
}

// NewCreateMessageResponseWithDefaults instantiates a new CreateMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageResponseWithDefaults() *CreateMessageResponse {
	this := CreateMessageResponse{}
	return &this
}

// GetMessageId returns the MessageId field value
func (o *CreateMessageResponse) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *CreateMessageResponse) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *CreateMessageResponse) SetMessageId(v string) {
	o.MessageId = v
}

func (o CreateMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message_id"] = o.MessageId
	return toSerialize, nil
}

func (o *CreateMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message_id",
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

	varCreateMessageResponse := _CreateMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMessageResponse)

	if err != nil {
		return err
	}

	*o = CreateMessageResponse(varCreateMessageResponse)

	return err
}

type NullableCreateMessageResponse struct {
	value *CreateMessageResponse
	isSet bool
}

func (v NullableCreateMessageResponse) Get() *CreateMessageResponse {
	return v.value
}

func (v *NullableCreateMessageResponse) Set(val *CreateMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageResponse(val *CreateMessageResponse) *NullableCreateMessageResponse {
	return &NullableCreateMessageResponse{value: val, isSet: true}
}

func (v NullableCreateMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


