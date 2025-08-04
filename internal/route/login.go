package route

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ajg/form"
	"github.com/bdreece/ephemera/internal/security"
	"github.com/bdreece/ephemera/pkg/database"
	"github.com/bdreece/ephemera/pkg/identity"
	"github.com/golang-jwt/jwt/v5"
)

type loginRequest struct {
	Username string            `form:"username"`
	Password identity.Password `form:"password"`
}

func (route *IdentityRoute) Login(w http.ResponseWriter, r *http.Request) {
	var input loginRequest
	if err := form.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := database.FindUserByDisplayNameParams{
		DisplayName: input.Username,
	}

	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	q := database.New(route.db)
	user, err := q.FindUserByDisplayName(ctx, query)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	hash := security.MustDecode[security.Hash](user.PasswordHash)
	salt := security.MustDecode[security.Salt](user.PasswordSalt)
	if !hash.Compare(input.Password, salt) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := identity.AccessClaims{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprint(user.ID),
			Issuer:    "ephemera.bdreece.dev",
			Audience:  jwt.ClaimStrings{"ephemera.bdreece.dev"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token, err := route.accessToken.Sign(claims)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := tokenResponse{token}
	if err = json.NewEncoder(w).Encode(output); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
