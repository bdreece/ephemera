package ephemera

import (
	"net/http/httputil"
	"net/url"

	"github.com/bdreece/ephemera/internal/route"
	"github.com/bdreece/ephemera/pkg/identity"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func newRouter(p *Params) chi.Router {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.GetHead,
	)

	accessHandler := identity.NewAccessTokenHandler(p.JWTSecret)
	refreshHandler := identity.NewRefreshTokenHandler(p.JWTSecret)
	identity := route.Identity(route.IdentityOptions{
		DB:                  p.DB,
		AccessTokenHandler:  accessHandler,
		RefreshTokenHandler: refreshHandler,
	})

	r.Post("/login", identity.Login)

	if p.Environment.IsDevelopment() {
		url, _ := url.Parse("http://localhost:5173")
		revproxy := httputil.NewSingleHostReverseProxy(url)
		r.Get("/*", revproxy.ServeHTTP)
	}

	return r
}
