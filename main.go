// @title Admin Panel API
// @version 1.0
// @description API для работы с пользователями и ролями.
// @host localhost:8080
// @BasePath /
package main

import (
	"log"

	_ "github.com/Skalette1/adminPanel/docs" // This line is important
	"github.com/Skalette1/adminPanel/internal/db"
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/Skalette1/adminPanel/internal/repository"
	"github.com/Skalette1/adminPanel/internal/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Admin Panel API
// @version 1.0
// @description This is an admin panel server.
// @termsOfService http://swagger.io/terms/

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	roleRepo := repository.NewRoleRepository(database)
	userRepo := repository.NewUserRepository(database)

	roleHandler := handlers.NewRoleHandler(roleRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RoleRoutes(r, roleHandler)
	routes.UserRoutes(r, userHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
