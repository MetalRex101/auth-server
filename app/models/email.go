package models

import (
	"time"
)

type Email struct {
	ID uint
	UserID *uint
	Oauth *string `gorm:"size:255"`
	Email *string `gorm:"size:100;not null"`
	Status bool `gorm:"default:false"`
	IsDefault bool `gorm:"default:false"`
	ConfirmDate *time.Time
	Code *string `gorm:"size:90"`
	CreatedAt *time.Time `gorm:"not null"`
}

func (Email) TableName() string {
	return "emails"
}