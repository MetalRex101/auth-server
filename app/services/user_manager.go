package services

import (
	"github.com/MetalRex101/auth-server/app/models"
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
		pass = HashPassword(pass)
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

func (um *UserManager) GetUserFromSession(sess *models.OauthSession) (*models.User, error) {
	var user models.User

	err := um.DB.Where("id = ?", sess.UserID).First(&user).Error
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Пользователь не найден")
	}

	return &user, nil
}

func (um *UserManager) UserNotHaveActivatedEmail (userID uint, emailAddr string) (bool, error) {
	var email models.Email

	err := um.DB.
		Where("user_id = ?", userID).
		Where("email = ?", emailAddr).
		Where("IS NOT NULL confirm_date").
		First(&email).Error

	if err != nil {
		return false, echo.NewHTTPError(
			http.StatusExpectationFailed,
			"Пользователь с этим адресом электронной почты уже активирован",
		)
	}

	return true, nil
}

func (um *UserManager) GetDefaultEmail (userID uint, update bool, c echo.Context) (*models.Email, error) {
	var email models.Email

	err := um.DB.
		Where("user_id = ?", userID).
		Order("is_default desc").
		Order("status desc").
		Order("created_at").
		First(&email).Error

	if err != nil {
		return nil, err
	}

	email.IsDefault = true
	um.DB.Save(&email)

	if update {
		um.updateDefaultEmail(&email, userID, c.Logger())
	}

	return &email, nil
}

func (um *UserManager) updateDefaultEmail (email *models.Email, userID uint, l echo.Logger) {
	err := um.DB.
		Table("emails").
		Where("user_id = ?", userID).
		Where("id != ?", email.ID).
		Update("is_default", false).Error

	if err != nil {
		l.Error(err)
	}
}

func (um *UserManager) GetDefaultPhone (userID uint, update bool, c echo.Context) (*models.Phone, error) {
	var phone models.Phone

	err := um.DB.
		Where("user_id = ?", userID).
		Order("is_default desc").
		Order("status desc").
		Order("created_at").
		First(&phone).Error

	if err != nil {
		return nil, err
	}

	phone.IsDefault = true
	um.DB.Save(&phone)

	if update {
		um.updateDefaultPhone(&phone, userID, c.Logger())
	}

	return &phone, nil
}

func (um *UserManager) updateDefaultPhone (phone *models.Phone, userID uint, l echo.Logger) {
	err := um.DB.
		Table("phones").
		Where("user_id = ?", userID).
		Where("id != ?", phone.ID).
		Update("is_default", false).Error

	if err != nil {
		l.Error(err)
	}
}