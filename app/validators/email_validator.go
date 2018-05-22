package validators

import (
	"net/mail"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"github.com/MetalRex101/auth-server/app/models"
	"time"
)

type EmailValidator struct {
	Base *Validator
}

func (ev *EmailValidator) ValidateEmail (addr string) error {
	_, err := mail.ParseAddress(addr)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusNotAcceptable,
			fmt.Sprintf("Указан невозможный адрес e-mail: '%s'", addr),
		)
	}

	return nil
}

func (ev *EmailValidator) ActivationCodeHasNotExpired (email *models.Email) error {
	if email.ConfirmDate.Before(time.Now().AddDate(0, 0, -1)) {
		return echo.NewHTTPError(http.StatusRequestTimeout, "Код активации устарел")
	}

	return nil
}