// https://developer.tesla.com/docs/fleet-api/endpoints/energy for reference

package batteries

import (
	"context"
	"sync"
	"time"

	pb "github.com/pbtrad/go_batteries/proto/batteries/v1/tesla"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeslaBatteryServer struct {
	pb.UnimplementedTeslaBatteryServiceServer
	batteryState *pb.TeslaBatteryState
	batteryLock  sync.Mutex
}

func NewTeslaBatteryServer() *TeslaBatteryServer {
	return &TeslaBatteryServer{
		batteryState: &pb.TeslaBatteryState{
			SiteId:             "tesla_site_123",
			BackupReserve:      20,
			BatteryCharging:    true,
			BatteryDischarging: false,
			ConsumptionW:       500,
			GridImportW:        100,
			GridExportW:        50,
			ProductionW:        600,
			StorageLevel:       75.5,
			SystemStatus:       "Operational",
			StormModeEnabled:   false,
			OperationMode:      "self_consumption",
			LastUpdated:        time.Now().Format(time.RFC3339),
		},
	}
}

func (s *TeslaBatteryServer) GetTeslaBatteryStatus(ctx context.Context, req *pb.GetTeslaBatteryStatusRequest) (*pb.TeslaBatteryState, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	s.batteryState.LastUpdated = time.Now().Format(time.RFC3339)
	return s.batteryState, nil
}

func (s *TeslaBatteryServer) UpdateBackupReserve(ctx context.Context, req *pb.UpdateBackupReserveRequest) (*pb.UpdateBackupReserveResponse, error) {
	if req.Percentage < 0 || req.Percentage > 100 {
		return nil, status.Error(codes.InvalidArgument, "Invalid backup reserve percentage")
	}

	s.batteryLock.Lock()
	s.batteryState.BackupReserve = int32(req.Percentage)
	s.batteryState.LastUpdated = time.Now().Format(time.RFC3339)
	s.batteryLock.Unlock()

	return &pb.UpdateBackupReserveResponse{
		Message:    "Backup reserve updated",
		Percentage: req.Percentage,
	}, nil
}

func (s *TeslaBatteryServer) ToggleStormMode(ctx context.Context, req *pb.ToggleStormModeRequest) (*pb.ToggleStormModeResponse, error) {
	s.batteryLock.Lock()
	s.batteryState.StormModeEnabled = req.Enable
	s.batteryState.LastUpdated = time.Now().Format(time.RFC3339)
	s.batteryLock.Unlock()

	return &pb.ToggleStormModeResponse{
		Message: "Storm mode updated",
		Enabled: req.Enable,
	}, nil
}

func (s *TeslaBatteryServer) SetOperationMode(ctx context.Context, req *pb.SetOperationModeRequest) (*pb.SetOperationModeResponse, error) {
	allowedModes := map[string]bool{
		"self_consumption": true,
		"backup":           true,
		"autonomous":       true,
	}

	if !allowedModes[req.Mode] {
		return nil, status.Error(codes.InvalidArgument, "Invalid operation mode")
	}

	s.batteryLock.Lock()
	s.batteryState.OperationMode = req.Mode
	s.batteryState.LastUpdated = time.Now().Format(time.RFC3339)
	s.batteryLock.Unlock()

	return &pb.SetOperationModeResponse{
		Message: "Operation mode updated",
		Mode:    req.Mode,
	}, nil
}

func (s *TeslaBatteryServer) GetPowerUsage(ctx context.Context, req *pb.GetPowerUsageRequest) (*pb.GetPowerUsageResponse, error) {
	s.batteryLock.Lock()
	defer s.batteryLock.Unlock()

	return &pb.GetPowerUsageResponse{
		GridImportW:         s.batteryState.GridImportW,
		GridExportW:         s.batteryState.GridExportW,
		BatteryStorageLevel: s.batteryState.StorageLevel,
		Timestamp:           time.Now().Format(time.RFC3339),
	}, nil
}
