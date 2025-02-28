package main

import (
	"log"
	"net"
	"os"

	"github.com/pbtrad/go_batteries/internal/handlers/batteries"
	givenergy "github.com/pbtrad/go_batteries/proto/batteries/v1/givenergy"
	sonnen "github.com/pbtrad/go_batteries/proto/batteries/v1/sonnen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	sonnen.RegisterSonnenBatteriesServiceServer(s, batteries.NewBatteryServer())
	givenergy.RegisterGivEnergyServiceServer(s, batteries.NewGivEnergyBatteryServer())

	reflection.Register(s)

	log.Printf("gRPC server running on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
