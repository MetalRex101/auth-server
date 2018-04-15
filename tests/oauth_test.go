package tests

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"net/url"
	"fmt"
	"github.com/MetalRex101/auth-server/app/controllers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"github.com/MetalRex101/auth-server/app/services"
	"github.com/MetalRex101/auth-server/app/repositories"
	"encoding/json"
	"github.com/MetalRex101/auth-server/app/db"
)

type authorizeResponse struct {
	Code string `json:"code"`
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

	conn := db.Init()

	managers := services.InitManagers(conn)
	repos := repositories.InitRepositories(conn)

	controllers := controllers.InitControllers(repos, managers)

	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.Oauth.AuthorizeClient(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NotEmpty(t, response.Code)
	}
}
