package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Bitencoo/event-booking-api.git/models"
	"github.com/Bitencoo/event-booking-api.git/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, events)
}

func getEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing the int ID provided"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found!"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message" : "Not authorized!"})
		return
	}

	err := utils.ValidateToken(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Not authorized!"})
		return
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse the data!"})
		return
	}

	event.ID = 1	
	event.UserID = 1
	err = event.Save()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the new Event, try again later!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message" : "Event Created!",
		"event" : event,
	})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing the int ID provided"})
		return
	}

	_, err = models.GetEventById(id)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the Event!"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data submitted!"})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.UpdateEvent()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update Event!"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message" : "Event updated succesfully!", "event" : updatedEvent})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing the int ID provided"})
		return
	}

	_, err = models.GetEventById(id)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the Event!"})
		return
	}

	err = models.DeleteEvent(id)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the Event!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Deleted the Event succesfully!"})
}