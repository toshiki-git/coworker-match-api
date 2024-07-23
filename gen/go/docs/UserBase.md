# UserBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Age** | Pointer to **NullableInt32** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Birthplace** | Pointer to **NullableString** |  | [optional] 
**JobType** | Pointer to **NullableString** |  | [optional] 
**LineAccount** | Pointer to **NullableString** |  | [optional] 
**DiscordAccount** | Pointer to **NullableString** |  | [optional] 
**Biography** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserBase

`func NewUserBase() *UserBase`

NewUserBase instantiates a new UserBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBaseWithDefaults

`func NewUserBaseWithDefaults() *UserBase`

NewUserBaseWithDefaults instantiates a new UserBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UserBase) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserBase) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserBase) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserBase) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserBase) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserBase) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserBase) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserBase) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetAge

`func (o *UserBase) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *UserBase) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *UserBase) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *UserBase) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *UserBase) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *UserBase) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetGender

`func (o *UserBase) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UserBase) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UserBase) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UserBase) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *UserBase) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *UserBase) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetBirthplace

`func (o *UserBase) GetBirthplace() string`

GetBirthplace returns the Birthplace field if non-nil, zero value otherwise.

### GetBirthplaceOk

`func (o *UserBase) GetBirthplaceOk() (*string, bool)`

GetBirthplaceOk returns a tuple with the Birthplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthplace

`func (o *UserBase) SetBirthplace(v string)`

SetBirthplace sets Birthplace field to given value.

### HasBirthplace

`func (o *UserBase) HasBirthplace() bool`

HasBirthplace returns a boolean if a field has been set.

### SetBirthplaceNil

`func (o *UserBase) SetBirthplaceNil(b bool)`

 SetBirthplaceNil sets the value for Birthplace to be an explicit nil

### UnsetBirthplace
`func (o *UserBase) UnsetBirthplace()`

UnsetBirthplace ensures that no value is present for Birthplace, not even an explicit nil
### GetJobType

`func (o *UserBase) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *UserBase) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *UserBase) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *UserBase) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *UserBase) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *UserBase) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetLineAccount

`func (o *UserBase) GetLineAccount() string`

GetLineAccount returns the LineAccount field if non-nil, zero value otherwise.

### GetLineAccountOk

`func (o *UserBase) GetLineAccountOk() (*string, bool)`

GetLineAccountOk returns a tuple with the LineAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAccount

`func (o *UserBase) SetLineAccount(v string)`

SetLineAccount sets LineAccount field to given value.

### HasLineAccount

`func (o *UserBase) HasLineAccount() bool`

HasLineAccount returns a boolean if a field has been set.

### SetLineAccountNil

`func (o *UserBase) SetLineAccountNil(b bool)`

 SetLineAccountNil sets the value for LineAccount to be an explicit nil

### UnsetLineAccount
`func (o *UserBase) UnsetLineAccount()`

UnsetLineAccount ensures that no value is present for LineAccount, not even an explicit nil
### GetDiscordAccount

`func (o *UserBase) GetDiscordAccount() string`

GetDiscordAccount returns the DiscordAccount field if non-nil, zero value otherwise.

### GetDiscordAccountOk

`func (o *UserBase) GetDiscordAccountOk() (*string, bool)`

GetDiscordAccountOk returns a tuple with the DiscordAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordAccount

`func (o *UserBase) SetDiscordAccount(v string)`

SetDiscordAccount sets DiscordAccount field to given value.

### HasDiscordAccount

`func (o *UserBase) HasDiscordAccount() bool`

HasDiscordAccount returns a boolean if a field has been set.

### SetDiscordAccountNil

`func (o *UserBase) SetDiscordAccountNil(b bool)`

 SetDiscordAccountNil sets the value for DiscordAccount to be an explicit nil

### UnsetDiscordAccount
`func (o *UserBase) UnsetDiscordAccount()`

UnsetDiscordAccount ensures that no value is present for DiscordAccount, not even an explicit nil
### GetBiography

`func (o *UserBase) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *UserBase) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *UserBase) SetBiography(v string)`

SetBiography sets Biography field to given value.

### HasBiography

`func (o *UserBase) HasBiography() bool`

HasBiography returns a boolean if a field has been set.

### SetBiographyNil

`func (o *UserBase) SetBiographyNil(b bool)`

 SetBiographyNil sets the value for Biography to be an explicit nil

### UnsetBiography
`func (o *UserBase) UnsetBiography()`

UnsetBiography ensures that no value is present for Biography, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


