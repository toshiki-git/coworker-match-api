# GetHobbyResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **string** |  | 
**CategoryName** | **string** |  | 
**Hobbies** | [**[]Hobby**](Hobby.md) |  | 

## Methods

### NewGetHobbyResponseInner

`func NewGetHobbyResponseInner(categoryId string, categoryName string, hobbies []Hobby, ) *GetHobbyResponseInner`

NewGetHobbyResponseInner instantiates a new GetHobbyResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHobbyResponseInnerWithDefaults

`func NewGetHobbyResponseInnerWithDefaults() *GetHobbyResponseInner`

NewGetHobbyResponseInnerWithDefaults instantiates a new GetHobbyResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *GetHobbyResponseInner) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *GetHobbyResponseInner) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *GetHobbyResponseInner) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetCategoryName

`func (o *GetHobbyResponseInner) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *GetHobbyResponseInner) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *GetHobbyResponseInner) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.


### GetHobbies

`func (o *GetHobbyResponseInner) GetHobbies() []Hobby`

GetHobbies returns the Hobbies field if non-nil, zero value otherwise.

### GetHobbiesOk

`func (o *GetHobbyResponseInner) GetHobbiesOk() (*[]Hobby, bool)`

GetHobbiesOk returns a tuple with the Hobbies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHobbies

`func (o *GetHobbyResponseInner) SetHobbies(v []Hobby)`

SetHobbies sets Hobbies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


