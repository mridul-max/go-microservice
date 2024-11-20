package services

import (
	"context"
	"drinks/config"
	"drinks/models"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type DrinkRecordService struct {
	client *azcosmos.ContainerClient
}

func NewDrinkRecordService() *DrinkRecordService {
	// Load the DrinkRecords container
	client := config.GetContainerClient("DrinkRecords")
	return &DrinkRecordService{
		client: client,
	}
}

// GetDrinkRecordById retrieves a drink record by its ID
func (drs *DrinkRecordService) GetDrinkRecordById(ctx context.Context, id string) (*models.DrinkRecord, error) {
	// Cosmos DB SQL-like query to find the drink record by ID
	query := fmt.Sprintf("SELECT * FROM c WHERE c.id = '%s'", id)
	partitionKey := azcosmos.NewPartitionKeyString(id)
	pager := drs.client.NewQueryItemsPager(query, partitionKey, nil)

	var drinkRecord models.DrinkRecord
	found := false

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query drink records: %v", err)
		}

		for _, item := range page.Items {
			if err := json.Unmarshal(item, &drinkRecord); err != nil {
				return nil, fmt.Errorf("failed to parse drink record data: %v", err)
			}
			found = true
			break
		}

		if found {
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("drink record with ID %s not found", id)
	}

	return &drinkRecord, nil
}
