# CreateHobbyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hobby** | [**HobbyDetails**](HobbyDetails.md) |  | 

## Methods

### NewCreateHobbyResponse

`func NewCreateHobbyResponse(hobby HobbyDetails, ) *CreateHobbyResponse`

NewCreateHobbyResponse instantiates a new CreateHobbyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHobbyResponseWithDefaults

`func NewCreateHobbyResponseWithDefaults() *CreateHobbyResponse`

NewCreateHobbyResponseWithDefaults instantiates a new CreateHobbyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHobby

`func (o *CreateHobbyResponse) GetHobby() HobbyDetails`

GetHobby returns the Hobby field if non-nil, zero value otherwise.

### GetHobbyOk

`func (o *CreateHobbyResponse) GetHobbyOk() (*HobbyDetails, bool)`

GetHobbyOk returns a tuple with the Hobby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHobby

`func (o *CreateHobbyResponse) SetHobby(v HobbyDetails)`

SetHobby sets Hobby field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


