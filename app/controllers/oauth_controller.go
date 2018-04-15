package controllers

import (
	"github.com/labstack/echo"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/validators"
	"net/http"
	"github.com/MetalRex101/auth-server/app/repositories"
)

type Response struct{}

type oauthController struct{}

func (controller *oauthController) AuthorizeClient(c echo.Context) error {
	oauthClientManager := services.OauthClientManager{}

	clientID, err := validators.Request.GetClientId(true, c)
	if err != nil {
		return err
	}

	client, err := repositories.OauthClient.GetForOauth(clientID)
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

	user, err := repositories.User.GetByEmailAndPassword(email, password, true)
	if err != nil {
		return err
	}

	if err := services.Auth.LogoutCurrentUser(c); err != nil {
		return err
	}

	if err := services.Auth.AuthenticateUser(user, c); err != nil {
		return err
	}

	oauthSession, err := oauthClientManager.StartSession(client, c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, oauthSession)
}
