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
)

func TestAuthorizeClient(t *testing.T) {
	q := make(url.Values)
	q.Set("client_id", "32131231")
	q.Set("password", "dmsladmskdl")
	q.Set("redirect_uri", "url")
	q.Set("email", "metal@gmail.com")

	e := echo.New()
	e.Debug = true
	//e.Use(session.Middleware(sessions.NewFilesystemStore("../storage/sessions")))

	req := httptest.NewRequest(echo.GET, fmt.Sprintf("/authorize?%s", q.Encode()), nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.Oauth.AuthorizeClient(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, "Code", rec.Body)
	}
}
