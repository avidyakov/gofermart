package routers

import (
	"github.com/go-chi/chi/v5"
	"gophermart/cmd/handlers/users"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Route("/api/user", func(r chi.Router) {
		r.Post("/register", users.Register)
		r.Post("/login", users.Login)
	})
	return router
}
