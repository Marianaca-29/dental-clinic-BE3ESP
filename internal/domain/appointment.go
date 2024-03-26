package domain

// Appointment represents information about appointments 
// @Summary Appointment information
// @Description Appointment information
type Appointment struct {
	ID          int    `json:"id"`
	IdPatient   int    `json:"id_patient"`
	IdDentist   int    `json:"id_dentist"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
}