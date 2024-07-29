# Question

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionId** | **string** |  | 
**QuestionText** | **string** |  | 
**Choice1** | [**Choice**](Choice.md) |  | 
**Choice2** | [**Choice**](Choice.md) |  | 

## Methods

### NewQuestion

`func NewQuestion(questionId string, questionText string, choice1 Choice, choice2 Choice, ) *Question`

NewQuestion instantiates a new Question object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionWithDefaults

`func NewQuestionWithDefaults() *Question`

NewQuestionWithDefaults instantiates a new Question object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionId

`func (o *Question) GetQuestionId() string`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *Question) GetQuestionIdOk() (*string, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *Question) SetQuestionId(v string)`

SetQuestionId sets QuestionId field to given value.


### GetQuestionText

`func (o *Question) GetQuestionText() string`

GetQuestionText returns the QuestionText field if non-nil, zero value otherwise.

### GetQuestionTextOk

`func (o *Question) GetQuestionTextOk() (*string, bool)`

GetQuestionTextOk returns a tuple with the QuestionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionText

`func (o *Question) SetQuestionText(v string)`

SetQuestionText sets QuestionText field to given value.


### GetChoice1

`func (o *Question) GetChoice1() Choice`

GetChoice1 returns the Choice1 field if non-nil, zero value otherwise.

### GetChoice1Ok

`func (o *Question) GetChoice1Ok() (*Choice, bool)`

GetChoice1Ok returns a tuple with the Choice1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice1

`func (o *Question) SetChoice1(v Choice)`

SetChoice1 sets Choice1 field to given value.


### GetChoice2

`func (o *Question) GetChoice2() Choice`

GetChoice2 returns the Choice2 field if non-nil, zero value otherwise.

### GetChoice2Ok

`func (o *Question) GetChoice2Ok() (*Choice, bool)`

GetChoice2Ok returns a tuple with the Choice2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice2

`func (o *Question) SetChoice2(v Choice)`

SetChoice2 sets Choice2 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


