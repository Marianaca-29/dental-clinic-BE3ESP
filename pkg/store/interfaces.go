package store

import (
	"DENTAL-CLINIC/internal/domain"
)

type StoreInterface interface {
	// Métodos CRUD para dentistas:
	CreateDentist (dentist domain.Dentist) (*domain.Dentist, error) // POST: Agregar dentista
	GetDentistById (id int) (*domain.Dentist, error) // GET: Traer dentista por ID
	UpdateDentist (dentist domain.Dentist) (*domain.Dentist, error) // PUT: Actualizar dentista
	UpdateDentistField (id int, field []string, value []string) (*domain.Dentist, error) // PATCH: Actualizar un campo específico del dentista
	DeleteDentist (id int) (error) // DELETE: Eliminar dentista
	GetDentistIdByLicense (license string) (int, error) // GET: Traer dentista por Matricula

	// Métodos CRUD para pacientes:
	CreatePatient (patient domain.Patient) (*domain.Patient, error) // POST: Agregar paciente
	GetPatientById (id int) (*domain.Patient, error) // GET: Traer paciente por ID
	UpdatePatient (patient domain.Patient) (*domain.Patient, error) // PUT: Actualizar paciente
	UpdatePatientField (id int, field []string, value []string) (*domain.Patient, error) // PATCH: Actualizar un campo específico del paciente
	DeletePatient (id int) (error) // DELETE: Eliminar paciente
	GetPatientIdByDNI (DNI string) (int, error) // GET: Traer paciente por DNI

	// Métodos CRUD para turnos:
	CreateAppointment (appointment domain.Appointment) (*domain.Appointment, error)  // POST: Agregar turno
	GetAppointmentById (id int) (*domain.Appointment, error) // GET: Traer turno por ID
	UpdateAppointment (appointment domain.Appointment) (*domain.Appointment, error) // PUT: Actualizar turno
	UpdateAppointmentField (id int, field []string, value []string) (*domain.Appointment, error) // PATCH: Actualizar un campo específico del turno
	DeleteAppointment (id int) (error) // DELETE: Eliminar turno
	CreateAppointmentByDNIAndLicense (DNI string, license string, appointment domain.Appointment) (*domain.Appointment, error) // POST: Agregar turno por DNI del paciente y matrícula del dentista
	GetAppointmentsByDNI (DNI string) ([]domain.Appointment, error) // GET: Traer turnos por DNI del paciente
}