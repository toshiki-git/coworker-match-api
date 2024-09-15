# CategoryInterest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **string** | Category&#39;s unique identifier | 
**CategoryName** | **string** | Name of the category | 
**InterestPercentage** | **int32** | User&#39;s interest percentage in this category (0-100%) | 

## Methods

### NewCategoryInterest

`func NewCategoryInterest(categoryId string, categoryName string, interestPercentage int32, ) *CategoryInterest`

NewCategoryInterest instantiates a new CategoryInterest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryInterestWithDefaults

`func NewCategoryInterestWithDefaults() *CategoryInterest`

NewCategoryInterestWithDefaults instantiates a new CategoryInterest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CategoryInterest) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryInterest) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryInterest) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetCategoryName

`func (o *CategoryInterest) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *CategoryInterest) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *CategoryInterest) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.


### GetInterestPercentage

`func (o *CategoryInterest) GetInterestPercentage() int32`

GetInterestPercentage returns the InterestPercentage field if non-nil, zero value otherwise.

### GetInterestPercentageOk

`func (o *CategoryInterest) GetInterestPercentageOk() (*int32, bool)`

GetInterestPercentageOk returns a tuple with the InterestPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPercentage

`func (o *CategoryInterest) SetInterestPercentage(v int32)`

SetInterestPercentage sets InterestPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


