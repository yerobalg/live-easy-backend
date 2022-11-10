package infrastructure

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuth struct {
	Config *oauth2.Config
}

func (o *OAuth) Init() *OAuth {
	return &OAuth{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			RedirectURL:  fmt.Sprintf("%s/auth/login/google/callback", os.Getenv("APP_BASE_URL")),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}
