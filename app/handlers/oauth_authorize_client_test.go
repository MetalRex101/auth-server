package handlers

import (
	"testing"
	"net/http/httptest"
	"fmt"
	"encoding/json"
	"net/url"
	"github.com/labstack/echo"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/db"
	"github.com/MetalRex101/auth-server/config"
	"github.com/stretchr/testify/assert"
	"net/http"
)

type AuthorizeClientResponse struct {
	Code string `json:"code"`
}

func TestAuthorizeClient(t *testing.T) {
	var response AuthorizeClientResponse

	q := make(url.Values)
	q.Set("client_id", "32131231")
	q.Set("password", "dmsladmskdl")
	q.Set("redirect_uri", "url")
	q.Set("email", "metal@gmail.com")

	e := echo.New()

	req := httptest.NewRequest(echo.GET, fmt.Sprintf("/authorize?%s", q.Encode()), nil)
	rec := httptest.NewRecorder()

	// todo mock services, depending database
	db := db.Init(config.GetConfig("testing"))

	handler := NewAuthorizeClientHandler(services.NewOauthClientManager(db), services.NewUserManager(db))

	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handler.Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NotEmpty(t, response.Code)
	}
}
