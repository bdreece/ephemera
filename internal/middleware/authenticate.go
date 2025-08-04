package middleware

import (
	"net/http"
	"strings"

	"github.com/bdreece/ephemera/pkg/identity"
)

type Authenticator func(next http.Handler) http.Handler

func Authenticate(verifier identity.AccessTokenVerifier) Authenticator {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			token := strings.TrimPrefix(header, "Bearer ")
			claims, err := verifier.Verify(token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			identity.SetClaims(r, claims)
			next.ServeHTTP(w, r)
		})
	}
}
