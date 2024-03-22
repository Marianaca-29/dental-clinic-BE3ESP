package domain

// Registrar apellido, nombre, domicilio, DNi y fecha de alta de los mismos
type Patient struct {
	ID               int    `json:"id"`
	LastName         string `json:"last_name" binding:"required"`
	FirstName        string `json:"first_name" binding:"required"`
	Address          string `json:"license" binding:"required"`
	DNI              string `json:"dni" binding:"required"`
	RegistrationDate string `json:"registration_date" binding:"required"`
}
