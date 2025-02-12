// https://jlunz.github.io/homeassistant/#/api/getApiV2Powermeter for reference
// Statis for now, will try add dynamic updates with battery state tracking in future

package batteries

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type BatteryState struct {
	ApparentOutput            int     `json:"Apparent_output"`
	BackupBuffer              int     `json:"BackupBuffer"`
	BatteryCharging           bool    `json:"BatteryCharging"`
	BatteryDischarging        bool    `json:"BatteryDischarging"`
	Consumption_W             int     `json:"Consumption_W"`
	Fac                       float64 `json:"Fac"`
	FlowConsumptionBattery    bool    `json:"FlowConsumptionBattery"`
	FlowConsumptionGrid       bool    `json:"FlowConsumptionGrid"`
	FlowConsumptionProduction bool    `json:"FlowConsumptionProduction"`
	GridFeedIn_W              int     `json:"GridFeedIn_W"`
	IsSystemInstalled         int     `json:"IsSystemInstalled"`
	OperatingMode             int     `json:"OperatingMode"`
	PacTotal_W                int     `json:"Pac_total_W"`
	Production_W              int     `json:"Production_W"`
	RSOC                      float64 `json:"RSOC"`
	SystemStatus              string  `json:"SystemStatus"`
	Timestamp                 string  `json:"Timestamp"`
}

var batteryState = BatteryState{
	ApparentOutput:            226,
	BackupBuffer:              0,
	BatteryCharging:           false,
	BatteryDischarging:        false,
	Consumption_W:             232,
	Fac:                       49.999,
	FlowConsumptionBattery:    false,
	FlowConsumptionGrid:       true,
	FlowConsumptionProduction: true,
	GridFeedIn_W:              -208,
	IsSystemInstalled:         1,
	OperatingMode:             2,
	PacTotal_W:                -5,
	Production_W:              28,
	RSOC:                      40.5,
	SystemStatus:              "OnGrid",
	Timestamp:                 time.Now().Format(time.RFC3339),
}

var batteryLock sync.Mutex

// Battery status
func GetSonnenStatus(c *gin.Context) {
	batteryLock.Lock()
	defer batteryLock.Unlock()

	batteryState.Timestamp = time.Now().Format(time.RFC3339)
	c.JSON(http.StatusOK, batteryState)
}

// Latest data
func GetLatestData(c *gin.Context) {
	batteryLock.Lock()
	defer batteryLock.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"RSOC":          batteryState.RSOC,
		"Consumption_W": batteryState.Consumption_W,
		"Production_W":  batteryState.Production_W,
		"GridFeedIn_W":  batteryState.GridFeedIn_W,
		"Timestamp":     batteryState.Timestamp,
	})
}

// Simulates charging
func ChargeSonnen(c *gin.Context) {
	watt, err := strconv.Atoi(c.Param("watt"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watt value"})
		return
	}
	go simulateCharging(watt)
	c.JSON(http.StatusOK, gin.H{"message": "Charging Sonnen battery", "watt": watt})
}

// Simulates discharging
func DischargeSonnen(c *gin.Context) {
	watt, err := strconv.Atoi(c.Param("watt"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watt value"})
		return
	}
	go simulateDischarging(watt)
	c.JSON(http.StatusOK, gin.H{"message": "Discharging Sonnen battery", "watt": watt})
}

// Mock power meter data
func GetPowerMeterData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Voltage_L1": 230.0,
		"Current_L1": 10.5,
		"Power_L1":   2415,
		"Voltage_L2": 230.0,
		"Current_L2": 8.0,
		"Power_L2":   1840,
		"Voltage_L3": 230.0,
		"Current_L3": 9.2,
		"Power_L3":   2116,
		"TotalPower": 6371,
		"Timestamp":  time.Now().Format(time.RFC3339),
	})
}

// Returns mock historical energy data
func GetEnergyData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"TotalConsumption_kWh":     1500.5,
		"TotalProduction_kWh":      1300.2,
		"TotalGridFeedIn_kWh":      100.8,
		"TotalGridConsumption_kWh": 200.5,
		"Timestamp":                time.Now().Format(time.RFC3339),
	})
}

// Simulate charging the battery
func simulateCharging(watt int) {
	batteryLock.Lock()
	defer batteryLock.Unlock()

	batteryState.BatteryCharging = true
	for i := 0; i < 10; i++ {
		if batteryState.RSOC >= 100 {
			batteryState.RSOC = 100
			break
		}
		batteryState.RSOC += float64(watt) / 10000
		time.Sleep(1 * time.Second)
	}
	batteryState.BatteryCharging = false
}

// Simulate discharging the battery
func simulateDischarging(watt int) {
	batteryLock.Lock()
	defer batteryLock.Unlock()

	batteryState.BatteryDischarging = true
	for i := 0; i < 10; i++ {
		if batteryState.RSOC <= 0 {
			batteryState.RSOC = 0
			break
		}
		batteryState.RSOC -= float64(watt) / 10000
		time.Sleep(1 * time.Second)
	}
	batteryState.BatteryDischarging = false
}
