# GetUserCategoryPercentagesRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User&#39;s unique identifier | 
**Categories** | [**[]CategoryInterest**](CategoryInterest.md) | List of categories with user&#39;s interest levels | 

## Methods

### NewGetUserCategoryPercentagesRes

`func NewGetUserCategoryPercentagesRes(userId string, categories []CategoryInterest, ) *GetUserCategoryPercentagesRes`

NewGetUserCategoryPercentagesRes instantiates a new GetUserCategoryPercentagesRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserCategoryPercentagesResWithDefaults

`func NewGetUserCategoryPercentagesResWithDefaults() *GetUserCategoryPercentagesRes`

NewGetUserCategoryPercentagesResWithDefaults instantiates a new GetUserCategoryPercentagesRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetUserCategoryPercentagesRes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetUserCategoryPercentagesRes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetUserCategoryPercentagesRes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCategories

`func (o *GetUserCategoryPercentagesRes) GetCategories() []CategoryInterest`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GetUserCategoryPercentagesRes) GetCategoriesOk() (*[]CategoryInterest, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GetUserCategoryPercentagesRes) SetCategories(v []CategoryInterest)`

SetCategories sets Categories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


