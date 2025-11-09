package routes

import (
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RoleRoutes(r *gin.Engine, h *handlers.RoleHandler) {
	r.POST("/roles", h.CreateRole)
	r.GET("/roles", h.GetAllRoles)
	r.GET("/roles/:id", h.GetRoleByID)
	r.PUT("/roles/:id", h.UpdateRole)
	r.DELETE("/roles/:id", h.DeleteRole)
}
