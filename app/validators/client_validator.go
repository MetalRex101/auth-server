package validators

import (
	"github.com/MetalRex101/auth-server/app/models"
	"github.com/viant/toolbox"
	"strings"
	"github.com/labstack/echo"
	"net/http"
)

type ClientValidator struct {
	Base *Validator
}

// Проверяет, имеется ли у клиента разрешение на использование функции
func (v ClientValidator) HasScope (client *models.Client, scopes []string) error {
	clientScopes := strings.Split(*client.Scope, ",")

	for _, scope := range scopes {
		if !toolbox.HasSliceAnyElements(clientScopes, scope) {
			return echo.NewHTTPError(http.StatusForbidden, "Использование функции запрещено")
		}
	}

	return nil
}

func (v ClientValidator) IsClientUrl (client *models.Client, url string, c echo.Context) error {
	equal, err := Url.compareUrls(*client.Url, url)

	if err != nil {
		c.Logger().Error(err)
	}

	if !equal {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Указанная страница не относится к сайту")
	}

	return nil
}