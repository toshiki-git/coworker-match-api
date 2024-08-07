# GetMessageResMessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionCardId** | **string** |  | 
**QuestionCardText** | **string** |  | 
**MessagePair** | [**GetMessageResMessagesInnerMessagePair**](GetMessageResMessagesInnerMessagePair.md) |  | 

## Methods

### NewGetMessageResMessagesInner

`func NewGetMessageResMessagesInner(questionCardId string, questionCardText string, messagePair GetMessageResMessagesInnerMessagePair, ) *GetMessageResMessagesInner`

NewGetMessageResMessagesInner instantiates a new GetMessageResMessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessageResMessagesInnerWithDefaults

`func NewGetMessageResMessagesInnerWithDefaults() *GetMessageResMessagesInner`

NewGetMessageResMessagesInnerWithDefaults instantiates a new GetMessageResMessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionCardId

`func (o *GetMessageResMessagesInner) GetQuestionCardId() string`

GetQuestionCardId returns the QuestionCardId field if non-nil, zero value otherwise.

### GetQuestionCardIdOk

`func (o *GetMessageResMessagesInner) GetQuestionCardIdOk() (*string, bool)`

GetQuestionCardIdOk returns a tuple with the QuestionCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionCardId

`func (o *GetMessageResMessagesInner) SetQuestionCardId(v string)`

SetQuestionCardId sets QuestionCardId field to given value.


### GetQuestionCardText

`func (o *GetMessageResMessagesInner) GetQuestionCardText() string`

GetQuestionCardText returns the QuestionCardText field if non-nil, zero value otherwise.

### GetQuestionCardTextOk

`func (o *GetMessageResMessagesInner) GetQuestionCardTextOk() (*string, bool)`

GetQuestionCardTextOk returns a tuple with the QuestionCardText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionCardText

`func (o *GetMessageResMessagesInner) SetQuestionCardText(v string)`

SetQuestionCardText sets QuestionCardText field to given value.


### GetMessagePair

`func (o *GetMessageResMessagesInner) GetMessagePair() GetMessageResMessagesInnerMessagePair`

GetMessagePair returns the MessagePair field if non-nil, zero value otherwise.

### GetMessagePairOk

`func (o *GetMessageResMessagesInner) GetMessagePairOk() (*GetMessageResMessagesInnerMessagePair, bool)`

GetMessagePairOk returns a tuple with the MessagePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagePair

`func (o *GetMessageResMessagesInner) SetMessagePair(v GetMessageResMessagesInnerMessagePair)`

SetMessagePair sets MessagePair field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


