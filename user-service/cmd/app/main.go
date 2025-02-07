package main

import (
	"net"

	"github.com/joho/godotenv"
	"github.com/lakeside763/synapse-crm/user-service/config"
	"github.com/lakeside763/synapse-crm/user-service/internals/adapters/repositories"
	"github.com/lakeside763/synapse-crm/user-service/internals/cores/services"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	users "github.com/lakeside763/synapse-crm-pkg/protos/users"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Info("Loading Configuration")
	}
}

func main() {
	// Initialize the configuration
	config := config.NewConfig()

	// Initialize the database
	dataRepo, err := repositories.NewDataRepository(config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	// User server
	server := services.NewUserServiceServer(dataRepo.User)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	users.RegisterUserServiceServer(grpcServer, server)

	log.Info("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}