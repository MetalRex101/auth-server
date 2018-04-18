package handlers

import (
	"net/http/httptest"
	"fmt"
	"encoding/json"
	"testing"
	"time"
	"github.com/elgs/gostrgen"
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/db"
	"github.com/MetalRex101/auth-server/config"
	"net/url"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Expires string `json:"expires"`
}

func TestGetAccessToken(t *testing.T) {
	var response AccessTokenResponse

	clientID := uint(32131231)
	userID := uint(4)
	expires := time.Now().Add(time.Hour)
	code, _ := gostrgen.RandGen(128, gostrgen.LowerUpperDigit, "", "")
	ip := "*"

	oauthSession := models.OauthSession{
		ClientID: &clientID,
		UserID: &userID,
		AccessExpiresAt: &expires,
		Code: &code,
		RemoteAddr: &ip,
	}

	// todo mock services, depending database
	db := db.Init(config.GetConfig("testing"))
	db.Create(&oauthSession)

	q := make(url.Values)
	q.Set("client_id", "32131231")
	q.Set("client_secret", "jhdjsad738123")
	q.Set("code", code)

	e := echo.New()

	req := httptest.NewRequest(echo.GET, fmt.Sprintf("/access_token?%s", q.Encode()), nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.Request().RemoteAddr = "127.0.0.1"

	handler := NewAccessTokenClientHandler(services.NewOauthClientManager(db), services.NewOauthSessionManager(db))

	// Assertions
	if assert.NoError(t, handler.Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NotEmpty(t, response.AccessToken)
	}
}
