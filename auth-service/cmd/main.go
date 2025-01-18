package main

import (
	"log"
	"net"

	"fmt"

	"github.com/VolodymyrShabat/testTask27NerdsAuthService/internal/config"
	"github.com/VolodymyrShabat/testTask27NerdsAuthService/internal/handlers"
	pb "github.com/VolodymyrShabat/testTask27NerdsAuthService/internal/proto" // Import the generated proto package
	"google.golang.org/grpc"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
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
	pb.RegisterAuthServiceServer(grpcServer, &handlers.AuthHandler{})

	log.Printf("Auth Service listening on port %d...", cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}

}
