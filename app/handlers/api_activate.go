package handlers

import (
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/labstack/echo"
	"github.com/MetalRex101/auth-server/app/validators"
	"github.com/MetalRex101/auth-server/app/models"
	"net/http"
)

type ActivateHandler struct{
	osm  services.IOauthSessionManager
	ocm  services.IOauthClientManager
	um   services.IUserManager
	em   services.IEmailManager
	uMer services.IUserMerger
}

func NewActivateHandler(
	osm services.IOauthSessionManager,
	ocm services.IOauthClientManager,
	um services.IUserManager,
	em services.IEmailManager,
	uMer services.IUserMerger,
) *ActivateHandler {
	return &ActivateHandler{osm, ocm, um, em, uMer}
}

func (act *ActivateHandler) Handle (c echo.Context) error {
	var client *models.Client
	var user *models.User

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

		user, err = act.um.GetUserFromSession(oauthSess)
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

	emailToActivate, err := act.em.GetEmailToActivate(emailAddr, code)
	if err != nil {
		return err
	}

	if _, err := act.um.UserNotHaveActivatedEmail(*emailToActivate.UserID, *emailToActivate.Email); err != nil {
		return err
	}

	if err := act.em.ActivateEmail(emailToActivate); err != nil {
		return err
	}

	otherUserActivatedEmail := act.em.FindOtherUserActivatedEmail(*emailToActivate.Email, *emailToActivate.UserID)

	if otherUserActivatedEmail != nil {
		act.uMer.MergerUsers(emailToActivate.User, otherUserActivatedEmail.User, c)
	}

	if user != nil {
		act.uMer.MergerUsers(user, emailToActivate.User, c)
	}

	return c.JSON(http.StatusOK, DefaultResponse{Status:1})
}