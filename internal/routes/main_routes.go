package routes

import (
	"github.com/go-chi/chi/v5"
)

func InitRoutes() chi.Router {
	r := chi.NewRouter()
	r.Mount("/api", RoleRoutes())
	r.Mount("/api", UserRoutes())
	return r
}
