package tests

import (
	"testing"
	"net/http/httptest"
	"fmt"
	"encoding/json"
	"net/url"
	"github.com/labstack/echo"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"github.com/MetalRex101/auth-server/app/handlers"
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/tests/testdata"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/MetalRex101/auth-server/app/db"
	"os"
	"github.com/MetalRex101/auth-server/config"
	"log"
	"github.com/elgs/gostrgen"
)

func TestMain(m *testing.M) {
	conf := config.GetConfig("testing")

	dbConn := db.Init(conf)
	if err := dbConn.Begin().Error; err != nil {
		log.Panic(err)
	}

	db.MigrateDown(dbConn, conf)
	db.MigrateUp(dbConn, conf)

	result := m.Run()

	db.MigrateDown(dbConn, conf)

	os.Exit(result)
}

type AuthorizeClientResponse struct {
	Code string `json:"code"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Expires string `json:"expires"`
}

func TestAuthorizeClient(t *testing.T) {
	withTx(func (tx *gorm.DB) error {
		var response AuthorizeClientResponse

		user := testdata.UserFact.MustCreate().(*models.User)
		tx.Save(user)

		client := testdata.ClientFact.MustCreate().(*models.Client)
		tx.Save(client)

		rawPass := "jdkajld739183"
		hashedPass := services.HashPassword(rawPass)

		password := testdata.PasswordFact.MustCreateWithOption(map[string]interface{}{
			"Password": &hashedPass,
			"UserID": &user.ID,
		}).(*models.Password)
		tx.Save(password)

		email := testdata.EmailFact.MustCreateWithOption(map[string]interface{}{
			"UserID": &user.ID,
		}).(*models.Email)
		tx.Save(email)

		q := make(url.Values)
		q.Set("email", *email.Email)
		q.Set("password", rawPass)
		q.Set("client_id", *client.ClientID)
		q.Set("redirect_uri", *client.Url)

		e := echo.New()

		req := httptest.NewRequest(echo.GET, fmt.Sprintf("/authorize?%s", q.Encode()), nil)
		rec := httptest.NewRecorder()

		handler := handlers.NewAuthorizeClientHandler(
			services.NewOauthClientManager(tx),
			services.NewUserManager(tx),
		)

		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, handler.Handle(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NotEmpty(t, response.Code)
		}

		return nil
	})
}

func TestGetAccessToken(t *testing.T) {
	withTx(func (tx *gorm.DB) error {
		var response AccessTokenResponse

		user := testdata.UserFact.MustCreate().(*models.User)
		tx.Save(user)

		client := testdata.ClientFact.MustCreate().(*models.Client)
		tx.Save(client)

		code, err := gostrgen.RandGen(128, gostrgen.LowerUpperDigit, "", "")
		if err != nil {
			return err
		}

		oauthSess := testdata.OauthSessFact.MustCreateWithOption(map[string]interface{}{
			"Code": &code,
			"ClientID": &client.ID,
			"UserID": &user.ID,
		}).(*models.OauthSession)
		tx.Save(oauthSess)

		q := make(url.Values)
		q.Set("client_id", *client.ClientID)
		q.Set("client_secret", *client.ClientSecret)
		q.Set("code", *oauthSess.Code)

		e := echo.New()

		req := httptest.NewRequest(echo.GET, fmt.Sprintf("/access_token?%s", q.Encode()), nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.Request().RemoteAddr = *oauthSess.RemoteAddr

		handler := handlers.NewAccessTokenClientHandler(
			services.NewOauthClientManager(tx),
			services.NewOauthSessionManager(tx),
		)

		// Assertions
		if assert.NoError(t, handler.Handle(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NotEmpty(t, response.AccessToken)
		}

		return nil
	})
}
