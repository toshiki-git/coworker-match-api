# GetQuestionRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Questions** | [**[]Question**](Question.md) |  | 

## Methods

### NewGetQuestionRes

`func NewGetQuestionRes(questions []Question, ) *GetQuestionRes`

NewGetQuestionRes instantiates a new GetQuestionRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuestionResWithDefaults

`func NewGetQuestionResWithDefaults() *GetQuestionRes`

NewGetQuestionResWithDefaults instantiates a new GetQuestionRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestions

`func (o *GetQuestionRes) GetQuestions() []Question`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *GetQuestionRes) GetQuestionsOk() (*[]Question, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *GetQuestionRes) SetQuestions(v []Question)`

SetQuestions sets Questions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


