package services

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/elgs/gostrgen"
	"github.com/labstack/echo"
	"time"
)

type OauthSessionManager struct {
	DB *gorm.DB
}

func NewOauthSessionManager (db *gorm.DB) *OauthSessionManager {
	return &OauthSessionManager{db}
}

func (osm *OauthSessionManager) StartSession (oauthSession *models.OauthSession, timeout bool, timeOffset int, c echo.Context) {
	var t time.Time

	accessToken, err := gostrgen.RandGen(64, gostrgen.All, "", "")
	if err != nil {
		c.Logger().Error(err)
	}

	if !timeout {
		t = time.Now().Add(time.Second * 4000)

		oauthSession.AccessExpiresAt = &t
	}

	oauthSession.AccessToken = &accessToken
	oauthSession.Offset = timeOffset
	oauthSession.Code = nil

	osm.DB.Save(&oauthSession)
}