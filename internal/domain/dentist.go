package domain

// Registrar apellido, nombre y matrícula de los mismos
type Dentist struct {
	ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    License   string `json:"license"`
}
