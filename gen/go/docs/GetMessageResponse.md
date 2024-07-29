# GetMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]GetMessageResponseMessagesInner**](GetMessageResponseMessagesInner.md) |  | 

## Methods

### NewGetMessageResponse

`func NewGetMessageResponse(messages []GetMessageResponseMessagesInner, ) *GetMessageResponse`

NewGetMessageResponse instantiates a new GetMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessageResponseWithDefaults

`func NewGetMessageResponseWithDefaults() *GetMessageResponse`

NewGetMessageResponseWithDefaults instantiates a new GetMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *GetMessageResponse) GetMessages() []GetMessageResponseMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetMessageResponse) GetMessagesOk() (*[]GetMessageResponseMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetMessageResponse) SetMessages(v []GetMessageResponseMessagesInner)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


