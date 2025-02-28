package batteries

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/pbtrad/go_batteries/proto/batteries/v1/givenergy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GivEnergyState struct {
	SystemSerialNumber string
	GridPower          float64
	SolarPower         float64
	BatteryPower       float64
	BatterySOC         float64
	BatteryCharging    bool
	BatteryDischarging bool
	Temperature        float64
	SystemStatus       string
	Timestamp          string
}

type GivEnergyBatteryServer struct {
	pb.UnimplementedGivEnergyServiceServer
	batteryState GivEnergyState
	batteryLock  sync.Mutex
}

func NewGivEnergyBatteryServer() *GivEnergyBatteryServer {
	return &GivEnergyBatteryServer{
		batteryState: GivEnergyState{
			SystemSerialNumber: "GIV-2025-12345",
			GridPower:          500.0,
			SolarPower:         1200.0,
			BatteryPower:       -100.0,
			BatterySOC:         60.0,
			BatteryCharging:    true,
			BatteryDischarging: false,
			Temperature:        25.5,
			SystemStatus:       "Normal",
			Timestamp:          time.Now().Format(time.RFC3339),
		},
	}
}

// Battery status
func (s *GivEnergyBatteryServer) GetBatteryState(ctx context.Context, req *pb.GetBatteryStateRequest) (*pb.GetBatteryStateResponse, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	s.batteryState.Timestamp = time.Now().Format(time.RFC3339)
	return &pb.GetBatteryStateResponse{
		SystemSerialNumber: s.batteryState.SystemSerialNumber,
		BatterySoc:         float32(s.batteryState.BatterySOC),
		BatteryCharging:    s.batteryState.BatteryCharging,
		BatteryDischarging: s.batteryState.BatteryDischarging,
		SystemStatus:       s.batteryState.SystemStatus,
		Timestamp:          s.batteryState.Timestamp,
	}, nil
}

// Simulates battery charging
func (s *GivEnergyBatteryServer) ChargeBattery(ctx context.Context, req *pb.ChargeBatteryRequest) (*pb.ChargeBatteryResponse, error) {
	if req.Watt <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Charge wattage must be positive")
	}

	go s.simulateCharging(int(req.Watt))
	log.Printf("Charging battery with %dW", req.Watt)

	return &pb.ChargeBatteryResponse{
		Success: true,
		Message: "Charging GivEnergy battery",
		Watt:    req.Watt,
	}, nil
}

// Simulates battery discharging
func (s *GivEnergyBatteryServer) DischargeBattery(ctx context.Context, req *pb.DischargeBatteryRequest) (*pb.DischargeBatteryResponse, error) {
	if req.Watt <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Discharge wattage must be positive")
	}

	go s.simulateDischarging(int(req.Watt))
	log.Printf("Discharging battery with %dW", req.Watt)

	return &pb.DischargeBatteryResponse{
		Success: true,
		Message: "Discharging GivEnergy battery",
		Watt:    req.Watt,
	}, nil
}

// Current power readings
func (s *GivEnergyBatteryServer) GetPowerData(ctx context.Context, req *pb.GetPowerFlowRequest) (*pb.GetPowerFlowResponse, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	return &pb.GetPowerFlowResponse{
		GridPower:    float32(s.batteryState.GridPower),
		SolarPower:   float32(s.batteryState.SolarPower),
		BatteryPower: float32(s.batteryState.BatteryPower),
		Timestamp:    s.batteryState.Timestamp,
	}, nil
}

// Battery temperature
func (s *GivEnergyBatteryServer) GetBatteryTemperature(ctx context.Context, req *pb.GetBatteryTemperatureRequest) (*pb.GetBatteryTemperatureResponse, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	return &pb.GetBatteryTemperatureResponse{
		Temperature: float32(s.batteryState.Temperature),
		Timestamp:   s.batteryState.Timestamp,
	}, nil
}

// Simulate battery charging
func (s *GivEnergyBatteryServer) simulateCharging(watt int) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	s.batteryState.BatteryCharging = true
	for s.batteryState.BatterySOC < 100 {
		s.batteryState.BatterySOC += float64(watt) / 10000
		if s.batteryState.BatterySOC > 100 {
			s.batteryState.BatterySOC = 100
		}
		time.Sleep(1 * time.Second)
	}
	s.batteryState.BatteryCharging = false
}

// Simulate battery discharging
func (s *GivEnergyBatteryServer) simulateDischarging(watt int) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	s.batteryState.BatteryDischarging = true
	for s.batteryState.BatterySOC > 0 {
		s.batteryState.BatterySOC -= float64(watt) / 10000
		if s.batteryState.BatterySOC < 0 {
			s.batteryState.BatterySOC = 0
		}
		time.Sleep(1 * time.Second)
	}
	s.batteryState.BatteryDischarging = false
}

func (s *GivEnergyBatteryServer) GetLatestData(ctx context.Context, req *pb.GetLatestDataRequest) (*pb.GetLatestDataResponse, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	return &pb.GetLatestDataResponse{
		BatterySoc:              float32(s.batteryState.BatterySOC),
		PowerFlowBatteryToLoad:  float32(s.batteryState.BatteryPower),
		PowerFlowSolarToBattery: float32(s.batteryState.SolarPower),
		PowerFlowGridToBattery:  float32(s.batteryState.GridPower),
		PowerFlowBatteryToGrid:  0.0,
		Timestamp:               s.batteryState.Timestamp,
	}, nil
}
