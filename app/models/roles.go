package models

import "time"

type Role struct {
	ID          uint
	Code        string     `gorm:"size:90;not null"`
	Name        string     `gorm:"size:50;not null"`
	Description *string    `gorm:"size:255"`
	Synthetic   bool       `gorm:"not null;default:false"`
	CreatedAt   *time.Time `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"not null"`
	creatorID   *uint
	editorID    *uint
	Status      bool       `gorm:"default:false"`
	Users       []User     `gorm:"many2many:user_role;"`
}

func (Role) TableName() string {
	return "roles"
}