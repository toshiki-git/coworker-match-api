package models

import "time"

type Message struct {
	MessageID      string       `json:"message_id"`       // メッセージID
	MatchingID     string       `json:"matching_id"`      // マッチングID
	UserID         string       `json:"user_id"`          // ユーザID
	QuestionCardID string       `json:"question_card_id"` // 質問カードID
	MessageText    string       `json:"message_text"`     // メッセージテキスト
	IsRead         bool         `json:"is_read"`          // 既読フラグ
	CreatedAt      time.Time    // 作成日時
	UpdatedAt      time.Time    // 更新日時
	Matching       Match        `json:"matching"`      // マッチング, MatchingIDとMatchの関連付け
	User           User         `json:"user"`          // ユーザ, UserIDとUserの関連付け
	QuestionCard   QuestionCard `json:"question_card"` // 質問カード, QuestionCardIDとQuestionCardの関連付け
}
