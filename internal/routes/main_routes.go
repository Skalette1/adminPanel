package routes

import (
	"github.com/Skalette1/adminPanel/docs"
	"github.com/Skalette1/adminPanel/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(userHandler *handlers.UserHandler, roleHandler *handlers.RoleHandler) *gin.Engine {
	r := gin.Default()

	// Swagger configuration
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	RoleRoutes(r, roleHandler)
	UserRoutes(r, userHandler)

	return r
}
