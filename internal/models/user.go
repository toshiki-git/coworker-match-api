package models

import (
	"time"
)

type GenderStatus string

const (
	Male   GenderStatus = "male"
	Female GenderStatus = "female"
	Other  GenderStatus = "other"
)

type User struct {
	UserID         string       `json:"user_id"`         // ユーザーID
	UserName       string       `json:"user_name"`       // ユーザー名
	Email          string       `json:"email"`           // メールアドレス
	AvatarURL      string       `json:"avatar_url"`      // アバターURL
	Age            int          `json:"age"`             // 年齢
	Gender         GenderStatus `json:"gender"`          // 性別
	Birthplace     string       `json:"birthplace"`      // 出身地
	JobType        string       `json:"job_type"`        // 職種
	LINEAccount    string       `json:"line_account"`    // LINEアカウント
	DiscordAccount string       `json:"discord_account"` // Discordアカウント
	Biography      string       `json:"biography"`       // 自己紹介
	CreatedAt      time.Time    // 作成日時
	UpdatedAt      time.Time    // 更新日時
}

func (g GenderStatus) IsValid() bool {
	switch g {
	case Male, Female, Other:
		return true
	}
	return false
}
