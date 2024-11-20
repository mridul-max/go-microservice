package controllers

import (
	"context"
	"drinks/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type DrinkRecordController struct {
	service *services.DrinkRecordService
}

func NewDrinkRecordController(service *services.DrinkRecordService) *DrinkRecordController {
	return &DrinkRecordController{
		service: service,
	}
}

// @Summary Get drink record by ID
// @Description Get a drink record by its ID
// @Tags drinkrecords
// @Param id query string true "Drink Record ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string "Missing or invalid query parameter"
// @Failure 404 {string} string "Drink record not found"
// @Failure 500 {string} string "Internal server error"
// @Router /api/drinkrecords/byid [get]
func (drc *DrinkRecordController) GetDrinkRecordById(w http.ResponseWriter, r *http.Request) {
	// Retrieve the "id" query parameter
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}

	// Call the service to get the drink record
	ctx := context.TODO()
	record, err := drc.service.GetDrinkRecordById(ctx, id)
	if err != nil {
		if err.Error() == fmt.Sprintf("drink record with ID %s not found", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error retrieving drink record: %v", err), http.StatusInternalServerError)
		}
		return
	}

	// Respond with the drink record
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(record)
}
