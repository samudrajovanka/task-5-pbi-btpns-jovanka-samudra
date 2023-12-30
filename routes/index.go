package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "PONG")
	})

	UserRoute(r)
	PhotoRoute(r)
}
