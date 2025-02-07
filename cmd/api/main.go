package main

import (
	"github.com/pbtrad/go_batteries/internal/handlers/batteries"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// SonnenBatterie API
	sonnen := router.Group("/api/sonnen")
	{
		sonnen.GET("/status", batteries.GetSonnenStatus)
		sonnen.POST("/setpoint/charge/:watt", batteries.ChargeSonnen)
		sonnen.POST("/setpoint/discharge/:watt", batteries.DischargeSonnen)
	}
	router.Run(":8080")
}
