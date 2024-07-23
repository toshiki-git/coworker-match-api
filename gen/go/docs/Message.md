# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** |  | [optional] 
**MessageText** | Pointer to **string** |  | [optional] 

## Methods

### NewMessage

`func NewMessage() *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *Message) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Message) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Message) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *Message) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageText

`func (o *Message) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *Message) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *Message) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *Message) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


