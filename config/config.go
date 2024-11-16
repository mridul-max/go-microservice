package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

var cosmosClient *azcosmos.ContainerClient

// Config holds the application's database configuration.
type Config struct {
	DBURI  string
	DBKey  string
	DBName string
}

// LoadConfig loads configuration from environment variables and checks for missing values.
func LoadConfig() (*Config, error) {
	dbURI := os.Getenv("DBURI")
	dbKey := os.Getenv("DBKEY")
	dbName := os.Getenv("DBNAME")

	// Validate that the required environment variables are set
	if dbURI == "" {
		return nil, fmt.Errorf("environment variable DBURI is not set")
	}
	if dbKey == "" {
		return nil, fmt.Errorf("environment variable DBKEY is not set")
	}
	if dbName == "" {
		return nil, fmt.Errorf("environment variable DBNAME is not set")
	}

	// Return the populated configuration struct
	return &Config{
		DBURI:  dbURI,
		DBKey:  dbKey,
		DBName: dbName,
	}, nil
}

// InitCosmosDB initializes the Cosmos DB client using the provided configuration.
func InitCosmosDB() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up Cosmos DB client options
	credential, err := azcosmos.NewKeyCredential(cfg.DBKey)
	if err != nil {
		log.Fatalf("Failed to create key credential: %v", err)
	}

	client, err := azcosmos.NewClientWithKey(cfg.DBURI, credential, nil)
	if err != nil {
		log.Fatalf("Failed to create Cosmos client: %v", err)
	}

	// Define the database and container (patient collection)
	databaseClient, err := client.NewDatabase(cfg.DBName)
	if err != nil {
		log.Fatalf("Failed to create database client: %v", err)
	}

	containerClient, err := databaseClient.NewContainer("Patients")
	if err != nil {
		log.Fatalf("Failed to create container client: %v", err)
	}

	cosmosClient = containerClient
	log.Println("Successfully connected to Azure Cosmos DB")
}

// GetCosmosClient returns the initialized Cosmos DB client.
func GetCosmosClient() *azcosmos.ContainerClient {
	if cosmosClient == nil {
		log.Fatal("Cosmos DB client is not initialized")
	}
	return cosmosClient
}
