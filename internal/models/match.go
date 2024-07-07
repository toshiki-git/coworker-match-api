package models

import (
	"time"
)

type Match struct {
	MatchingID     string      `json:"matching_id"`      // マッチID
	SenderUserID   string      `json:"sender_user_id"`   // 送信者のユーザーID
	ReceiverUserID string      `json:"receiver_user_id"` // 受信者のユーザーID
	Status         MatchStatus `json:"status"`           // マッチステータス
	CreatedAt      time.Time   // 作成日時
	UpdatedAt      time.Time   // 更新日時
	SenderUser     User        `json:"sender_user"`   // 送信者のユーザー情報
	ReceiverUser   User        `json:"receiver_user"` // 受信者のユーザー情報
}

type MatchStatus string

const (
	MatchStatusPending  MatchStatus = "pending"  // マッチ待ち
	MatchStatusAccepted MatchStatus = "accepted" // マッチ成立
	MatchStatusRejected MatchStatus = "rejected" // マッチ拒否
)
