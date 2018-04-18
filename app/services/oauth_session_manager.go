package services

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/elgs/gostrgen"
	"github.com/labstack/echo"
	"time"
	"github.com/uniplaces/carbon"
	"net/http"
)

type OauthSessionManager struct {
	DB *gorm.DB
}

func NewOauthSessionManager (db *gorm.DB) IOauthSessionManager {
	return &OauthSessionManager{db}
}

func (osm *OauthSessionManager) StartSession (oauthSession *models.OauthSession, timeout bool, timeOffset int, c echo.Context) error {
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

	err = osm.DB.Save(&oauthSession).Error

	return err
}

func (osm *OauthSessionManager) FindByClientIDAndCode (clientID int, code string) (*models.OauthSession, error) {
	var oauthSession models.OauthSession

	err := osm.DB.Where("client_id = ?", clientID).
		Where("code = ?", code).
		Where("access_expires_at > ?", carbon.Now().Time).
		Where("access_token IS NULL").
		First(&oauthSession).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Сессия не найдена")
	}

	return &oauthSession, nil
}