package routes

import (
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RoleRoutes(r *gin.Engine, h *handlers.RoleHandler) {
	r.POST("/roles", h.Create)
	r.GET("/roles", h.GetAll)
	r.GET("/roles/:id", h.GetByID)
	r.PUT("/roles/:id", h.Update)
	r.DELETE("/roles/:id", h.Delete)
}
