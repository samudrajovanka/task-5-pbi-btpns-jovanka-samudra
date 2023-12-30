package main

import (
	"photo-api/database"
	"photo-api/helpers"
	"photo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	helpers.LoadEnv()

	r := gin.Default()

	database.Init()

	routes.Router(r)

	r.Run()
}
