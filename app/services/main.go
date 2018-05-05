package services

import (
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/labstack/echo"
	"crypto/md5"
	"encoding/hex"
)

type Managers struct {
	OauthClient IOauthClientManager
	OauthSession IOauthSessionManager
	User IUserManager
}

type IOauthClientManager interface {
	StartSession(client *models.Client, user *models.User, c echo.Context) (*models.OauthSession, error)
	GetForOauth(clientID int) (*models.Client, error)
	GetForApi(clientID int, clientSecret string, ip string) (*models.Client, error)
}

type IOauthSessionManager interface {
	StartSession (oauthSession *models.OauthSession, timeout bool, timeOffset int, c echo.Context) error
	FindByClientAndCode(client *models.Client, code string) (*models.OauthSession, error)
}

type IUserManager interface {
	GetByEmailAndPassword(email string, pass string, hash bool) (*models.User, error)
}

func InitManagers (db *gorm.DB) *Managers {
	return &Managers{
		OauthClient: NewOauthClientManager(db),
		OauthSession: NewOauthSessionManager(db),
		User: NewUserManager(db),
	}
}

func HashPassword (pass string) string {
	hash := md5.New()
	hash.Write([]byte(pass))

	return hex.EncodeToString(hash.Sum(nil))
}