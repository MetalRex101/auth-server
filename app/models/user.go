package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID         uint
	LastVisit  *time.Time
	CreatorID  *uint
	EditorID   *uint
	Status     bool       `gorm:"default:false"`
	Merged     *string
	Nickname   *string    `gorm:"size:50"`
	FirstName  *string    `gorm:"size:100"`
	FatherName *string    `gorm:"size:100"`
	LastName   *string    `gorm:"size:100"`
	BirthDate  *time.Time
	Gender     *string
	Language   *string
	CreatedAt  *time.Time `gorm:"not null"`
	UpdatedAt  *time.Time `gorm:"not null"`
	ChangedAt  *time.Time

	Passwords []Password
	Emails    []Email
	Phones    []Phone
	Editor    *User
	Creator   *User
	Roles     []Role `gorm:"many2many:user_role;"`
}

func (User) TableName() string {
	return "users"
}

func (User) WhereHasPassword(pass string) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("exists(?)", db.
			Table("passwords").
			Select("1").
			Where("password = ? and user_id = users.id", pass).
			QueryExpr())
	}
}

func (User) WhereHasEmail(email string) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("exists(?)", db.
			Table("emails").
			Select("1").
			Where("email = ? and user_id = users.id", email).
			QueryExpr())
	}
}
