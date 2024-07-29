# CreateQuestionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchingId** | **string** |  | 
**SenderUserId** | **string** |  | 
**ReceiverUserId** | **string** |  | 
**MatchingDate** | **time.Time** |  | 

## Methods

### NewCreateQuestionResponse

`func NewCreateQuestionResponse(matchingId string, senderUserId string, receiverUserId string, matchingDate time.Time, ) *CreateQuestionResponse`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


