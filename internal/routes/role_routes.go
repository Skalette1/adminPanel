package routes

import (
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RoleRoutes(h *handlers.RoleHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	return r
}
