package routes

import (
	"drinks/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(patientController *controllers.PatientController, getPatientPhoneNumberByIdController *controllers.GetPatientPhoneNumberByIdController) *mux.Router {
	router := mux.NewRouter()

	// Get daily goal for a patient by Id
	router.HandleFunc("/patients/dailygoal", patientController.GetPatientDailyGoal).Methods(http.MethodGet)

	// Drink record routes
	router.HandleFunc("/patient/phoneNumber", getPatientPhoneNumberByIdController.GetPatientPhoneNumberById).Methods(http.MethodGet)

	return router
}
