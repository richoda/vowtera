package routes

import (
	"net/http"

	"admin-api/handlers"
	"admin-api/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func Setup(secret string) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.RequestID)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", handlers.Login)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(secret))
			r.Get("/me", handlers.Me)
		})
	})

	return r
}
