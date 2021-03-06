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

const timeoutSeconds = 4000

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
		t = time.Now().Add(time.Second * timeoutSeconds)

		oauthSession.AccessExpiresAt = &t
	}

	oauthSession.AccessToken = &accessToken
	oauthSession.Offset = timeOffset
	oauthSession.Code = nil

	err = osm.DB.Save(&oauthSession).Error

	return err
}

func (osm *OauthSessionManager) FindByClientAndCode (client *models.Client, code string) (*models.OauthSession, error) {
	var oauthSession models.OauthSession

	err := osm.DB.Where("client_id = ?", client.ID).
		Where("code = ?", code).
		Where("access_expires_at > ?", carbon.Now().Time).
		Where("access_token IS NULL").
		First(&oauthSession).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Сессия не найдена")
	}

	return &oauthSession, nil
}

func (osm *OauthSessionManager) FindByToken(accessToken string) (*models.OauthSession, error) {
	var oauthSession models.OauthSession

	err := osm.DB.Where("access_token LIKE ?", accessToken).First(&oauthSession).Error
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusForbidden, "Access Token не найден")
	}

	if oauthSession.AccessExpiresAt.Before(time.Now()) {
		return nil, echo.NewHTTPError(http.StatusRequestTimeout, "Срок действия Access Token истек")
	}

	return &oauthSession, nil
}