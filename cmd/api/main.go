package main

import (
	"log"
	"os"

	"github.com/pbtrad/go_batteries/internal/handlers/batteries"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// SonnenBatterie API
	sonnen := router.Group("/api/sonnen")
	{
		sonnen.GET("/status", batteries.GetSonnenStatus)
		sonnen.POST("/setpoint/charge/:watt", batteries.ChargeSonnen)
		sonnen.POST("/setpoint/discharge/:watt", batteries.DischargeSonnen)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API running on port %s", port)
	router.Run(":" + port)
}
