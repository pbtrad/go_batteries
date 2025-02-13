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

	batteryState.Timestamp = time.Now().Format(time.RFC3339)
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
	batteryLock.Lock()
	defer batteryLock.Unlock()

	current := float64(batteryState.Consumption_W) / 230.0
	totalPower := batteryState.Consumption_W + batteryState.Production_W

	c.JSON(http.StatusOK, gin.H{
		"Voltage_L1": 230.0,
		"Current_L1": current,
		"Power_L1":   batteryState.Consumption_W,
		"Voltage_L2": 230.0,
		"Current_L2": current * 0.8,
		"Power_L2":   batteryState.Production_W,
		"Voltage_L3": 230.0,
		"Current_L3": current * 1.1,
		"Power_L3":   batteryState.Production_W * 1,
		"TotalPower": totalPower,
		"Timestamp":  time.Now().Format(time.RFC3339),
	})
}

// Returns mock historical energy data
func GetEnergyData(c *gin.Context) {
	batteryLock.Lock()
	defer batteryLock.Unlock()

	batteryState.Timestamp = time.Now().Format(time.RFC3339)
	c.JSON(http.StatusOK, gin.H{
		"TotalConsumption_kWh":     1500.5,
		"TotalProduction_kWh":      1300.2,
		"TotalGridFeedIn_kWh":      100.8,
		"TotalGridConsumption_kWh": 200.5,
		"Timestamp":                batteryState.Timestamp,
	})
}

// Simulate charging the battery
func simulateCharging(watt int) {
	batteryLock.Lock()
	defer batteryLock.Unlock()

	if batteryState.RSOC >= 100 {
		batteryState.BatteryCharging = false
		return
	}

	batteryState.BatteryCharging = true
	for batteryState.RSOC < 100 {
		batteryState.RSOC += float64(watt) / 10000
		if batteryState.RSOC > 100 {
			batteryState.RSOC = 100
		}
		time.Sleep(1 * time.Second)
	}
	batteryState.BatteryCharging = false
}

// Simulate discharging the battery
func simulateDischarging(watt int) {
	batteryLock.Lock()
	defer batteryLock.Unlock()

	if batteryState.RSOC <= 0 {
		batteryState.BatteryDischarging = false
		return
	}

	batteryState.BatteryDischarging = true
	for batteryState.RSOC > 0 {
		batteryState.RSOC -= float64(watt) / 10000
		if batteryState.RSOC < 0 {
			batteryState.RSOC = 0
		}
		time.Sleep(1 * time.Second)
	}
	batteryState.BatteryDischarging = false
}
