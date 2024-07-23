# CreateQuestionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | Pointer to [**[]Answer**](Answer.md) |  | [optional] 

## Methods

### NewCreateQuestionRequest

`func NewCreateQuestionRequest() *CreateQuestionRequest`

NewCreateQuestionRequest instantiates a new CreateQuestionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuestionRequestWithDefaults

`func NewCreateQuestionRequestWithDefaults() *CreateQuestionRequest`

NewCreateQuestionRequestWithDefaults instantiates a new CreateQuestionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *CreateQuestionRequest) GetAnswers() []Answer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *CreateQuestionRequest) GetAnswersOk() (*[]Answer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *CreateQuestionRequest) SetAnswers(v []Answer)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *CreateQuestionRequest) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


