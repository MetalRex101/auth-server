package handlers

import (
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/validators"
	"github.com/labstack/echo"
	"net/http"
)

type AuthorizeClientHandler struct{
	userManager        services.IUserManager
	oauthClientManager services.IOauthClientManager
}

func NewAuthorizeClientHandler(
	ocm services.IOauthClientManager,
	um services.IUserManager,
) *AuthorizeClientHandler {
	return &AuthorizeClientHandler{um, ocm}
}

func (ac *AuthorizeClientHandler) Handle(c echo.Context) error {
	clientID, err := validators.Request.GetClientId(true, c)
	if err != nil {
		return err
	}

	client, err := ac.oauthClientManager.GetForOauth(clientID)
	if err != nil {
		return err
	}

	if err = validators.Client.HasScope(client, []string{"oauth"}); err != nil {
		return err
	}

	url, err := validators.Request.GetRedirectUri(false, c)
	if err != nil {
		return err
	}

	if err = validators.Client.IsClientUrl(client, url, c); err != nil {
		return err
	}

	password, err := validators.Request.GetPassword(true, c, false)
	if err != nil {
		return err
	}

	email, err := validators.Request.GetEmail(true, c, false)
	if err != nil {
		return err
	}

	user, err := ac.userManager.GetByEmailAndPassword(email, password, true)
	if err != nil {
		return err
	}

	oauthSession, err := ac.oauthClientManager.StartSession(client, user, c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, oauthSession)
}
