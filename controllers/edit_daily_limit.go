package controllers

import (
	"context"
	"drinks/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type PatientController struct {
	service *services.PatientService
}

func NewPatientController(service *services.PatientService) *PatientController {
	return &PatientController{
		service: service,
	}
}

// @Summary Get patient's daily goal
// @Description Get the daily goal of a patient by their ID
// @Tags patients
// @Param Id query string true "Patient ID"
// @Success 200 {object} map[string]int
// @Failure 400 {string} string "Missing or invalid query parameter"
// @Failure 500 {string} string "Internal server error"
// @Router /patients/dailygoal [get]
func (pc *PatientController) GetPatientDailyGoal(w http.ResponseWriter, r *http.Request) {
	// Get the Id from query parameters
	Id := r.URL.Query().Get("Id")
	if Id == "" {
		http.Error(w, "Missing Id query parameter", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	dailygoal, err := pc.service.GetPatientDailyGoal(ctx, Id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving patient: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the daily limit
	response := map[string]int{"Dailygoal": dailygoal}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
