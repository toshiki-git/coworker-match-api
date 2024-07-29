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

// checks if the Category type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Category{}

// Category struct for Category
type Category struct {
	CategoryId string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

type _Category Category

// NewCategory instantiates a new Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategory(categoryId string, categoryName string) *Category {
	this := Category{}
	this.CategoryId = categoryId
	this.CategoryName = categoryName
	return &this
}

// NewCategoryWithDefaults instantiates a new Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryWithDefaults() *Category {
	this := Category{}
	return &this
}

// GetCategoryId returns the CategoryId field value
func (o *Category) GetCategoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *Category) GetCategoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *Category) SetCategoryId(v string) {
	o.CategoryId = v
}

// GetCategoryName returns the CategoryName field value
func (o *Category) GetCategoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value
// and a boolean to check if the value has been set.
func (o *Category) GetCategoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryName, true
}

// SetCategoryName sets field value
func (o *Category) SetCategoryName(v string) {
	o.CategoryName = v
}

func (o Category) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Category) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["categoryId"] = o.CategoryId
	toSerialize["categoryName"] = o.CategoryName
	return toSerialize, nil
}

func (o *Category) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"categoryId",
		"categoryName",
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

	varCategory := _Category{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategory)

	if err != nil {
		return err
	}

	*o = Category(varCategory)

	return err
}

type NullableCategory struct {
	value *Category
	isSet bool
}

func (v NullableCategory) Get() *Category {
	return v.value
}

func (v *NullableCategory) Set(val *Category) {
	v.value = val
	v.isSet = true
}

func (v NullableCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategory(val *Category) *NullableCategory {
	return &NullableCategory{value: val, isSet: true}
}

func (v NullableCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


