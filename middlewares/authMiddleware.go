package middlewares

import (
	"photo-api/app"
	"photo-api/database"
	"photo-api/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer helpers.RecoverErrorResponse(ctx)

		bearerToken := ctx.GetHeader("Authorization")

		if bearerToken == "" {
			helpers.UnauthorizedError("You not authorized")
		}

		accessToken := strings.Split(bearerToken, " ")[1]

		tokenClaims := helpers.ParseToken(accessToken)

		var user app.User

		if err := database.DB.Where("username=?", tokenClaims.Username).First(&user).Error; err != nil {
			helpers.UnauthorizedError("You not authorized")
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
