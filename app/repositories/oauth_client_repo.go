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

func (cr *OauthClientRepo) GetForOauth(ClientID int) (*models.Client, error) {
	client := models.Client{}

	err := cr.DB.Where("client_id = ? and status = ?", ClientID, true).First(&client).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Клиент не найден")
	}

	return &client, nil
}