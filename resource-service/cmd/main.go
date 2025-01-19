package main

import (
	"fmt"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/config"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/handlers"
	pb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ResourceServer struct {
	pb.UnimplementedResourceServiceServer
}

func main() {

	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	address := fmt.Sprintf(":%d", cfg.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", cfg.Server.Port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterResourceServiceServer(grpcServer, &handlers.ResourceHandler{})

	log.Printf("Auth Service listening on port %d...", cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}

}
