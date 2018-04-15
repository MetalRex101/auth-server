package services

import "github.com/jinzhu/gorm"

type Managers struct {
	OauthClient *OauthClientManager
}

func InitManagers (db *gorm.DB) *Managers {
	return &Managers{
		OauthClient: NewOauthClientManager(db),
	}
}