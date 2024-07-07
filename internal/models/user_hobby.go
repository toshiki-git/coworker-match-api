package models

import "time"

type UserHobby struct {
	HobbyID   string    `json:"hobby_id"` // 趣味ID
	UserID    string    `json:"user_id"`  // ユーザーID
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時
	Hobby     Hobby     `json:"hobby"` // 趣味, HobbyIDとHobbyの関連付け
}
