package routes

import (
	"photo-api/controllers"
	"photo-api/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRoute(r *gin.Engine) {
	photoRoute := r.Group("/api/photos")

	photoRoute.Use(middlewares.AuthMiddleware())
	photoRoute.POST("/", controllers.CreatePhoto)
	photoRoute.GET("/", controllers.GetAllPhoto)
	photoRoute.GET("/:photoId", controllers.GetByIdPhoto)
	photoRoute.PUT("/:photoId", controllers.UpdateByIdPhoto)
	photoRoute.DELETE("/:photoId", controllers.DeleteByIdPhoto)
}
