package main

import (
	"log"

	_ "github.com/Skalette1/adminPanel/docs"
	"github.com/Skalette1/adminPanel/internal/db"
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/Skalette1/adminPanel/internal/repository"
	routes "github.com/Skalette1/adminPanel/internal/routes"
	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer database.Close()

	// Создание репозиториев
	roleRepo := repository.NewRoleRepository(database)
	userRepo := repository.NewUserRepository(database)

	// Создание обработчиков
	roleHandler := handlers.NewRoleHandler(roleRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	// Инициализация роутера
	r := routes.InitRouter(userHandler, roleHandler)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
