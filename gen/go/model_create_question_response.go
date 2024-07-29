/*
CoWorkerMatch API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CreateQuestionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateQuestionResponse{}

// CreateQuestionResponse struct for CreateQuestionResponse
type CreateQuestionResponse struct {
	MatchingId string `json:"matchingId"`
	SenderUserId string `json:"senderUserId"`
	ReceiverUserId string `json:"receiverUserId"`
	MatchingDate time.Time `json:"matchingDate"`
}

type _CreateQuestionResponse CreateQuestionResponse

// NewCreateQuestionResponse instantiates a new CreateQuestionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateQuestionResponse(matchingId string, senderUserId string, receiverUserId string, matchingDate time.Time) *CreateQuestionResponse {
	this := CreateQuestionResponse{}
	this.MatchingId = matchingId
	this.SenderUserId = senderUserId
	this.ReceiverUserId = receiverUserId
	this.MatchingDate = matchingDate
	return &this
}

// NewCreateQuestionResponseWithDefaults instantiates a new CreateQuestionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateQuestionResponseWithDefaults() *CreateQuestionResponse {
	this := CreateQuestionResponse{}
	return &this
}

// GetMatchingId returns the MatchingId field value
func (o *CreateQuestionResponse) GetMatchingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchingId
}

// GetMatchingIdOk returns a tuple with the MatchingId field value
// and a boolean to check if the value has been set.
func (o *CreateQuestionResponse) GetMatchingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchingId, true
}

// SetMatchingId sets field value
func (o *CreateQuestionResponse) SetMatchingId(v string) {
	o.MatchingId = v
}

// GetSenderUserId returns the SenderUserId field value
func (o *CreateQuestionResponse) GetSenderUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderUserId
}

// GetSenderUserIdOk returns a tuple with the SenderUserId field value
// and a boolean to check if the value has been set.
func (o *CreateQuestionResponse) GetSenderUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderUserId, true
}

// SetSenderUserId sets field value
func (o *CreateQuestionResponse) SetSenderUserId(v string) {
	o.SenderUserId = v
}

// GetReceiverUserId returns the ReceiverUserId field value
func (o *CreateQuestionResponse) GetReceiverUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverUserId
}

// GetReceiverUserIdOk returns a tuple with the ReceiverUserId field value
// and a boolean to check if the value has been set.
func (o *CreateQuestionResponse) GetReceiverUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverUserId, true
}

// SetReceiverUserId sets field value
func (o *CreateQuestionResponse) SetReceiverUserId(v string) {
	o.ReceiverUserId = v
}

// GetMatchingDate returns the MatchingDate field value
func (o *CreateQuestionResponse) GetMatchingDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.MatchingDate
}

// GetMatchingDateOk returns a tuple with the MatchingDate field value
// and a boolean to check if the value has been set.
func (o *CreateQuestionResponse) GetMatchingDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchingDate, true
}

// SetMatchingDate sets field value
func (o *CreateQuestionResponse) SetMatchingDate(v time.Time) {
	o.MatchingDate = v
}

func (o CreateQuestionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateQuestionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matchingId"] = o.MatchingId
	toSerialize["senderUserId"] = o.SenderUserId
	toSerialize["receiverUserId"] = o.ReceiverUserId
	toSerialize["matchingDate"] = o.MatchingDate
	return toSerialize, nil
}

func (o *CreateQuestionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"matchingId",
		"senderUserId",
		"receiverUserId",
		"matchingDate",
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

	varCreateQuestionResponse := _CreateQuestionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateQuestionResponse)

	if err != nil {
		return err
	}

	*o = CreateQuestionResponse(varCreateQuestionResponse)

	return err
}

type NullableCreateQuestionResponse struct {
	value *CreateQuestionResponse
	isSet bool
}

func (v NullableCreateQuestionResponse) Get() *CreateQuestionResponse {
	return v.value
}

func (v *NullableCreateQuestionResponse) Set(val *CreateQuestionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateQuestionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateQuestionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateQuestionResponse(val *CreateQuestionResponse) *NullableCreateQuestionResponse {
	return &NullableCreateQuestionResponse{value: val, isSet: true}
}

func (v NullableCreateQuestionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateQuestionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


