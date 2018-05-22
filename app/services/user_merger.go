package services

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/app/models"
	"strings"
	"github.com/labstack/echo"
)

type UserMerger struct {
	DB *gorm.DB
	um IUserManager
}

func NewUserMerger (db *gorm.DB, um *UserManager) IUserMerger {
	return &UserMerger{db, um}
}

type MergeError struct {
	message string
}

func newMergeError(message string) *MergeError {
	return &MergeError{
		message: message,
	}
}

func (e *MergeError) Error() string {
	return e.message
}

func (um *UserMerger) MergerUsers(receiver *models.User, sender *models.User, c echo.Context) (uint, error) {
	if receiver.ID == sender.ID {
		return 0, newMergeError("Сливаемые пользователи равны")
	}

	um.transferEmails(receiver, sender, c)

	um.um.GetDefaultEmail(sender.ID, true, c)
}

func (um *UserMerger) transferEmails(receiver *models.User, sender *models.User, c echo.Context) {
	var receiverEmail models.Email

	for _, email := range sender.Emails {
		err := um.DB.
			Where("user_id = ?", sender.ID).
			Where("email = ?", email.Email).
			First(&receiverEmail).Error

		if err == nil {
			if email.ConfirmDate != nil {
				receiverEmail.ConfirmDate = email.ConfirmDate
				if err := um.DB.Save(&receiverEmail).Error; err != nil {
					c.Logger().Error(err)
				}
			}

			if err := um.addEmailOauthProvider(&receiverEmail, *email.Oauth); err != nil {
				c.Logger().Errorf("Произошла ошибка при добавлении oauth провайдера: %s", err)
			}
		} else {
			email.UserID = &receiver.ID
			if err := um.DB.Save(&email).Error; err != nil {
				c.Logger().Error(err)
			}

			c.Logger().Debugf(
				"Пользователь %s передал email %s пользователю %s",
				sender.ID,
				email.ID,
				receiver.ID,
			)
		}
	}
}

func (um *UserMerger) addEmailOauthProvider(email *models.Email, provider string) error {
	emailOauth := strings.Split(*email.Oauth, ",")
	providerOauth := strings.Split(provider, ",")

	oauthsMap := make(map[string]bool)

	for _, oauth := range append(emailOauth, providerOauth...) {
		if oauth != "" {
			oauthsMap[oauth] = true
		}
	}

	uniqueOauths := make([]string, 0, len(oauthsMap))

	for key := range oauthsMap {
		uniqueOauths = append(uniqueOauths, key)
	}

	joined := strings.Join(uniqueOauths, ",")
	email.Email = &joined

	if err := um.DB.Save(email).Error; err != nil {
		return err
	}

	return nil
}