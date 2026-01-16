package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	server.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Health check successful",
		})
	})

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
