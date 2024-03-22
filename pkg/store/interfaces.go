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
	UpdatePatient (dentist domain.Patient) (*domain.Patient, error)
	UpdatePatientField (dentist domain.Patient) (*domain.Patient, error)
	DeletePatient (id int) (error)

	// Métodos para turnos:
}