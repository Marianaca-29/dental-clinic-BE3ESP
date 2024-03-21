package domain
// Registrar apellido, nombre y matrícula de los mismos
type Dentist struct {
	ID        int    `json:"id"`
    LastName  string `json:"last_name" binding:"required"`
    FirstName string `json:"first_name" binding:"required"`
    License   string `json:"license" binding:"required"`
}
