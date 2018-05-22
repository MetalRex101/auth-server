package handlers

import (
	"github.com/MetalRex101/auth-server/app/services"
)

type Handlers struct{
	Oauth *Oauth
	Api *Api
}

type Oauth struct {
	AuthorizeClientHandler *AuthorizeClientHandler
	AccessTokenHandler *AccessTokenHandler
}

type Api struct {
	ActivateHandler *ActivateHandler
}

func InitHandlers(managers *services.Managers) *Handlers {
	return &Handlers{
		Oauth: &Oauth {
			AuthorizeClientHandler: NewAuthorizeClientHandler(managers.OauthClient, managers.User),
			AccessTokenHandler: NewAccessTokenClientHandler(managers.OauthClient, managers.OauthSession),
		},
		Api: &Api {

		},
	}
}