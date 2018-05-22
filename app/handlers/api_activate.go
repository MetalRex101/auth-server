package handlers

import (
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/labstack/echo"
	"github.com/MetalRex101/auth-server/app/validators"
	"github.com/MetalRex101/auth-server/app/models"
)

type ActivateHandler struct{
	osm services.IOauthSessionManager
	ocm services.IOauthClientManager
	um services.IUserManager
	em services.IEmailManager
}

func NewActivateHandler(
	osm *services.IOauthSessionManager,
	ocm *services.IOauthClientManager,
	um *services.IUserManager,
	em *services.IEmailManager,
) *ActivateHandler {
	return &ActivateHandler{osm, ocm, um, em}
}

func (act *ActivateHandler) handle (c echo.Context) error {
	var client *models.Client

	if err := validators.Request.OauthTID(c); err != nil {
		return err
	}

	accessToken, err := validators.Request.GetAccessToken(false, c)
	if err != nil {
		return err
	}

	if accessToken != "" {
		oauthSess, err := act.osm.FindByToken(accessToken)
		if err != nil {
			return err
		}

		client, err = act.ocm.GetClientFromSession(oauthSess)
		if err != nil {
			return err
		}

		user, err := act.um.GetUserFromSession(oauthSess)
		if err != nil {
			return err
		}
	} else {
		clientID, err := validators.Request.GetClientId(true, c)
		if err != nil {
			return err
		}

		clientSecret, err := validators.Request.GetClientSecret(true, c)
		if err != nil {
			return err
		}

		client, err = act.ocm.GetForApi(clientID, clientSecret, c.Request().RemoteAddr)
		if err != nil {
			return err
		}
	}

	if err = validators.Client.HasScope(client, []string{"oauth"}); err != nil {
		return err
	}

	code, err := validators.Request.GetCode(validators.CodeTypeActivation, c)
	if err != nil {
		return err
	}

	emailAddr, _ := validators.Request.GetEmail(false, c, false)

	if err := validators.Email.ValidateEmail(emailAddr); err != nil {
		return err
	}

	email, err := act.em.GetEmailToActivate(emailAddr, code)
	if err != nil {
		return err
	}

	if _, err := act.um.UserNotHaveActivatedEmail(*email.UserID, *email.Email); err != nil {
		return err
	}

	if err := act.em.ActivateEmail(email); err != nil {
		return err
	}

	otherUserActivatedEmail := act.em.FindOtherUserActivatedEmail(*email.Email, *email.UserID)

	if otherUserActivatedEmail != nil {

	}

	return nil
}