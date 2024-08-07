# CreateQuestionReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | [**[]Answer**](Answer.md) |  | 

## Methods

### NewCreateQuestionReq

`func NewCreateQuestionReq(answers []Answer, ) *CreateQuestionReq`

NewCreateQuestionReq instantiates a new CreateQuestionReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuestionReqWithDefaults

`func NewCreateQuestionReqWithDefaults() *CreateQuestionReq`

NewCreateQuestionReqWithDefaults instantiates a new CreateQuestionReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *CreateQuestionReq) GetAnswers() []Answer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *CreateQuestionReq) GetAnswersOk() (*[]Answer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *CreateQuestionReq) SetAnswers(v []Answer)`

SetAnswers sets Answers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


