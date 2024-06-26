package appointment

import (
	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/store"
)

type Repository interface {
    CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error)  // POST: Agregar turno
    GetAppointmentById(id int) (*domain.Appointment, error) // GET: Traer turno por ID
    UpdateAppointment(appointment domain.Appointment) (*domain.Appointment, error) // PUT: Actualizar turno
    UpdateAppointmentField(id int, field []string, value []string) (*domain.Appointment, error) // PATCH: Actualizar un campo específico del turno
    DeleteAppointment(id int) (error) // DELETE: Eliminar turno
    CreateAppointmentByDNIAndLicense(appointmentData domain.AppointmentData) (*domain.Appointment, error) // POST: Agregar turno por DNI del paciente y matrícula del dentista
    GetAppointmentsByDNI(DNI string) ([]domain.Appointment, error) // GET: Traer turnos por DNI del paciente
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetAppointmentById(id int) (*domain.Appointment, error) {
    appointment, err := r.storage.GetAppointmentById(id)
    if err != nil {
        return nil, err
    }
    return appointment, nil
}

func (r *repository) CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
    createdAppointment, err := r.storage.CreateAppointment(appointment)
    if err != nil {
        return nil, err
    }
    return createdAppointment, nil
}

func (r *repository) DeleteAppointment(id int) error {
    err := r.storage.DeleteAppointment(id)
    if err != nil {
        return err
    }
	return nil
}

func (r *repository) UpdateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
    appointmentUpdated, err := r.storage.UpdateAppointment(appointment)
    if err != nil {
        return nil, err
    }
    return appointmentUpdated, nil
}

func (r *repository) UpdateAppointmentField(id int, field []string, value []string) (*domain.Appointment, error) {
    appointment, err := r.storage.UpdateAppointmentField(id, field, value)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (r *repository) GetAppointmentsByDNI(DNI string) ([]domain.Appointment, error) {
	appointments, err := r.storage.GetAppointmentsByDNI(DNI)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *repository) CreateAppointmentByDNIAndLicense(appointmentData domain.AppointmentData) (*domain.Appointment, error) {
    createdAppointment, err := r.storage.CreateAppointmentByDNIAndLicense(appointmentData)
    if err != nil {
        return nil, err
    }
    return createdAppointment, nil
}