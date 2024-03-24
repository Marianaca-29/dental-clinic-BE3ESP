package domain

// Registrar dentista, paciente, fecha, hora y descripcion de los mismos
type Appointment struct {
	ID          int    `json:"id"`
	IdDentist   int    `json:"id_dentist" binding:"required"`
	IdPatient   int    `json:"id_patient" binding:"required"`
	Date        string `json:"date" binding:"required"`
	Time        string `json:"time" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type AppointmentData struct {
	DNI         string `json:"dni"`
	License     string `json:"license"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
}