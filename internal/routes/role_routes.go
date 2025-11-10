package routes

import (
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RoleRoutes(r *gin.Engine, h *handlers.RoleHandler) {
	r.POST("/", h.CreateRoleHandler)
	r.GET("/", h.GetAllRolesHandler)
	r.GET("/{id}", h.GetRoleByIDHandler)
	r.PUT("/{id}", h.UpdateRoleHandler)
	r.DELETE("/{id}", h.DeleteRoleHandler)
}
