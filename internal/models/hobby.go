package models

import (
	"time"
)

type Hobby struct {
	HobbyID    string    `json:"hobby_id"`    // 趣味ID
	HobbyName  string    `json:"hobby_name"`  // 趣味名
	CategoryID string    `json:"category_id"` // カテゴリID
	CreatorID  string    `json:"creator_id"`  // 作成者ID
	CreatedAt  time.Time // 作成日時
	UpdatedAt  time.Time // 更新日時
	Category   Category  `json:"category"` // カテゴリ, CategoryIDとCategoryの関連付け
	User       User      `json:"user"`     // 作成者, CreatorIDとUserの関連付け
}
