package services

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo"
	"github.com/gorilla/sessions"
)

type SessionService struct {}

const sessionName = "session"

func (ss SessionService) Get (key string, c echo.Context) (interface{}, bool) {
	sess, _ := session.Get(sessionName, c)

	value, ok := sess.Values[key]

	return value, ok
}

func (ss SessionService) Put (key string, value interface{}, c echo.Context) error {
	sess, _ := session.Get(sessionName, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values[key] = value

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		c.Logger().Error(err)

		return err
	}

	return nil
}

func (ss SessionService) Delete (key string, c echo.Context) error {
	sess, _ := session.Get(sessionName, c)
	delete(sess.Values, sessionVar)

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		c.Logger().Error(err)

		return err
	}

	return nil
}

func NewSession () *SessionService {
	return &SessionService{}
}