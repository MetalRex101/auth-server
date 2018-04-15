package services

import (
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"time"
	"github.com/elgs/gostrgen"
	"github.com/jinzhu/gorm"
)

type OauthClientManager struct {
	DB *gorm.DB
}

func NewOauthClientManager (db *gorm.DB) *OauthClientManager {
	return &OauthClientManager{db}
}

func (cm *OauthClientManager) StartSession(client *models.Client, user *models.User, c echo.Context) (*models.OauthSession, error) {
	oauthSession := &models.OauthSession{}

	cm.DB.FirstOrInit(oauthSession, models.OauthSession{
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

	cm.DB.Save(oauthSession)

	return oauthSession, nil
}
