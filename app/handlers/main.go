package handlers

import (
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/jinzhu/gorm"
)

type Handlers struct{
	Oauth *Oauth
	Api *Api
}

type Oauth struct {
	AuthorizeClientHandler *AuthorizeClientHandler
	AccessTokenHandler *AccessTokenHandler
	ProfileHandler *ProfileHandler
}

type Api struct {
	ActivateHandler *ActivateHandler
	RegisterHandler *RegisterHandler
	ProfileHandler *ProfileHandler
}

type DefaultResponse struct {
	Status int
}

func InitHandlers(managers *services.Managers, merger *services.UserMerger, db *gorm.DB) *Handlers {
	return &Handlers{
		Oauth: &Oauth {
			AuthorizeClientHandler: NewAuthorizeClientHandler(managers.OauthClient, managers.User),
			AccessTokenHandler:     NewAccessTokenClientHandler(managers.OauthClient, managers.OauthSession),
		},
		Api: &Api {
			ActivateHandler: NewActivateHandler(
				managers.OauthSession,
				managers.OauthClient,
				managers.User,
				managers.Email,
				merger,
			),
			RegisterHandler: NewRegisterHandler(
				managers.OauthSession,
				managers.OauthClient,
				managers.User,
				managers.Email,
				merger,
				db,
			),
			ProfileHandler: NewProfileHandler(
				managers.OauthSession,
				managers.OauthClient,
				managers.User,
				managers.Email,
			),
		},
	}
}