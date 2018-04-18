package resources

import (
	"github.com/MetalRex101/auth-server/app/models"
	"time"
)

type AccessTokenResource struct {
	AccessToken string `json:"access_token"`
	Expires string `json:"expires"`
}

func NewAccessTokenResource(oauthSession *models.OauthSession) *AccessTokenResource {
	return &AccessTokenResource{
		AccessToken: *oauthSession.AccessToken,
		Expires: oauthSession.AccessExpiresAt.Format(time.RFC1123Z),
	}
}