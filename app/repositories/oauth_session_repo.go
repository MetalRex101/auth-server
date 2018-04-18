package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/uniplaces/carbon"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"net/http"
)

type OauthSessionRepo struct {
	DB *gorm.DB
}

func NewOauthSessionRepo (db *gorm.DB) *OauthSessionRepo {
	return &OauthSessionRepo{db}
}

func (osr *OauthSessionRepo) Find (clientID int, code string) (*models.OauthSession, error) {
	var oauthSession models.OauthSession

	err := osr.DB.Where("client_id = ?", clientID).
		Where("code = ?", code).
		Where("access_expires_at > ?", carbon.Now().Time).
		Where("access_token IS NULL").
		First(&oauthSession).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Сессия не найдена")
	}

	return &oauthSession, nil
}
