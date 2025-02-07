package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/joho/godotenv"
	"github.com/lakeside763/synapse-crm/api-gateway-service/config"
	"github.com/lakeside763/synapse-crm/api-gateway-service/internals/adapters/repositories"
	"github.com/lakeside763/synapse-crm/api-gateway-service/internals/ports/http/routes"
	log "github.com/sirupsen/logrus"
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
	dataRepo, err := repositories.NewDataRepo(config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	// Main router
	router := httprouter.New()
	routes.UserRoutes(router, dataRepo)


	port := "5400"
	log.Info("Starting server on port: ", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}