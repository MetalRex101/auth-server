package repositories

import "github.com/jinzhu/gorm"

type Repositories struct {
	User *UserRepo
	OauthClient *OauthClientRepo
}

func InitRepositories (db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	oauthClientRepo := NewOauthClientRepo(db)

	return &Repositories{
		User:userRepo,
		OauthClient:oauthClientRepo,
	}
}
