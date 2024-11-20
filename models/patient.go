package models

// PatientResponseDTO is the data transfer object for a patient response.
type PatientDTO struct {
	Id          string `json:"Id"`
	DailyGoal   int    `json:"DailyGoal"`
	PhoneNumber string `json:"PhoneNumber"`
}
