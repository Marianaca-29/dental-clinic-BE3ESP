package domain

// AppointmentDNILicense represents information about appointments created by patient´s DNI and dentist´s license
// @Summary AppointmentData information
// @Description AppointmentData information
type AppointmentData struct {
	DNI         string `json:"dni"`
	License     string `json:"license"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
}