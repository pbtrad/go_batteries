// https://jlunz.github.io/homeassistant/#/api/getApiV2Powermeter for reference
// Statis for now, will try add dynamic updates with battery state tracking in future

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Battery status
	router.GET("/api/v2/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Apparent_output":           226,
			"BackupBuffer":              0,
			"BatteryCharging":           false,
			"BatteryDischarging":        false,
			"Consumption_W":             232,
			"Fac":                       49.999,
			"FlowConsumptionBattery":    false,
			"FlowConsumptionGrid":       true,
			"FlowConsumptionProduction": true,
			"FlowGridBattery":           false,
			"FlowProductionBattery":     false,
			"FlowProductionGrid":        false,
			"GridFeedIn_W":              -208,
			"IsSystemInstalled":         1,
			"OperatingMode":             2,
			"Pac_total_W":               -5,
			"Production_W":              28,
			"RSOC":                      4,
			"Sac1":                      75,
			"Sac2":                      75,
			"Sac3":                      76,
			"SystemStatus":              "OnGrid",
			"Timestamp":                 "2025-02-07 13:17:08",
			"USOC":                      0,
			"Uac":                       230,
			"Ubat":                      50,
			"dischargeNotAllowed":       false,
			"generator_autostart":       false,
		})
	})

	// Latest dtaa
	router.GET("/api/v2/latestdata", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Consumption_W": 232,
			"Production_W":  28,
			"RSOC":          4,
			"USOC":          0,
			"GridFeedIn_W":  -208,
			"Pac_total_W":   -5,
			"Timestamp":     "2025-02-07 13:17:08",
		})
	})

	// Set the charge setpoint
	router.POST("/api/v2/setpoint/charge/:watt", func(c *gin.Context) {
		watt := c.Param("watt")
		// Charge command stuff
		c.JSON(http.StatusOK, gin.H{
			"message": "Charge command received",
			"watt":    watt,
		})
	})

	// Set the discharge setpoint
	router.POST("/api/v2/setpoint/discharge/:watt", func(c *gin.Context) {
		watt := c.Param("watt")
		// Discharge command stuff
		c.JSON(http.StatusOK, gin.H{
			"message": "Discharge command received",
			"watt":    watt,
		})
	})

	router.Run(":8080")
}
