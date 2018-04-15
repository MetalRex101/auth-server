package models

import "time"

type Password struct{
	ID uint
	UserID *uint
	Password *string `gorm:"size:32;not null"`
	CreatedAt *time.Time `gorm:"not null"`
	Code *string `gorm:"size:100"`

	User *User
}

func (Password) TableName() string {
	return "passwords"
}
