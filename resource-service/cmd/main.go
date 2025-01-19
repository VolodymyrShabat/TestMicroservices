package main

import (
	"fmt"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/config"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/handlers"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/middlewares"
	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/services"
	pb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	path := os.Getenv("CFG_PATH")
	if path == "" {
		path = "./config"
	}
	cfg, err := config.LoadConfig(path)
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	address := fmt.Sprintf(":%d", cfg.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", cfg.Server.Port, err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middlewares.UnaryLoggingInterceptor))

	userService := services.NewUserService()
	bookService := services.NewBookService()

	resourcesHandler := handlers.NewResourceHandler(userService, bookService)
	pb.RegisterResourceServiceServer(grpcServer, resourcesHandler)

	log.Printf("Auth Service listening on port %d...", cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}

}
