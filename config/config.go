package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

var (
	cosmosClientMap map[string]*azcosmos.ContainerClient // Map to hold container clients
)

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

	return &Config{
		DBURI:  dbURI,
		DBKey:  dbKey,
		DBName: dbName,
	}, nil
}

// InitCosmosDB initializes the Cosmos DB client and containers dynamically.
func InitCosmosDB() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create the Cosmos DB credential
	credential, err := azcosmos.NewKeyCredential(cfg.DBKey)
	if err != nil {
		log.Fatalf("Failed to create key credential: %v", err)
	}

	// Create the Cosmos DB client
	client, err := azcosmos.NewClientWithKey(cfg.DBURI, credential, nil)
	if err != nil {
		log.Fatalf("Failed to create Cosmos client: %v", err)
	}

	// Initialize the map for container clients
	cosmosClientMap = make(map[string]*azcosmos.ContainerClient)

	// Define the database client
	databaseClient, err := client.NewDatabase(cfg.DBName)
	if err != nil {
		log.Fatalf("Failed to create database client: %v", err)
	}

	// Initialize containers: "Patients" and "DrinkRecords"
	containers := []string{"Patients"}
	for _, containerName := range containers {
		containerClient, err := databaseClient.NewContainer(containerName)
		if err != nil {
			log.Fatalf("Failed to create container client for %s: %v", containerName, err)
		}
		cosmosClientMap[containerName] = containerClient
	}

	log.Println("Successfully connected to Azure Cosmos DB and initialized containers")
}

// GetContainerClient returns the Cosmos DB container client for a specified container name.
func GetContainerClient(containerName string) *azcosmos.ContainerClient {
	if cosmosClientMap == nil || cosmosClientMap[containerName] == nil {
		log.Fatalf("Cosmos DB client for container %s is not initialized", containerName)
	}
	return cosmosClientMap[containerName]
}
