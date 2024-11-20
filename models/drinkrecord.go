package models

type DrinkRecord struct {
	ID        string  `json:"id" db:"id"`
	PatientID string  `json:"patient_id" db:"patient_id"`
	AmountML  float64 `json:"amount_ml" db:"amount_ml"`
	Date      string  `json:"date" db:"date"`
}
