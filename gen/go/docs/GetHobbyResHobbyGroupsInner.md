# GetHobbyResHobbyGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **string** |  | 
**CategoryName** | **string** |  | 
**Hobbies** | [**[]Hobby**](Hobby.md) |  | 

## Methods

### NewGetHobbyResHobbyGroupsInner

`func NewGetHobbyResHobbyGroupsInner(categoryId string, categoryName string, hobbies []Hobby, ) *GetHobbyResHobbyGroupsInner`

NewGetHobbyResHobbyGroupsInner instantiates a new GetHobbyResHobbyGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHobbyResHobbyGroupsInnerWithDefaults

`func NewGetHobbyResHobbyGroupsInnerWithDefaults() *GetHobbyResHobbyGroupsInner`

NewGetHobbyResHobbyGroupsInnerWithDefaults instantiates a new GetHobbyResHobbyGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *GetHobbyResHobbyGroupsInner) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *GetHobbyResHobbyGroupsInner) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *GetHobbyResHobbyGroupsInner) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetCategoryName

`func (o *GetHobbyResHobbyGroupsInner) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *GetHobbyResHobbyGroupsInner) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *GetHobbyResHobbyGroupsInner) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.


### GetHobbies

`func (o *GetHobbyResHobbyGroupsInner) GetHobbies() []Hobby`

GetHobbies returns the Hobbies field if non-nil, zero value otherwise.

### GetHobbiesOk

`func (o *GetHobbyResHobbyGroupsInner) GetHobbiesOk() (*[]Hobby, bool)`

GetHobbiesOk returns a tuple with the Hobbies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHobbies

`func (o *GetHobbyResHobbyGroupsInner) SetHobbies(v []Hobby)`

SetHobbies sets Hobbies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


