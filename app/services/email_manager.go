package services

import (
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type EmailManager struct {
	DB *gorm.DB
}

func NewEmailManager(db *gorm.DB) IEmailManager {
	return &EmailManager{db}
}

func (em *EmailManager) GetEmailToActivate(emailAddr string, code string) (*models.Email, error) {
	var email models.Email

	err := em.DB.
		Where("email = ?", emailAddr).
		Where("code = ?", code).
		Where("confirm_date IS NULL").
		First(&email).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Email не найден")
	}

	return &email, nil
}

func (em *EmailManager) ActivateEmail(email *models.Email) error {
	now := time.Now()

	email.ConfirmDate = &now
	email.Code = nil

	if err := em.DB.Save(email).Error; err != nil {
		return err
	}

	return nil
}

func (em *EmailManager) FindOtherUserActivatedEmail(addr string, userID uint) *models.Email {
	var email models.Email

	err := em.DB.
		Where("email = ?", addr).
		Where("user_id != ?", userID).
		Where("NOT NULL confirm_date").
		First(&email).Error

	if err != nil {
		return nil
	}

	return &email
}

func (em *EmailManager) EmailNotUsed (addr string) error {
	var email models.Email

	err := em.DB.
		Where("email = ?", email).
		Where("IS NOT NULL confirm_date").
		First(&email).Error

	if err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(
			http.StatusConflict,
			"Указанный адрес e-mail уже используется другим пользователем",
		)
	}

	return nil
}