package route

import (
	"github.com/bdreece/ephemera/pkg/database"
	"github.com/bdreece/ephemera/pkg/identity"
)

type IdentityRoute struct {
	db           database.DB
	accessToken  identity.AccessTokenHandler
	refreshToken identity.RefreshTokenHandler
}

type IdentityOptions struct {
	DB                  database.DB
	AccessTokenHandler  identity.AccessTokenHandler
	RefreshTokenHandler identity.RefreshTokenHandler
}

func Identity(opts IdentityOptions) *IdentityRoute {
	return &IdentityRoute{
		db:           opts.DB,
		accessToken:  opts.AccessTokenHandler,
		refreshToken: opts.RefreshTokenHandler,
	}
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}
