package patient

import (
	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/store"
	"DENTAL-CLINIC/pkg/web"
	"fmt"
)

type Repository interface {
	CreatePatient(patient domain.Patient) (*domain.Patient, error)
	GetPatientById(id int) (*domain.Patient, error)
	UpdatePatient(patient domain.Patient) (*domain.Patient, error)
	UpdatePatientField(id int, field []string, value []string) (*domain.Patient, error)
	DeletePatient(id int) (error)
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) CreatePatient(patient domain.Patient) (*domain.Patient, error) {
	exists, _ := r.storage.GetPatientIdByDNI(patient.DNI)
	if exists != 0 {
		return nil, web.NewBadRequestApiError("Ese DNI ya existe")
	}
	
	patientResponse, err := r.storage.CreatePatient(patient)
	if (err != nil) {
		return nil, web.NewBadRequestApiError("No se ha podido crear el paciente")
	}

	return patientResponse, nil
}

func (r *repository) GetPatientById(id int) (*domain.Patient, error) {
	patientResponse, err := r.storage.GetPatientById(id)
	if (err != nil) {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("No se ha encontrado ningun paciente con el id %d", id))
	}

	return patientResponse, nil
}

func (r *repository) UpdatePatient(patient domain.Patient) (*domain.Patient, error) {
	updatedPatient, err := r.storage.UpdatePatient(patient)
	if err != nil {
		return nil, err
	}
	return updatedPatient, nil
}

func (r *repository) UpdatePatientField(id int, field []string, value []string) (*domain.Patient, error){
	updatedPatient, err := r.storage.UpdatePatientField(id, field, value)
    if err != nil {
        return nil, err
    }
    return updatedPatient, nil
}

func (r *repository) DeletePatient(id int) (error) {
	err := r.storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}
