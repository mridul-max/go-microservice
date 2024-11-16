package services

import (
	"context"
	"drinks/config"
	"drinks/models"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type PatientService struct {
	client *azcosmos.ContainerClient
}

func NewPatientService() *PatientService {
	// Load the already initialized Cosmos DB client
	client := config.GetCosmosClient()

	return &PatientService{
		client: client,
	}
}

// GetPatientDailyLimit retrieves the daily limit for a patient by Id
func (ps *PatientService) GetPatientDailyGoal(ctx context.Context, Id string) (int, error) {
	// Cosmos DB SQL-like query to find the patient by Id
	query := fmt.Sprintf("SELECT * FROM c WHERE c.Id = '%s'", Id)
	partitionKey := azcosmos.NewPartitionKeyString(Id)
	pager := ps.client.NewQueryItemsPager(query, partitionKey, nil)

	var patient models.PatientDTO
	found := false

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return 0, fmt.Errorf("failed to query patients: %v", err)
		}

		for _, item := range page.Items {
			if err := json.Unmarshal(item, &patient); err != nil {
				return 0, fmt.Errorf("failed to parse patient data: %v", err)
			}
			found = true
		}
	}

	if !found {
		return 0, fmt.Errorf("patient with Id %s not found", Id)
	}

	return patient.DailyGoal, nil // Assuming DailyLimit is part of your PatientDTO
}
