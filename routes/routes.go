package routes

import (
	"github.com/Bitencoo/event-booking-api.git/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	

	server.POST("/signup", signup)
	server.POST("/login", login)
}