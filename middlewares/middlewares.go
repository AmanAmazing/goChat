package middlewares

import (
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
)

func UnloggedInRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, _ := jwtauth.FromContext(r.Context())
		if token == nil || jwt.Validate(token) != nil {
			http.Redirect(w, r, "/login", 302)
		}

		next.ServeHTTP(w, r)
	})
}
