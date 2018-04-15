package services

import (
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"github.com/MetalRex101/auth-server/app/db"
	"time"
	"github.com/elgs/gostrgen"
)

type OauthClientManager struct {
	AuthService AuthService
}

func (cm OauthClientManager) StartSession(client *models.Client, c echo.Context) (*models.OauthSession, error) {
	oauthSession := &models.OauthSession{}
	user, err := cm.AuthService.AuthenticatedUser(c)

	if err != nil {
		return nil, err
	}

	db.Gorm.FirstOrInit(oauthSession, models.OauthSession{
		ClientID: &client.ID,
		UserID: &user.ID,
	})

	accessGranted := time.Now()
	accessExpires := time.Now().Add(time.Hour)

	oauthSession.AccessGrantedAt = &accessGranted
	oauthSession.AccessExpiresAt = &accessExpires

	code, err := gostrgen.RandGen(128, gostrgen.Lower | gostrgen.Upper | gostrgen.Digit, "", "")

	if err != nil {
		return nil, err
	}

	oauthSession.Code = &code

	userAgent := c.Request().UserAgent()
	oauthSession.UserAgent = &userAgent

	remoteAddr := c.Request().RemoteAddr
	oauthSession.UserAgent = &remoteAddr

	db.Gorm.Save(oauthSession)

	return oauthSession, nil
}
