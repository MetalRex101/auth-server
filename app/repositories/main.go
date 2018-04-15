package repositories

var OauthClient *oauthClientRepo
var User *userRepo

func init() {
	OauthClient = &oauthClientRepo{}
	User = &userRepo{}
}
