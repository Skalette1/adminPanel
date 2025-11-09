package routes

import (
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *handlers.UserHandler) {
	r.POST("/users", h.CreateUserHandler)
	r.GET("/users", h.GetAllUsersHandler)
	r.GET("/users/:id", h.GetUserByIDHandler)
	r.PUT("/users/:id", h.UpdateUserByIDHandler)
	r.DELETE("/users/:id", h.DeleteUserByIDHandler)
}
