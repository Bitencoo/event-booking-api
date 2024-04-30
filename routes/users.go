package routes

import (
	"fmt"
	"net/http"

	"github.com/Bitencoo/event-booking-api.git/models"
	"github.com/Bitencoo/event-booking-api.git/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Provided data for User is wrong!"})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Error saving User to the db!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "User created succesfully!", "user" : user,})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message" : "Error parsing user login!"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message" : "Wrong password!"})
		return
	}

	fmt.Print(user)
	token, err := utils.GenerateJWT(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message" : "Failed to generate JWT!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Login succesfull!", "token" : token})
}