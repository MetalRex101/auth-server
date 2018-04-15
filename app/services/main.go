package services

var Session *SessionService
var Auth *AuthService

func init() {
	Session = &SessionService{}
	Auth = &AuthService{}
}
