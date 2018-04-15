package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/davecgh/go-spew/spew"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"net/http"
	"crypto/md5"
	"encoding/hex"
)

type UserRepo struct{
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (ur *UserRepo) GetByEmailAndPassword(email string, pass string, hash bool) (*models.User, error) {
	var user models.User
	var userEmail models.Email

	if hash {
		hash := md5.New()
		hash.Write([]byte(pass))
		pass = hex.EncodeToString(hash.Sum(nil))
	}

	err := ur.DB.Preload("Emails", "email = ?", email).
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
