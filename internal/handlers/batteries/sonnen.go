// https://jlunz.github.io/homeassistant/#/api/getApiV2Powermeter for reference
// Statis for now, will try add dynamic updates with battery state tracking in future

package batteries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Returns mock status
func GetSonnenStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"RSOC":               40.5,
		"BatteryCharging":    true,
		"BatteryDischarging": false,
		"Consumption_W":      500,
		"Production_W":       200,
		"GridFeedIn_W":       -150,
		"SystemStatus":       "OnGrid",
		"Timestamp":          "2025-02-07 14:00:00",
	})
}

// Simulate charging
func ChargeSonnen(c *gin.Context) {
	watt := c.Param("watt")
	c.JSON(http.StatusOK, gin.H{
		"message": "Charging Sonnen battery",
		"watt":    watt,
	})
}

// Simulate discharging
func DischargeSonnen(c *gin.Context) {
	watt := c.Param("watt")
	c.JSON(http.StatusOK, gin.H{
		"message": "Discharging Sonnen battery",
		"watt":    watt,
	})
}
