# GetMessageResponseMessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionCardId** | Pointer to **string** |  | [optional] 
**QuestionCardText** | Pointer to **string** |  | [optional] 
**MessagePair** | Pointer to [**GetMessageResponseMessagesInnerMessagePair**](GetMessageResponseMessagesInnerMessagePair.md) |  | [optional] 

## Methods

### NewGetMessageResponseMessagesInner

`func NewGetMessageResponseMessagesInner() *GetMessageResponseMessagesInner`

NewGetMessageResponseMessagesInner instantiates a new GetMessageResponseMessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessageResponseMessagesInnerWithDefaults

`func NewGetMessageResponseMessagesInnerWithDefaults() *GetMessageResponseMessagesInner`

NewGetMessageResponseMessagesInnerWithDefaults instantiates a new GetMessageResponseMessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionCardId

`func (o *GetMessageResponseMessagesInner) GetQuestionCardId() string`

GetQuestionCardId returns the QuestionCardId field if non-nil, zero value otherwise.

### GetQuestionCardIdOk

`func (o *GetMessageResponseMessagesInner) GetQuestionCardIdOk() (*string, bool)`

GetQuestionCardIdOk returns a tuple with the QuestionCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionCardId

`func (o *GetMessageResponseMessagesInner) SetQuestionCardId(v string)`

SetQuestionCardId sets QuestionCardId field to given value.

### HasQuestionCardId

`func (o *GetMessageResponseMessagesInner) HasQuestionCardId() bool`

HasQuestionCardId returns a boolean if a field has been set.

### GetQuestionCardText

`func (o *GetMessageResponseMessagesInner) GetQuestionCardText() string`

GetQuestionCardText returns the QuestionCardText field if non-nil, zero value otherwise.

### GetQuestionCardTextOk

`func (o *GetMessageResponseMessagesInner) GetQuestionCardTextOk() (*string, bool)`

GetQuestionCardTextOk returns a tuple with the QuestionCardText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionCardText

`func (o *GetMessageResponseMessagesInner) SetQuestionCardText(v string)`

SetQuestionCardText sets QuestionCardText field to given value.

### HasQuestionCardText

`func (o *GetMessageResponseMessagesInner) HasQuestionCardText() bool`

HasQuestionCardText returns a boolean if a field has been set.

### GetMessagePair

`func (o *GetMessageResponseMessagesInner) GetMessagePair() GetMessageResponseMessagesInnerMessagePair`

GetMessagePair returns the MessagePair field if non-nil, zero value otherwise.

### GetMessagePairOk

`func (o *GetMessageResponseMessagesInner) GetMessagePairOk() (*GetMessageResponseMessagesInnerMessagePair, bool)`

GetMessagePairOk returns a tuple with the MessagePair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagePair

`func (o *GetMessageResponseMessagesInner) SetMessagePair(v GetMessageResponseMessagesInnerMessagePair)`

SetMessagePair sets MessagePair field to given value.

### HasMessagePair

`func (o *GetMessageResponseMessagesInner) HasMessagePair() bool`

HasMessagePair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


