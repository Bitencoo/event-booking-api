package routes

import (
	"net/http"
	"strconv"

	"github.com/Bitencoo/event-booking-api.git/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID!"})
		return
	}

	var event *models.Event
	event, err = models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Invalid ID"})
		return
	}

	registrationID, err := event.Register(userId)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Error registering!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration created succesfully", "registrationID": registrationID})
	
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID!"})
		return
	}

	var event *models.Event
	event, err = models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve event from the database!"})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error canceling the registration!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration Canceled!"})
}