package middlewares

import (
	"net/http"

	"github.com/Bitencoo/event-booking-api.git/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Not authorized!"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : "Not authorized!"})
		return
	}

	c.Set("userId", userId)
	c.Next()
}