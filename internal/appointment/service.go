package appointment

import (
	"DENTAL-CLINIC/internal/domain"
)

type IService interface {
CreateAppointment (appointment domain.Appointment) (*domain.Appointment, error)  // POST: Agregar turno
GetAppointmentById (id int) (*domain.Appointment, error) // GET: Traer turno por ID
UpdateAppointment (appointment domain.Appointment) (*domain.Appointment, error) // PUT: Actualizar turno
UpdateAppointmentField(id int, field string, value string) (*domain.Appointment, error) // PATCH: Actualizar un campo específico del turno
DeleteAppointment (id int) (error) // DELETE: Eliminar turno
CreateAppointmentByDNIAndLicense (DNI string, license string, appointment domain.Appointment) (*domain.Appointment, error) // POST: Agregar turno por DNI del paciente y matrícula del dentista
GetAppointmentsByDNI (DNI string) ([]domain.Appointment, error) // GET: Traer turnos por DNI del paciente

}

type service struct {
	repository Repository
}

func NewService(r Repository) IService {
	return &service{r}
}

func (s *service) GetAppointmentById(id int) (*domain.Appointment, error) {
	appointment, err := s.repository.GetAppointmentById(id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *service) CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
	createdAppointment, err := s.repository.CreateAppointment(appointment)
	if err != nil {
		return nil, err
	}
	return createdAppointment, nil
}

func (s *service) CreateAppointmentByDNIAndLicense(DNI string, license string, appointment domain.Appointment) (*domain.Appointment, error) {
	createdAppointment, err := s.repository.CreateAppointmentByDNIAndLicense(DNI, license, appointment)
	if err != nil {
		return nil, err
	}
	return createdAppointment, nil
}

func (s *service) GetAppointmentsByDNI(DNI string) ([]domain.Appointment, error) {
	appointments, err := s.repository.GetAppointmentsByDNI(DNI)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

// completar
func (s *service) UpdateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
	return nil, nil
}

func (s *service) UpdateAppointmentField(id int, field string, value string) (*domain.Appointment, error) {
	return nil, nil
}

func (s *service) DeleteAppointment(id int) error {
	return nil
}
