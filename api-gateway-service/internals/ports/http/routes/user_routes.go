package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lakeside763/synapse-crm/api-gateway-service/internals/adapters/repositories"
	"github.com/lakeside763/synapse-crm/api-gateway-service/internals/cores/domains/users/services"
	"github.com/lakeside763/synapse-crm/api-gateway-service/internals/ports/http/handlers"
	log "github.com/sirupsen/logrus"
)


func UserRoutes(router *httprouter.Router, dataRepo *repositories.DataRepo) {

	// Initialize the user service
	userService, err := services.NewUserServiceClient("localhost:50051") // gRPC server address
	if err != nil {
		log.Fatalf("Failed to initialize user service client: %v", err)
	}

	userHandler := handlers.NewUserHandler(userService)

	// router.GET("/users", GetUsers)
	// router.GET("/users/:id", GetUser)
	router.POST("/users", userHandler.CreateUser)
}