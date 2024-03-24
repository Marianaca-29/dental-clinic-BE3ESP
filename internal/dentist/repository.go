package dentist

import (
	"fmt"

	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/store"
	"DENTAL-CLINIC/pkg/web"
)

type Repository interface {
	CreateDentist(d domain.Dentist) (*domain.Dentist, error)
	GetDentistById(id int) (*domain.Dentist, error)
	UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentistField(id int, values []string, fields[]string) (*domain.Dentist, error)
	DeleteDentist(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) CreateDentist(d domain.Dentist) (*domain.Dentist, error) {
	exists, _ := r.storage.GetDentistIdByLicense(d.License)
	if exists != 0 {
		return nil, web.NewBadRequestApiError("Esa matrícula ya existe")
	}

	dent, err := r.storage.CreateDentist(d)
	if (err != nil) {
		return nil, web.NewBadRequestApiError("No se ha podido crear el dentista")
	}
	
	return dent, nil
}

func (r *repository) GetDentistById(id int) (*domain.Dentist, error) {
	dentistResponse, err := r.storage.GetDentistById(id)
	if err != nil {
		return nil, web.NewNotFoundApiError("/nDentista inexistente")
	}

	return dentistResponse, nil
}

func (r *repository) UpdateDentist(d domain.Dentist) (*domain.Dentist, error) {
	dentist, err := r.storage.UpdateDentist(d)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

func (r *repository) UpdateDentistField(id int, field []string, values []string) (*domain.Dentist, error) {
	dentist, err := r.storage.UpdateDentistField(id, field, values)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

func (r *repository) DeleteDentist(id int) error {
	_, err := r.storage.GetDentistById(id)
	if err != nil {
		return web.NewBadRequestApiError(fmt.Sprintf("/nNo existe un dentista con el id %d", id))
	}
	errDelete := r.storage.DeleteDentist(id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}