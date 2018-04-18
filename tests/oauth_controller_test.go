package tests

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"net/url"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"encoding/json"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/elgs/gostrgen"
	"github.com/MetalRex101/auth-server/app"
	"github.com/MetalRex101/auth-server/config"
	"time"
)

type authorizeResponse struct {
	Code string `json:"code"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Expires string `json:"expires"`
}

func TestAuthorizeClient(t *testing.T) {
	var response authorizeResponse

	q := make(url.Values)
	q.Set("client_id", "32131231")
	q.Set("password", "dmsladmskdl")
	q.Set("redirect_uri", "url")
	q.Set("email", "metal@gmail.com")

	e := echo.New()

	req := httptest.NewRequest(echo.GET, fmt.Sprintf("/authorize?%s", q.Encode()), nil)
	rec := httptest.NewRecorder()
	app := app.NewApp(config.GetConfig("testing"))
	app.InitializeServices()

	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, app.Controllers.Oauth.AuthorizeClient(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NotEmpty(t, response.Code)
	}
}

func TestGetAccessToken(t *testing.T) {
	var response AccessTokenResponse

	clientID := uint(32131231)
	userID := uint(4)
	expires := time.Now().Add(time.Hour)
	code, _ := gostrgen.RandGen(128, gostrgen.Lower | gostrgen.Upper | gostrgen.Digit, "", "")
	ip := "*"

	oauthSession := models.OauthSession{
		ClientID: &clientID,
		UserID: &userID,
		AccessExpiresAt: &expires,
		Code: &code,
		RemoteAddr: &ip,
	}

	app := app.NewApp(config.GetConfig("testing"))
	app.InitializeServices()

	app.DB.Create(&oauthSession)

	q := make(url.Values)
	q.Set("client_id", "32131231")
	q.Set("client_secret", "jhdjsad738123")
	q.Set("code", code)

	e := echo.New()

	req := httptest.NewRequest(echo.GET, fmt.Sprintf("/access_token?%s", q.Encode()), nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.Request().RemoteAddr = "127.0.0.1"

	// Assertions
	if assert.NoError(t, app.Controllers.Oauth.GetAccessToken(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NotEmpty(t, response.AccessToken)
	}
}