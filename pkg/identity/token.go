package identity

import (
	"context"
	"flag"
	"net/http"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"

	"github.com/bdreece/ephemera/internal/security"
)

// AccessClaims are the claims embedded in an access token.
type AccessClaims struct {
	jwt.RegisteredClaims

	FirstName   string `json:"given_name"`
	LastName    string `json:"family_name"`
	DisplayName string `json:"preferred_username"`
	AvatarURL   string `json:"picture"`
}

// Parses the user ID from the Subject field.
func (claims AccessClaims) UserID() (int, error) {
	return strconv.Atoi(claims.Subject)
}

type AccessTokenSigner interface {
	security.TokenSigner[AccessClaims]
}

type AccessTokenVerifier interface {
	security.TokenVerifier[AccessClaims]
}

// An AccessTokenHandler is a [security.TokenHandler] for [AccessClaims].
type AccessTokenHandler interface {
	AccessTokenSigner
	AccessTokenVerifier
}

type TokenConfig struct {
	Issuer   string
	Audience []string
	Lifetime time.Duration
}

func NewAccessTokenHandler(secret security.Key) AccessTokenHandler {
	const lifetime time.Duration = 24 * time.Hour

	aud := jwt.ClaimStrings{flag.Lookup("jwt.aud").Value.String()}
	iss := flag.Lookup("jwt.iss").Value.String()

	return security.NewJWTHandler[AccessClaims](secret,
		security.SignJWTWith(func(t *jwt.Token) {
			claims := t.Claims.(*AccessClaims)
			uuid, _ := uuid.NewV4()
			claims.ID = uuid.String()
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(lifetime))
			claims.Audience = aud
			claims.Issuer = iss
		}),
		security.VerifyJWTWith(
			jwt.WithAllAudiences(aud...),
			jwt.WithIssuer(iss),
			jwt.WithIssuedAt(),
			jwt.WithExpirationRequired(),
		),
	)
}

// RefreshClaims are the claims embedded in a refresh token.
type RefreshClaims struct {
	jwt.RegisteredClaims

	AccessTokenMD5 string `json:"at_hash"`
}

// A RefreshTokenHandler is a [security.TokenHandler] for [RefreshClaims].
type RefreshTokenHandler interface {
	security.TokenHandler[RefreshClaims]
}

func NewRefreshTokenHandler(secret security.Key) RefreshTokenHandler {
	const lifetime time.Duration = 7 * 24 * time.Hour

	aud := jwt.ClaimStrings{flag.Lookup("jwt.aud").Value.String()}
	iss := flag.Lookup("jwt.iss").Value.String()

	return security.NewJWTHandler[RefreshClaims](secret,
		security.SignJWTWith(func(t *jwt.Token) {
			claims := t.Claims.(*AccessClaims)
			uuid, _ := uuid.NewV4()
			claims.ID = uuid.String()
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(lifetime))
			claims.Audience = aud
			claims.Issuer = iss
		}),
		security.VerifyJWTWith(
			jwt.WithAllAudiences(aud...),
			jwt.WithIssuer(iss),
			jwt.WithIssuedAt(),
			jwt.WithExpirationRequired(),
		),
	)
}

type claimsKey int

const (
	accessKey claimsKey = iota
	refreshKey
)

type Claims interface {
	*AccessClaims | *RefreshClaims
}

func GetClaims[C Claims](r *http.Request) (claims C, ok bool) {
	var key claimsKey
	switch any(claims).(type) {
	case *AccessClaims:
		key = accessKey
	case *RefreshClaims:
		key = refreshKey
	default:
		panic("invalid claims type")
	}

	claims, ok = r.Context().Value(key).(C)
	return
}

func SetClaims[C Claims](r *http.Request, claims C) {
	var key claimsKey
	switch any(claims).(type) {
	case *AccessClaims:
		key = accessKey
	case *RefreshClaims:
		key = refreshKey
	default:
		panic("invalid claims type")
	}

	*r = *r.WithContext(context.WithValue(r.Context(), key, claims))
}
