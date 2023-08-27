package handlers

import (
	"github.com/go-chi/chi/v5"
)

func (h *Handlers) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Route("/api/user", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
	})
	return router
}
