package handlers

import (
	"github.com/go-chi/chi/v5"
)

func (h *Handlers) NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Route("/api/user", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Post("/orders", h.CreateOrder)
		r.Get("/orders", h.GetOrders)
		r.Get("/balance", h.GetBalance)
	})
	return router
}
