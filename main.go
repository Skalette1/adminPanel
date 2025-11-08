package main

import (
	"net/http"

	"github.com/Skalette1/adminPanel/internal/db"
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/Skalette1/adminPanel/internal/repository"
	"github.com/Skalette1/adminPanel/internal/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	roleRepo := repository.NewRoleRepository(database)
	roleHandler := handlers.NewRoleHandler(roleRepo)

	r := chi.NewRouter()
	r.Mount("/roles", routes.RoleRoutes(roleHandler))

	http.ListenAndServe(":8080", r)
}
