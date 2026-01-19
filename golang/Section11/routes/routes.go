package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/signup", signUpUser)
	server.POST("/login", loginUser)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	{
		authenticated.POST("/events", createEvent)
		authenticated.PUT("/events/:id", updateEvent)
		authenticated.DELETE("/events/:id", deleteEvent)
	}
}
