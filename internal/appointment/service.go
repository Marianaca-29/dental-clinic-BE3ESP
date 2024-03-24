package appointment

import (
	"DENTAL-CLINIC/internal/domain"
)

type IService interface {
	CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error)  // POST: Agregar turno
	GetAppointmentById(id int) (*domain.Appointment, error) // GET: Traer turno por ID
	UpdateAppointment(appointment domain.Appointment) (*domain.Appointment, error) // PUT: Actualizar turno
	UpdateAppointmentField(appointment domain.Appointment) (*domain.Appointment, error) // PATCH: Actualizar un campo específico del turno
	DeleteAppointment(id int) (error) // DELETE: Eliminar turno
	CreateAppointmentByDNIAndLicense(appointmentData domain.AppointmentData) (*domain.Appointment, error) // POST: Agregar turno por DNI del paciente y matrícula del dentista
	GetAppointmentsByDNI(DNI string) ([]domain.Appointment, error) // GET: Traer turnos por DNI del paciente
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

func (s *service) UpdateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
	appointmentFound, err := s.repository.GetAppointmentById(appointment.ID)
	if err != nil {
		return nil, err
	}

	if appointment.IdDentist == 0 {
		appointment.IdDentist = appointmentFound.IdDentist
	}
	if appointment.IdPatient == 0 {
		appointment.IdPatient = appointmentFound.IdPatient
	}
	if appointment.Date == "" {
		appointment.Date = appointmentFound.Date
	}
	if appointment.Time == "" {
		appointment.Time = appointmentFound.Time
	}
	if appointment.Description == "" {
		appointment.Description = appointmentFound.Description
	}
	
	appointmentUpdated, err := s.repository.UpdateAppointment(appointment)
	if err != nil {
		return nil, err
	}
	
	return appointmentUpdated, nil
}

func (s *service) UpdateAppointmentField(appointment domain.Appointment) (*domain.Appointment, error) {
	var field []string
	var values []string
	if appointment.Date != "" {
		field = append(field, "date")
		values = append(values, appointment.Date)
	}
	if appointment.Time != "" {
		field = append(field, "time")
		values = append(values, appointment.Time)
	}
	if  appointment.Description != "" {
		field = append(field, "description")
		values = append(values, appointment.Description)
	}

	appointmentUpdated, err := s.repository.UpdateAppointmentField(appointment.ID, field, values) 
	if err != nil {
		return nil, err
	}
	return appointmentUpdated, nil
}

func (s *service) DeleteAppointment(id int) error {
	err := s.repository.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateAppointmentByDNIAndLicense(appointmentData domain.AppointmentData) (*domain.Appointment, error) {
	createdAppointment, err := s.repository.CreateAppointmentByDNIAndLicense(appointmentData)
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