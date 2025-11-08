package routes

import (
	userapi "github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/users/{id}", userapi.GetRoleById)
	r.Get("/users", userapi.GetAllRoles)
	r.Post("/users", userapi.CreateRole)
	r.Put("/users/{id}", userapi.UpdateRole)
	r.Delete("/users{id}", userapi.DeleteRole)
	return r
}
