package models

import (
	"time"
)

type Phone struct {
	ID          uint
	UserID      *uint
	Oauth       *string    `gorm:"size:255"`
	Phone       *uint64    `gorm:"size:100;not null"`
	Status      bool       `gorm:"default:false"`
	IsDefault   bool       `gorm:"default:false"`
	ConfirmDate *time.Time
	Code        *string    `gorm:"size:90"`
	CreatedAt   *time.Time `gorm:"not null"`
}

func (Phone) TableName() string {
	return "phones"
}
