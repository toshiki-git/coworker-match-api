# CreateQuestionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchingId** | Pointer to **string** |  | [optional] 
**SenderUserId** | Pointer to **string** |  | [optional] 
**ReceiverUserId** | Pointer to **string** |  | [optional] 
**MatchingDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateQuestionResponse

`func NewCreateQuestionResponse() *CreateQuestionResponse`

NewCreateQuestionResponse instantiates a new CreateQuestionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuestionResponseWithDefaults

`func NewCreateQuestionResponseWithDefaults() *CreateQuestionResponse`

NewCreateQuestionResponseWithDefaults instantiates a new CreateQuestionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchingId

`func (o *CreateQuestionResponse) GetMatchingId() string`

GetMatchingId returns the MatchingId field if non-nil, zero value otherwise.

### GetMatchingIdOk

`func (o *CreateQuestionResponse) GetMatchingIdOk() (*string, bool)`

GetMatchingIdOk returns a tuple with the MatchingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingId

`func (o *CreateQuestionResponse) SetMatchingId(v string)`

SetMatchingId sets MatchingId field to given value.

### HasMatchingId

`func (o *CreateQuestionResponse) HasMatchingId() bool`

HasMatchingId returns a boolean if a field has been set.

### GetSenderUserId

`func (o *CreateQuestionResponse) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *CreateQuestionResponse) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *CreateQuestionResponse) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.

### HasSenderUserId

`func (o *CreateQuestionResponse) HasSenderUserId() bool`

HasSenderUserId returns a boolean if a field has been set.

### GetReceiverUserId

`func (o *CreateQuestionResponse) GetReceiverUserId() string`

GetReceiverUserId returns the ReceiverUserId field if non-nil, zero value otherwise.

### GetReceiverUserIdOk

`func (o *CreateQuestionResponse) GetReceiverUserIdOk() (*string, bool)`

GetReceiverUserIdOk returns a tuple with the ReceiverUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverUserId

`func (o *CreateQuestionResponse) SetReceiverUserId(v string)`

SetReceiverUserId sets ReceiverUserId field to given value.

### HasReceiverUserId

`func (o *CreateQuestionResponse) HasReceiverUserId() bool`

HasReceiverUserId returns a boolean if a field has been set.

### GetMatchingDate

`func (o *CreateQuestionResponse) GetMatchingDate() time.Time`

GetMatchingDate returns the MatchingDate field if non-nil, zero value otherwise.

### GetMatchingDateOk

`func (o *CreateQuestionResponse) GetMatchingDateOk() (*time.Time, bool)`

GetMatchingDateOk returns a tuple with the MatchingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingDate

`func (o *CreateQuestionResponse) SetMatchingDate(v time.Time)`

SetMatchingDate sets MatchingDate field to given value.

### HasMatchingDate

`func (o *CreateQuestionResponse) HasMatchingDate() bool`

HasMatchingDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


