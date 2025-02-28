// https://jlunz.github.io/homeassistant/#/api/getApiV2Powermeter for reference

package batteries

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/pbtrad/go_batteries/proto/batteries/v1/sonnen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BatteryState struct {
	ApparentOutput            int
	BackupBuffer              int
	BatteryCharging           bool
	BatteryDischarging        bool
	Consumption_W             int
	Fac                       float64
	FlowConsumptionBattery    bool
	FlowConsumptionGrid       bool
	FlowConsumptionProduction bool
	GridFeedIn_W              int
	IsSystemInstalled         int
	OperatingMode             int
	PacTotal_W                int
	Production_W              int
	RSOC                      float64
	SystemStatus              string
	Timestamp                 string
}

type BatteryServer struct {
	pb.UnimplementedSonnenBatteriesServiceServer
	batteryState BatteryState
	batteryLock  sync.Mutex
}

func NewBatteryServer() *BatteryServer {
	return &BatteryServer{
		batteryState: BatteryState{
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
		},
	}
}

func (s *BatteryServer) GetSonnenStatus(ctx context.Context, req *pb.GetSonnenStatusRequest) (*pb.GetSonnenStatusResponse, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	s.batteryState.Timestamp = time.Now().Format(time.RFC3339)
	return &pb.GetSonnenStatusResponse{
		ApparentOutput:            int32(s.batteryState.ApparentOutput),
		BackupBuffer:              int32(s.batteryState.BackupBuffer),
		BatteryCharging:           s.batteryState.BatteryCharging,
		BatteryDischarging:        s.batteryState.BatteryDischarging,
		ConsumptionW:              int32(s.batteryState.Consumption_W),
		Fac:                       s.batteryState.Fac,
		FlowConsumptionBattery:    s.batteryState.FlowConsumptionBattery,
		FlowConsumptionGrid:       s.batteryState.FlowConsumptionGrid,
		FlowConsumptionProduction: s.batteryState.FlowConsumptionProduction,
		GridFeedInW:               int32(s.batteryState.GridFeedIn_W),
		IsSystemInstalled:         int32(s.batteryState.IsSystemInstalled),
		OperatingMode:             int32(s.batteryState.OperatingMode),
		PacTotalW:                 int32(s.batteryState.PacTotal_W),
		ProductionW:               int32(s.batteryState.Production_W),
		Rsoc:                      s.batteryState.RSOC,
		SystemStatus:              s.batteryState.SystemStatus,
		Timestamp:                 s.batteryState.Timestamp,
	}, nil
}

func (s *BatteryServer) GetLatestData(ctx context.Context, req *pb.GetLatestDataRequest) (*pb.GetLatestDataResponse, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	s.batteryState.Timestamp = time.Now().Format(time.RFC3339)
	return &pb.GetLatestDataResponse{
		Rsoc:         s.batteryState.RSOC,
		ConsumptionW: int32(s.batteryState.Consumption_W),
		ProductionW:  int32(s.batteryState.Production_W),
		GridFeedInW:  int32(s.batteryState.GridFeedIn_W),
		Timestamp:    s.batteryState.Timestamp,
	}, nil
}

func (s *BatteryServer) GetPowerMeterData(ctx context.Context, req *pb.GetPowerMeterDataRequest) (*pb.GetPowerMeterDataResponse, error) {

	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	current := float64(s.batteryState.Consumption_W) / 230.0
	totalPower := float64(s.batteryState.Consumption_W + s.batteryState.Production_W)

	log.Printf("Preparing response with current: %f, totalPower: %f", current, totalPower)

	response := &pb.GetPowerMeterDataResponse{
		VoltageL1:  230.0,
		CurrentL1:  current,
		PowerL1:    float64(s.batteryState.Consumption_W),
		VoltageL2:  230.0,
		CurrentL2:  current * 0.8,
		PowerL2:    float64(s.batteryState.Production_W),
		VoltageL3:  230.0,
		CurrentL3:  current * 1.1,
		PowerL3:    float64(s.batteryState.Production_W),
		TotalPower: totalPower,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	return response, nil
}

func (s *BatteryServer) GetEnergyData(ctx context.Context, req *pb.GetEnergyDataRequest) (*pb.GetEnergyDataResponse, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	return &pb.GetEnergyDataResponse{
		TotalConsumptionKwh:     1500.5,
		TotalProductionKwh:      1300.2,
		TotalGridFeedInKwh:      100.8,
		TotalGridConsumptionKwh: 200.5,
		Timestamp:               time.Now().Format(time.RFC3339),
	}, nil
}

func (s *BatteryServer) ChargeSonnen(ctx context.Context, req *pb.ChargeSonnenRequest) (*pb.ChargeSonnenResponse, error) {
	if req.Watt <= 0 {
		return nil, status.Error(codes.InvalidArgument, "watt value must be positive")
	}

	go s.simulateCharging(int(req.Watt))
	return &pb.ChargeSonnenResponse{
		Success: true,
		Message: "Charging Sonnen battery",
		Watt:    req.Watt,
	}, nil
}

func (s *BatteryServer) DischargeSonnen(ctx context.Context, req *pb.DischargeSonnenRequest) (*pb.DischargeSonnenResponse, error) {
	if req.Watt <= 0 {
		return nil, status.Error(codes.InvalidArgument, "watt value must be positive")
	}

	go s.simulateDischarging(int(req.Watt))
	return &pb.DischargeSonnenResponse{
		Success: true,
		Message: "Discharging Sonnen battery",
		Watt:    req.Watt,
	}, nil
}

func (s *BatteryServer) simulateCharging(watt int) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	if s.batteryState.RSOC >= 100 {
		s.batteryState.BatteryCharging = false
		return
	}

	s.batteryState.BatteryCharging = true
	for s.batteryState.RSOC < 100 {
		s.batteryState.RSOC += float64(watt) / 10000
		if s.batteryState.RSOC > 100 {
			s.batteryState.RSOC = 100
		}
		time.Sleep(1 * time.Second)
	}
	s.batteryState.BatteryCharging = false
}

func (s *BatteryServer) simulateDischarging(watt int) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	if s.batteryState.RSOC <= 0 {
		s.batteryState.BatteryDischarging = false
		return
	}

	s.batteryState.BatteryDischarging = true
	for s.batteryState.RSOC > 0 {
		s.batteryState.RSOC -= float64(watt) / 10000
		if s.batteryState.RSOC < 0 {
			s.batteryState.RSOC = 0
		}
		time.Sleep(1 * time.Second)
	}
	s.batteryState.BatteryDischarging = false
}
