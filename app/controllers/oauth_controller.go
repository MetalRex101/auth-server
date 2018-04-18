package controllers

import (
	"github.com/labstack/echo"
	"github.com/MetalRex101/auth-server/app/validators"
	"net/http"
	"github.com/MetalRex101/auth-server/app/resources"
)

type Response struct{}

type OauthController struct{
	Base *BaseController
}

func NewOauthController (base *BaseController) *OauthController {
	return &OauthController{base}
}

func (oc *OauthController) AuthorizeClient(c echo.Context) error {
	clientID, err := validators.Request.GetClientId(true, c)
	if err != nil {
		return err
	}

	client, err := oc.Base.Repos.OauthClient.GetForOauth(clientID)
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

	user, err := oc.Base.Repos.User.GetByEmailAndPassword(email, password, true)
	if err != nil {
		return err
	}

	oauthSession, err := oc.Base.Managers.OauthClient.StartSession(client, user, c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, oauthSession)
}

func (oc *OauthController) GetAccessToken (c echo.Context) error {
	clientID, err := validators.Request.GetClientId(true, c)
	if err != nil {
		return err
	}

	clientSecret, err := validators.Request.GetClientSecret(true, c)
	if err != nil {
		return err
	}

	client, err := oc.Base.Repos.OauthClient.GetForApi(clientID, clientSecret, c.Request().RemoteAddr)
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

	oauthSession, err :=oc.Base.Repos.OauthSession.Find(clientID, code)
	if err != nil {
		return err
	}

	oc.Base.Managers.OauthSession.StartSession(oauthSession, false, 0, c)

	return c.JSON(http.StatusOK, resources.NewAccessTokenResource(oauthSession))
}