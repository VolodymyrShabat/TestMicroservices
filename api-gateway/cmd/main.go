package main

import (
	"fmt"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/config"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/handlers"
	"github.com/VolodymyrShabat/TestMicroservices/api-gateway/internal/routes"
	authpb "github.com/VolodymyrShabat/TestMicroservices/auth-service/pkg/proto"
	resourcepb "github.com/VolodymyrShabat/TestMicroservices/resource-service/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

func main() {

	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	connAuth, err := grpc.Dial(fmt.Sprintf("localhost:%d", cfg.Server.AuthPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Auth service: %v", err)
	}
	defer connAuth.Close()

	// 2. Create the gRPC client
	authClient := authpb.NewAuthServiceClient(connAuth)

	connResources, err := grpc.Dial(fmt.Sprintf("localhost:%d", cfg.Server.ResourcesPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Auth service: %v", err)
	}
	defer connResources.Close()

	resourceClient := resourcepb.NewResourceServiceClient(connResources)

	authHandlers := &handlers.AuthHandlers{
		AuthClient: authClient,
	}
	resourceHandlers := &handlers.ResourceHandlers{
		ResourceClient: resourceClient,
	}

	r := routes.SetupRouter(*authHandlers, *resourceHandlers)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port), // API Gateway listening on port 8080
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
	}

	log.Println(fmt.Sprintf("API Gateway listening on port - %d ...", cfg.Server.Port))
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
