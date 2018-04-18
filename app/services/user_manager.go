package services

import (
	"github.com/MetalRex101/auth-server/app/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

type UserManager struct {
	DB *gorm.DB
}

func NewUserManager (db *gorm.DB) IUserManager {
	return &UserManager{db}
}

func (um *UserManager) GetByEmailAndPassword(email string, pass string, hash bool) (*models.User, error) {
	var user models.User
	var userEmail models.Email

	if hash {
		hash := md5.New()
		hash.Write([]byte(pass))
		pass = hex.EncodeToString(hash.Sum(nil))
	}

	err := um.DB.Preload("Emails", "email = ?", email).
		Scopes(user.WhereHasPassword(pass), user.WhereHasEmail(email)).First(&user).Error

	if err != nil {
		spew.Dump(err)

		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Пользователь не найден")
		}

		return nil, err
	}

	userEmail = user.Emails[0]

	if userEmail.ConfirmDate == nil {
		return nil, echo.NewHTTPError(
			http.StatusFailedDependency,
			"Этот адрес электронной почты не подтвержден",
		)
	}

	if !userEmail.Status {
		return nil, echo.NewHTTPError(
			http.StatusLocked,
			"Пользователь заблокирован. Обратитесь к администрации сайта",
		)
	}

	return &user, nil
}
