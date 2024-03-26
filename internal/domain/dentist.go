package domain

// Dentist represents information about a dentist.
// @Summary Dentist information
// @Description Dentist information including ID, first name, last name, and license number.
type Dentist struct {
	ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    License   string `json:"license"`
}
