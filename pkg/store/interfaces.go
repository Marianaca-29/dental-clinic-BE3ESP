package store

import (
	"DENTAL-CLINIC/internal/domain"
)

type StoreInterface interface {
	// Métodos para dentistas:
	CreateDentist (dentist domain.Dentist) (*domain.Dentist, error)
	GetDentistById (id int) (*domain.Dentist, error)
	UpdateDentist (dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentistField (dentist domain.Dentist) (*domain.Dentist, error)
	DeleteDentist (id int) (error)

	// Métodos para pacientes:
	CreatePatient (patient domain.Patient) (*domain.Patient, error)
	GetPatientById (id int) (*domain.Patient, error)
	UpdatePatient (patient domain.Patient) (*domain.Patient, error)
	UpdatePatientField (patient domain.Patient) (*domain.Patient, error)
	DeletePatient (id int) (error)

	// Métodos para turnos:
	CreateAppointment (appointment domain.Appointment) (*domain.Appointment, error)
	GetAppointmentById (id int) (*domain.Appointment, error)
	UpdateAppointment (appointment domain.Appointment) (*domain.Appointment, error)
	UpdateAppointmentField (appointment domain.Appointment) (*domain.Appointment, error)
	DeleteAppointment (id int) (error)
	CreateAppointmentByDNIAndLicense (DNI int, license string, appointment domain.Appointment) (*domain.Appointment, error)
	GetAppointmentByDNI (DNI int) (*domain.Appointment, error)
}