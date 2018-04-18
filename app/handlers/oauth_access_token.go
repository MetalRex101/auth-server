package handlers

import (
	"github.com/MetalRex101/auth-server/app/validators"
	"github.com/MetalRex101/auth-server/app/resources"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/labstack/echo"
	"net/http"
)

type AccessTokenHandler struct {
	oauthClientManager services.IOauthClientManager
	oauthSessionManager services.IOauthSessionManager
}

func NewAccessTokenClientHandler(
	ocm services.IOauthClientManager,
	osm services.IOauthSessionManager,
) *AccessTokenHandler {
	return &AccessTokenHandler{ocm, osm}
}

func (ac *AccessTokenHandler) Handle (c echo.Context) error {
	clientID, err := validators.Request.GetClientId(true, c)
	if err != nil {
		return err
	}

	clientSecret, err := validators.Request.GetClientSecret(true, c)
	if err != nil {
		return err
	}

	client, err := ac.oauthClientManager.GetForApi(clientID, clientSecret, c.Request().RemoteAddr)
	if err != nil {
		return err
	}

	if err = validators.Client.HasScope(client, []string{"oauth"}); err != nil {
		return err
	}

	code, err := validators.Request.GetCode(validators.CodeTypeAccess, c)
	if err != nil {
		return err
	}

	oauthSession, err := ac.oauthSessionManager.FindByClientIDAndCode(clientID, code)
	if err != nil {
		return err
	}

	ac.oauthSessionManager.StartSession(oauthSession, false, 0, c)

	return c.JSON(http.StatusOK, resources.NewAccessTokenResource(oauthSession))
}
