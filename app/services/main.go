package services

import "github.com/jinzhu/gorm"

type Managers struct {
	OauthClient *OauthClientManager
}

func InitManagers (db *gorm.DB) *Managers {
	oauthClientManager := NewOauthClientManager(db)

	return &Managers{
		OauthClient: oauthClientManager,
	}
}