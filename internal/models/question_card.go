package models

import "time"

type QuestionCard struct {
	QuestionCardID   string    `json:"question_card_id"`   // 質問カードID
	QuestionCardText string    `json:"question_card_text"` // 質問カードテキスト
	CreatorID        string    `json:"creator_id"`         // 作成者ID
	CreatedAt        time.Time // 作成日時
	UpdatedAt        time.Time // 更新日時
	Creator          User      `json:"creator"` // 作成者, CreatorIDとUserの関連付け
}
