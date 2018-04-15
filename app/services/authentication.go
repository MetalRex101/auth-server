package services

import (
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"time"
	"github.com/MetalRex101/auth-server/app/db"
)

const sessionVar = "authorized"

type AuthService struct{}

func (as AuthService) AuthenticatedUser(c echo.Context) (*models.User, error) {
	var user models.User

	userId, ok := Session.Get(sessionVar, c)

	if ok {
		if err := db.Gorm.First(&user, userId.(uint)).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				c.Logger().Error(err)
			}

			return nil, err
		}
	}

	return &user, nil
}

func (as AuthService) LogoutCurrentUser(c echo.Context) error {
	if _, err := as.AuthenticatedUser(c); err == nil {
		if err = Session.Delete(sessionVar, c); err != nil {
			return err
		}
	}

	return nil
}

func (as AuthService) AuthenticateUser(user *models.User, c echo.Context) error {
	Session.Put(sessionVar, user.ID, c)
	now := time.Now()

	user.LastVisit = &now

	if err := db.Gorm.Save(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.Logger().Error(err)
		}

		return err
	}

	return nil
}