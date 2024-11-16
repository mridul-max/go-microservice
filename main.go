package main

import (
	"drinks/config"
	"drinks/controllers"
	"drinks/routes"
	"drinks/services"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize Cosmos DB connection
	config.InitCosmosDB()

	// Initialize the patient service and controller
	patientService := services.NewPatientService()
	patientController := controllers.NewPatientController(patientService)

	// Setup the application routes
	router := routes.SetupRoutes(patientController)

	// Start the HTTP server on port 8080
	log.Println("Server is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
