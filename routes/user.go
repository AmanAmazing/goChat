package routes

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/AmanAmazing/goChat/middlewares"
	"github.com/AmanAmazing/goChat/services"
	"github.com/AmanAmazing/goChat/utils"
	"github.com/AmanAmazing/goChat/views/components"
	"github.com/AmanAmazing/goChat/views/pages"
	"github.com/AmanAmazing/goChat/views/partials"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()

	// for routes logged in users should not see
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		//  TODO: need to add logged in redirector
		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			pages.GetLogin().Render(context.Background(), w)
		})
	})

	// public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Welcome to the public router"))
		})

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			username := r.FormValue("username")
			password := r.FormValue("password")

			// getting the jwt token
			jwtToken, err := services.PostLogin(username, password)
			if err != nil {
				if errors.Is(err, services.ErrInvalidCredentials) {
					w.WriteHeader(http.StatusOK) // FIX:Need to use HTMX Response extension to apply correct http code (unauthorized)
					components.ErrorLogin().Render(context.Background(), w)
					return
				} else {
					// TODO: need to change the error message that is displayed
					w.WriteHeader(http.StatusOK) // FIX:Need to use HTMX Response extension to apply correct http code (internal server error)
					components.ErrorLogin().Render(context.Background(), w)
					return
				}
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    jwtToken,
				HttpOnly: true,
				// secure: true, // TODO: set to true if using https
				SameSite: http.SameSiteStrictMode,
				Expires:  time.Now().Add(7 * 24 * time.Hour),
			})
			w.Header().Add("HX-Retarget", "body")
			w.Header().Add("HX-Reswap", "innerHTML")
			w.Header().Add("HX-Push-Url", "/home")
			partials.GetHome().Render(context.Background(), w)
		})

	})

	// For routes that the user needs to be logged in for
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(middlewares.UnloggedInRedirector)

		r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
			if w.Header().Get("HX-Request") != "" {
				w.Header().Add("HX-Retarget", "body")
				w.Header().Add("HX-Reswap", "innerHTML")
				w.Header().Add("HX-Push-Url", "/home")
				partials.GetHome().Render(context.Background(), w)
				return
			}
			pages.GetHome().Render(context.Background(), w)
		})
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    "",
				HttpOnly: true,
				Path:     "/",
				// secure:   true, // set to true if using https
				SameSite: http.SameSiteStrictMode,
				MaxAge:   -1,
				Expires:  time.Unix(0, 0),
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
		})

		r.Get("/communities", func(w http.ResponseWriter, r *http.Request) {
			services.GetCommunities()
			pages.GetCommunities().Render(context.Background(), w)
		})

	})

	// returning http handler
	return r
}
