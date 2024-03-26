package domain

// Patient represents information about a patient.
// @Summary Patient information
// @Description Patient information including ID, last name, first name, address, DNI, and registration date.
type Patient struct {
	ID               int    `json:"id"`
	LastName         string `json:"last_name"`
	FirstName        string `json:"first_name"`
	Address          string `json:"address"`
	DNI              string `json:"dni"`
	RegistrationDate string `json:"registration_date"`
}
