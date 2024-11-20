package main

import (
	"drinks/config"
	"drinks/controllers"
	"drinks/routes"
	"drinks/services"
	"log"
	"net/http"

	_ "drinks/docs" // Import the generated Swagger docs

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Load environment variables from .env file
	//for local test
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// Initialize Cosmos DB connection
	config.InitCosmosDB()

	// Initialize the patient service and controller
	patientService := services.NewPatientService()

	patientController := controllers.NewPatientController(patientService)
	getPatientPhoneNumberByIdController := controllers.NewGetPatientPhoneNumberByIdController(patientService)

	// Set up routes with controllers
	router := routes.SetupRoutes(patientController, getPatientPhoneNumberByIdController)

	// Serve Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the HTTP server on port 8082
	log.Println("Server is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
