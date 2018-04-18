package repositories

import (
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"net/http"
	"github.com/jinzhu/gorm"
)

type OauthClientRepo struct {
	DB *gorm.DB
}

func NewOauthClientRepo(db *gorm.DB) *OauthClientRepo {
	return &OauthClientRepo{db}
}

func (cr *OauthClientRepo) GetForOauth(clientID int) (*models.Client, error) {
	var client models.Client

	err := cr.DB.Where("client_id = ? and status = ?", clientID, true).First(&client).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Клиент не найден")
	}

	return &client, nil
}

func (cr *OauthClientRepo) GetForApi(clientID int, clientSecret string, ip string) (*models.Client, error) {
	var client models.Client

	err := cr.DB.Where(
		"client_id = ? and client_secret = ? and status = ?",
		clientID,
		clientSecret,
		true,
	).
		Where("ip = '*' OR ip = ?", ip).
		First(&client).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Клиент не найден")
	}

	return &client, nil
}