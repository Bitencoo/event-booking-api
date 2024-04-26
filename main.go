package main

import (
	"github.com/Bitencoo/event-booking-api.git/db"
	"github.com/Bitencoo/event-booking-api.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
  db.InitDB()
  r := gin.Default()
  routes.RegisterRoutes(r)
  r.Run()
}