package domain

// Registrar dentista, paciente, fecha, hora y descripcion de los mismos
type Appointment struct {
	ID          int    `json:"id"`
	IdPatient   int    `json:"id_patient"`
	IdDentist   int    `json:"id_dentist"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
}

type AppointmentData struct {
	DNI         string `json:"dni"`
	License     string `json:"license"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
}