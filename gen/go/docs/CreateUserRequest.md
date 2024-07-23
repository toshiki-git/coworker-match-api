# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**AvatarUrl** | **string** |  | 
**Age** | Pointer to **NullableInt32** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Birthplace** | Pointer to **NullableString** |  | [optional] 
**JobType** | Pointer to **NullableString** |  | [optional] 
**LineAccount** | Pointer to **NullableString** |  | [optional] 
**DiscordAccount** | Pointer to **NullableString** |  | [optional] 
**Biography** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(name string, email string, avatarUrl string, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAvatarUrl

`func (o *CreateUserRequest) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *CreateUserRequest) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *CreateUserRequest) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetAge

`func (o *CreateUserRequest) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *CreateUserRequest) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *CreateUserRequest) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *CreateUserRequest) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *CreateUserRequest) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *CreateUserRequest) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetGender

`func (o *CreateUserRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CreateUserRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CreateUserRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CreateUserRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CreateUserRequest) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CreateUserRequest) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetBirthplace

`func (o *CreateUserRequest) GetBirthplace() string`

GetBirthplace returns the Birthplace field if non-nil, zero value otherwise.

### GetBirthplaceOk

`func (o *CreateUserRequest) GetBirthplaceOk() (*string, bool)`

GetBirthplaceOk returns a tuple with the Birthplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthplace

`func (o *CreateUserRequest) SetBirthplace(v string)`

SetBirthplace sets Birthplace field to given value.

### HasBirthplace

`func (o *CreateUserRequest) HasBirthplace() bool`

HasBirthplace returns a boolean if a field has been set.

### SetBirthplaceNil

`func (o *CreateUserRequest) SetBirthplaceNil(b bool)`

 SetBirthplaceNil sets the value for Birthplace to be an explicit nil

### UnsetBirthplace
`func (o *CreateUserRequest) UnsetBirthplace()`

UnsetBirthplace ensures that no value is present for Birthplace, not even an explicit nil
### GetJobType

`func (o *CreateUserRequest) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *CreateUserRequest) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *CreateUserRequest) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *CreateUserRequest) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *CreateUserRequest) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *CreateUserRequest) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetLineAccount

`func (o *CreateUserRequest) GetLineAccount() string`

GetLineAccount returns the LineAccount field if non-nil, zero value otherwise.

### GetLineAccountOk

`func (o *CreateUserRequest) GetLineAccountOk() (*string, bool)`

GetLineAccountOk returns a tuple with the LineAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAccount

`func (o *CreateUserRequest) SetLineAccount(v string)`

SetLineAccount sets LineAccount field to given value.

### HasLineAccount

`func (o *CreateUserRequest) HasLineAccount() bool`

HasLineAccount returns a boolean if a field has been set.

### SetLineAccountNil

`func (o *CreateUserRequest) SetLineAccountNil(b bool)`

 SetLineAccountNil sets the value for LineAccount to be an explicit nil

### UnsetLineAccount
`func (o *CreateUserRequest) UnsetLineAccount()`

UnsetLineAccount ensures that no value is present for LineAccount, not even an explicit nil
### GetDiscordAccount

`func (o *CreateUserRequest) GetDiscordAccount() string`

GetDiscordAccount returns the DiscordAccount field if non-nil, zero value otherwise.

### GetDiscordAccountOk

`func (o *CreateUserRequest) GetDiscordAccountOk() (*string, bool)`

GetDiscordAccountOk returns a tuple with the DiscordAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordAccount

`func (o *CreateUserRequest) SetDiscordAccount(v string)`

SetDiscordAccount sets DiscordAccount field to given value.

### HasDiscordAccount

`func (o *CreateUserRequest) HasDiscordAccount() bool`

HasDiscordAccount returns a boolean if a field has been set.

### SetDiscordAccountNil

`func (o *CreateUserRequest) SetDiscordAccountNil(b bool)`

 SetDiscordAccountNil sets the value for DiscordAccount to be an explicit nil

### UnsetDiscordAccount
`func (o *CreateUserRequest) UnsetDiscordAccount()`

UnsetDiscordAccount ensures that no value is present for DiscordAccount, not even an explicit nil
### GetBiography

`func (o *CreateUserRequest) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *CreateUserRequest) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *CreateUserRequest) SetBiography(v string)`

SetBiography sets Biography field to given value.

### HasBiography

`func (o *CreateUserRequest) HasBiography() bool`

HasBiography returns a boolean if a field has been set.

### SetBiographyNil

`func (o *CreateUserRequest) SetBiographyNil(b bool)`

 SetBiographyNil sets the value for Biography to be an explicit nil

### UnsetBiography
`func (o *CreateUserRequest) UnsetBiography()`

UnsetBiography ensures that no value is present for Biography, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


