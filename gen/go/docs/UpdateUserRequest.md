# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *UpdateUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UpdateUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UpdateUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UpdateUserRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UpdateUserRequest) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UpdateUserRequest) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UpdateUserRequest) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UpdateUserRequest) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


