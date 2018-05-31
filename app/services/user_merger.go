package services

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/app/models"
	"strings"
	"github.com/labstack/echo"
	"strconv"
	"time"
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

	um.transferEmails(receiver, sender, c.Logger())

	um.um.GetDefaultEmail(sender.ID, true, c)

	um.transferPhones(receiver, sender, c.Logger())

	um.um.GetDefaultPhone(sender.ID, true, c)

	um.transferPasswords(receiver, sender, c.Logger())

	um.updateRegistrationTime(receiver, sender, c.Logger())

	um.transferRoles(receiver, sender, c.Logger())

	uid := sender.ID

	um.DB.Delete(sender)

	now := time.Now()

	receiver.ChangedAt = &now
	if err := um.DB.Save(receiver).Error; err != nil {
		c.Logger().Error(err)
	}

	return uid, nil
}

func (um *UserMerger) transferEmails(receiver *models.User, sender *models.User, l echo.Logger) {
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
					l.Error(err)
				}
			}

			if err := um.addEmailOauthProvider(&receiverEmail, *email.Oauth); err != nil {
				l.Errorf("Произошла ошибка при добавлении oauth провайдера: %s", err)
			}
		} else {
			email.UserID = &receiver.ID
			if err := um.DB.Save(&email).Error; err != nil {
				l.Errorf(
					"При передаче email %s произошла ошибка: %s",
					email.Email,
					err,
				)
			} else {
				l.Debugf(
					"Пользователь %s передал email %s пользователю %s",
					sender.ID,
					email.Email,
					receiver.ID,
				)
			}
		}
	}
}

func (um *UserMerger) transferPhones(receiver *models.User, sender *models.User, l echo.Logger) {
	var receiverPhone models.Phone

	for _, phone := range sender.Phones {
		err := um.DB.
			Where("user_id = ?", sender.ID).
			Where("email = ?", phone.Phone).
			First(&receiverPhone).Error

		if err == nil {
			if phone.ConfirmDate != nil {
				receiverPhone.ConfirmDate = phone.ConfirmDate
				if err := um.DB.Save(&receiverPhone).Error; err != nil {
					l.Error(err)
				}
			}

			if err := um.addPhoneOauthProvider(&receiverPhone, *phone.Oauth); err != nil {
				l.Errorf("Произошла ошибка при добавлении oauth провайдера: %s", err)
			}
		} else {
			phone.UserID = &receiver.ID
			if err := um.DB.Save(&phone).Error; err != nil {
				l.Errorf(
					"При передаче телефона %s произошла ошибка: %s",
					phone.Phone,
					err,
				)
			} else {
				l.Debugf(
					"Пользователь %s передал телефон %s пользователю %s",
					sender.ID,
					phone.Phone,
					receiver.ID,
				)
			}
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

func (um *UserMerger) addPhoneOauthProvider(phone *models.Phone, provider string) error {
	emailOauth := strings.Split(*phone.Oauth, ",")
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
	phone.Oauth = &joined

	if err := um.DB.Save(phone).Error; err != nil {
		return err
	}

	return nil
}

func (um *UserMerger) transferPasswords (receiver *models.User, sender *models.User, l echo.Logger) {
	err := um.DB.
		Table("passwords").
		Where("user_id = ?", sender.ID).
		Update("user_id", receiver.ID).Error

	if err != nil {
		l.Errorf(
			"При передаче паролей пользователем %d пользователю %d произошла ошибка: %s",
			sender.ID,
			receiver.ID,
			err,
		)
	} else {
		l.Debugf(
			"Пользователь %s передал пароли пользователю %s",
			sender.ID,
			receiver.ID,
		)
	}
}

func (um *UserMerger) transferRoles (receiver *models.User, sender *models.User, l echo.Logger) {
	var roles []models.Role

	err := um.DB.Model(sender).Related(&roles).Error

	if err != nil {
		l.Error(err)
	}

	err = um.DB.Model(receiver).Association("Roles").Append(roles).Error

	if err != nil {
		l.Error(err)
	}

	if err == nil {
		l.Debugf(
			"Пользователь %d Передали роли пользователю %d",
			sender.ID,
			receiver.ID,
		)
	}
}

func (um *UserMerger) updateRegistrationTime (receiver *models.User, sender *models.User, l echo.Logger) {
	if !receiver.CreatedAt.Before(*sender.CreatedAt) {
		receiver.CreatedAt = sender.CreatedAt
	}

	if err := um.DB.Save(receiver).Error; err != nil {
		l.Errorf(
		"При передачи пользователем %d даты регистрации пользователю %d произошла ошибка: %s",
			sender.ID,
			receiver.ID,
			err,
		)
	}
}

func (um *UserMerger) appendMerged(receiver *models.User, sender *models.User, l echo.Logger) {
	merged := strings.Split(*receiver.Merged, ",")
	merged = append(merged, strconv.FormatUint(uint64(sender.ID), 10))

	mergedString := strings.Join(merged, ",")
	receiver.Merged = &mergedString
	if err := um.DB.Save(&receiver).Error; err != nil {
		l.Error(err)
	}
}