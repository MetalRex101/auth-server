package repositories

import (
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"net/http"
	"github.com/MetalRex101/auth-server/app/db"
)

type oauthClientRepo struct {

}

func (oauthClientRepo) GetForOauth(ClientID int) (*models.Client, error) {
	client := models.Client{}

	err := db.Gorm.Where("client_id = ? and status = ?", ClientID, true).First(&client).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Клиент не найден")
	}

	return &client, nil
}