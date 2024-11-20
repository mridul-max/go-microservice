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
	client := config.GetContainerClient("Patients")

	return &PatientService{
		client: client,
	}
}

// GetPatientDailygoal retrieves the daily goal for a patient by Id
func (ps *PatientService) GetPatientDailyGoal(ctx context.Context, Id string) (float64, error) {
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

	return patient.DailyGoal, nil // Assuming Dailygoal is part of your PatientDTO
}

// GetDrinkRecordById retrieves a phone number by its ID
func (ps *PatientService) GetPatientPhoneNumberById(ctx context.Context, Id string) (string, error) {
	// Cosmos DB SQL-like query to find the patient by Id
	query := fmt.Sprintf("SELECT * FROM c WHERE c.Id = '%s'", Id)
	partitionKey := azcosmos.NewPartitionKeyString(Id)
	pager := ps.client.NewQueryItemsPager(query, partitionKey, nil)

	var patient models.PatientDTO
	found := false

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to query patients: %v", err)
		}

		for _, item := range page.Items {
			if err := json.Unmarshal(item, &patient); err != nil {
				return "", fmt.Errorf("failed to parse patient data: %v", err)
			}
			found = true
		}
	}

	if !found {
		return "", fmt.Errorf("patient with Id %s not found", Id)
	}

	return patient.PhoneNumber, nil
}
