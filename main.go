package main

import (
	"net/http"

	"github.com/Bitencoo/event-booking-api.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/events", getEvents)
  r.POST("/events", createEvent)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse the data!"})
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	c.JSON(http.StatusCreated, gin.H{
		"message" : "Event Created!",
		"event" : event,
	})
}