package controllers

import (
	"github.com/MetalRex101/auth-server/app/repositories"
	"github.com/MetalRex101/auth-server/app/services"
)

type Controllers struct{
	Oauth *OauthController
}

type BaseController struct {
	Repos *repositories.Repositories
	Managers *services.Managers
}

func InitControllers (repos *repositories.Repositories, managers *services.Managers) *Controllers {
	base := NewBaseController(repos, managers)

	return &Controllers{
		Oauth:NewOauthController(base),
	}
}

func NewBaseController (repos *repositories.Repositories, managers *services.Managers) *BaseController {
	return &BaseController{repos, managers}
}