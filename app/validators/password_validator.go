package validators

import (
	"github.com/labstack/echo"
	"net/http"
)

type PasswordValidator struct {
	Base *Validator
}

func (pv *PasswordValidator) ValidatePassword (pass string) error {
	passLen := len([]rune(pass))

	if passLen < 6 {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Пароль не должен быть меньше 6 символов")
	}

	return nil
}