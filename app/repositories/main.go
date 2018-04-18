package repositories

import "github.com/jinzhu/gorm"

type Repositories struct {
	User *UserRepo
	OauthClient *OauthClientRepo
	OauthSession *OauthSessionRepo
}

func InitRepositories (db *gorm.DB) *Repositories {
	return &Repositories{
		User:NewUserRepo(db),
		OauthClient:NewOauthClientRepo(db),
		OauthSession:NewOauthSessionRepo(db),
	}
}
