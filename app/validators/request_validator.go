package validators

import (
	"github.com/labstack/echo"
	"strconv"
	"net/http"
	"fmt"
	"time"
)

const CodeTypeActivation = "activation"
const CodeTypeAccess = "access"

type RequestValidator struct {
	Base *Validator
}

func (v RequestValidator) OauthTID(c echo.Context) error {
	accessToken  := c.QueryParam("access_token")
	clientId     := c.QueryParam("client_id")
	clientSecret := c.QueryParam("client_secret")

	if (accessToken == "" && clientId == "") || clientSecret == "" {
		return NewValidationError("Не указан client_id, client_secret или access_token")
	}

	return nil
}

// Возвращает client_id, если тот был передан в запросе
func (v RequestValidator) GetClientId(validate bool, c echo.Context) (int, error) {
	clientID := c.QueryParam("client_id")

	if clientID == "" {
		clientID = c.Request().Header.Get("Client-ID")
	}

	if validate && clientID == "" {
		return -1, echo.NewHTTPError(http.StatusBadRequest, "Не указан client id")
	}

	i, err := strconv.Atoi(clientID)

	if err != nil {
		return -1, echo.NewHTTPError(http.StatusBadRequest, "Неверный формат client id")
	}

	return i, nil
}

// Проверяет, был ли передан access_token
func (v RequestValidator) GetAccessToken(validate bool, c echo.Context) (string, error) {
	accessToken := c.QueryParam("access_token")

	if accessToken == "" && validate {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Не указан access_token или Authorization")
	}

	return accessToken, nil
}

func (v RequestValidator) GetClientSecret(validate bool, c echo.Context) (string, error) {
	clientSecret := c.QueryParam("client_secret")

	if validate && clientSecret == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Не указан client secret")
	}

	return clientSecret, nil
}

func (v RequestValidator) GetRedirectUri(validate bool, c echo.Context) (string, error) {
	redirectUri := c.QueryParam("redirect_uri")

	if redirectUri == "" && validate {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Не указан redirect_uri")
	}

	return redirectUri, nil
}

func (v RequestValidator) GetPassword(validate bool, c echo.Context, fromBody bool) (string, error) {
	var password string

	if fromBody {
		m := echo.Map{}

		if err := c.Bind(&m); err != nil {
			return "", err
		}

		password = m["password"].(string)
	} else {
		password = c.QueryParam("password")
	}

	if password == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Не указан пароль")
	}

	return password, nil
}

func (v RequestValidator) GetEmail(validate bool, c echo.Context, fromBody bool) (string, error) {
	var email string

	if fromBody {
		m := echo.Map{}

		if err := c.Bind(&m); err != nil {
			return "", err
		}

		email = m["email"].(string)
	} else {
		email = c.QueryParam("email")
	}

	if email == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Не указан email")
	}

	return email, nil
}

func (v RequestValidator) GetCode(codeType string, c echo.Context) (string, error) {
	code := c.QueryParam("code")

	if code == "" {
		m := echo.Map{}

		if err := c.Bind(&m); err != nil {
			return "", echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Не указан %s код", codeType))
		}

		code = m["code"].(string)
	}

	return code, nil
}

func (v RequestValidator) GetUrl(c echo.Context) (string, error) {
	url := c.QueryParam("url")

	if url == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Не указан url")
	}

	return url, nil
}

func (v RequestValidator) GetGender (validate bool, c echo.Context) (string, error) {
	gender := c.QueryParam("gender")

	if gender == "" && validate {
		return "", echo.NewHTTPError(http.StatusNotFound, "Не указан gender")
	}

	if validate {
		genders := []string{"male", "female"}
		has := false

		for _, gend := range genders {
			if gend == gender {
				has = true
			}
		}

		if !has {
			return "", echo.NewHTTPError(
				http.StatusNotAcceptable,
				"Аргумент gender должен иметь значение male или female",
			)
		}
	}

	return gender, nil
}

func (v RequestValidator) GetBirthDate (validate bool, c echo.Context) (*time.Time, error) {
	birthDate := c.QueryParam("birthDate")

	if birthDate == "" {
		if validate {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Не указан birthDate")
		} else {
			return nil, nil
		}
	}

	timestamp, err := strconv.ParseInt(birthDate, 10, 64)
	if err != nil && validate {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "Неверный формат birthdate")
	}

	date := time.Unix(timestamp, 0)

	return &date, nil
}