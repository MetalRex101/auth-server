package tests

import (
	"testing"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/models"
	"net/http/httptest"
	"fmt"
	"github.com/MetalRex101/auth-server/app/handlers"
	"github.com/jinzhu/gorm"
	"github.com/MetalRex101/auth-server/tests/testdata"
	"net/url"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"github.com/Pallinder/go-randomdata"
)

func TestRegister(t *testing.T) {
	withTx(func (tx *gorm.DB) error {
		user := testdata.UserFact.MustCreate().(*models.User)
		tx.Save(user)
		client := testdata.ClientFact.MustCreate().(*models.Client)
		tx.Save(client)
		sess := testdata.OauthSessFact.MustCreate().(*models.OauthSession)
		sess.UserID = &user.ID
		sess.ClientID = &client.ID
		tx.Save(sess)

		rawPass := randomdata.RandStringRunes(16)
		email := randomdata.Email()
		gender := "male"

		q := make(url.Values)
		q.Set("email", email)
		q.Set("password", rawPass)
		q.Set("client_id", *client.ClientID)
		q.Set("client_secret", *client.ClientSecret)
		q.Set("gender", gender)

		e := echo.New()

		req := httptest.NewRequest(echo.GET, fmt.Sprintf("/register?%s", q.Encode()), nil)
		rec := httptest.NewRecorder()

		managers := services.InitManagers(tx)
		merger := services.NewUserMerger(tx, managers.User.(*services.UserManager))

		handler := handlers.NewRegisterHandler(
			managers.OauthSession,
			managers.OauthClient,
			managers.User,
			managers.Email,
			merger,
			tx,
		)

		c := e.NewContext(req, rec)
		c.Request().RemoteAddr = "127.0.0.1"

		if assert.NoError(t, handler.Handle(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			//json.Unmarshal(rec.Body.Bytes(), &response)
			//assert.NotEmpty(t, response.Code)
		}

		return nil
	})
}