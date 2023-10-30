package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"

	"projector-test-app/app/handlers"
)

func NewRouter(db *sqlx.DB) http.Handler {
	rootRouter := chi.NewRouter()

	rootRouter.Route("/", func(router chi.Router) {
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)

		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		userHandler := handlers.NewUser(db)
		router.Route("/", func(r chi.Router) {
			r.Post("/users", userHandler.HandleCreateUser)
		})
	})

	return rootRouter
}
