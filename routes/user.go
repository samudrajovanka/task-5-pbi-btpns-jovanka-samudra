package routes

import (
	"photo-api/controllers"
	"photo-api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	userRoute := r.Group("/api/users")

	userRoute.POST("/register", controllers.UserRegister)
	userRoute.POST("/login", controllers.UserLogin)

	userRoute.Use(middlewares.AuthMiddleware())
	userRoute.PUT("/:userId", controllers.UpdateByIdUser)
	userRoute.DELETE("/:userId", controllers.DeleteByIdUser)
}
