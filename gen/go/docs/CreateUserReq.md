# CreateUserReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** |  | 
**Email** | **string** |  | 
**AvatarUrl** | **string** |  | 

## Methods

### NewCreateUserReq

`func NewCreateUserReq(userName string, email string, avatarUrl string, ) *CreateUserReq`

NewCreateUserReq instantiates a new CreateUserReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserReqWithDefaults

`func NewCreateUserReqWithDefaults() *CreateUserReq`

NewCreateUserReqWithDefaults instantiates a new CreateUserReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *CreateUserReq) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateUserReq) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateUserReq) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetEmail

`func (o *CreateUserReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAvatarUrl

`func (o *CreateUserReq) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *CreateUserReq) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *CreateUserReq) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


