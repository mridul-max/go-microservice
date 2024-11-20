package controllers

import (
	"context"
	"drinks/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetPatientPhoneNumberByIdController struct {
	service *services.PatientService
}

func NewGetPatientPhoneNumberByIdController(service *services.PatientService) *GetPatientPhoneNumberByIdController {
	return &GetPatientPhoneNumberByIdController{
		service: service,
	}
}

// @Summary Get patient's phoneNumber
// @Description Get the phoneNumber of a patient by their ID
// @Tags patients
// @Param Id query string true "Patient ID"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Missing or invalid query parameter"
// @Failure 500 {string} string "Internal server error"
// @Router /patient/phoneNumber [get]
func (pc *GetPatientPhoneNumberByIdController) GetPatientPhoneNumberById(w http.ResponseWriter, r *http.Request) {
	// Get the Id from query parameters
	Id := r.URL.Query().Get("Id")
	if Id == "" {
		http.Error(w, "Missing Id query parameter", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	patientphonenumber, err := pc.service.GetPatientPhoneNumberById(ctx, Id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving patient: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the daily limit
	response := map[string]string{"PhoneNumber": patientphonenumber}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
