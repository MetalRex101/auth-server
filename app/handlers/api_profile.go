package handlers

import (
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/labstack/echo"
	"github.com/MetalRex101/auth-server/app/validators"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/MetalRex101/auth-server/app/resources"
	"net/http"
)

type ProfileHandler struct{
	osm  services.IOauthSessionManager
	ocm  services.IOauthClientManager
	um   services.IUserManager
	em   services.IEmailManager
}

func NewProfileHandler(
	osm services.IOauthSessionManager,
	ocm services.IOauthClientManager,
	um services.IUserManager,
	em services.IEmailManager,
) *ProfileHandler {
	return &ProfileHandler{osm, ocm, um, em}
}

func (pf *ProfileHandler) Handle (c echo.Context) error {
	var user *models.User

	accessToken, err := validators.Request.GetAccessToken(false, c)
	if err != nil {
		return err
	}

	oauthSess, err := pf.osm.FindByToken(accessToken)
	if err != nil {
		return err
	}

	user, err = pf.um.GetUserFromSession(oauthSess)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resources.NewRegisteredResource(user, pf.um.(*services.UserManager), c))
}