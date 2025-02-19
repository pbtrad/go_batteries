package main

import (
	"log"
	"net"
	"os"

	"github.com/pbtrad/go_batteries/internal/handlers/batteries"
	pb "github.com/pbtrad/go_batteries/proto/batteries/v1"
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
	pb.RegisterBatteriesServiceServer(s, batteries.NewBatteryServer())
	reflection.Register(s)

	log.Printf("gRPC server running on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
