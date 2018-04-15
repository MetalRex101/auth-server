package controllers

import (
	"github.com/MetalRex101/auth-server/app/repositories"
	"github.com/MetalRex101/auth-server/app/services"
)

type BaseController struct {
	Repos *repositories.Repositories
	Managers *services.Managers
}

func NewBaseController (repos *repositories.Repositories, managers *services.Managers) *BaseController {
	return &BaseController{repos, managers}
}