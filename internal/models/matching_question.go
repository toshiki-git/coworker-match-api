package models

type MatchingQuestion struct {
	QuestionID         string   `json:"question_id"`          // 質問ID
	QuestionText       string   `json:"question_text"`        // 質問文
	QuestionCategoryID string   `json:"question_category_id"` // カテゴリID(future use)
	QuestionCategory   Category `json:"category"`             // カテゴリ, QuestionCategoryIDとCategoryの関連付け
}
