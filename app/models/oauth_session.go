package models

import (
	"time"
)

type OauthSession struct {
	ID uint `gorm:"AUTO_INCREMENT"json:"-"`
	ClientID *uint `gorm:"not null"json:"-"`
	UserID *uint `gorm:"not null"json:"-"`
	AccessGrantedAt *time.Time `gorm:"not null"json:"-"`
	AccessExpiresAt *time.Time `gorm:"not null"json:"-"`
	Offset int `gorm:"default:0;not null"json:"-"`
	Code *string `json:"code"`
	AccessToken *string `gorm:"size:100"json:"-"`
	UserAgent *string `gorm:"type:text"json:"-"`
	RemoteAddr *string `gorm:"size:25"json:"-"`
	HttpReferer *string `json:"-"`
}

func (OauthSession) TableName() string {
	return "oauth_sessions"
}
