package dentist

import (
	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/store"
	"DENTAL-CLINIC/pkg/web"
)

type Repository interface {
	CreateDentist(d domain.Dentist) (*domain.Dentist, error)
	GetDentistById(id int) (*domain.Dentist, error)
	UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentistField(id int, p domain.Dentist) (*domain.Dentist, error)
	DeleteDentist(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) CreateDentist(d domain.Dentist) (*domain.Dentist, error) {
	exists, err := r.storage.GetDentistIdByLicense(d.License)
	if err != nil {
		return nil, err
	}
	if exists != 0 {
		return nil, web.NewBadRequestApiError("/nEsa matrícula ya existe")
	}
	dent, err := r.storage.CreateDentist(d)
	if err != nil {
		return nil, err
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


// Completar
func (r *repository) DeleteDentist(id int) error {
	return nil
}

func (r *repository) UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error) {
	return nil, nil
}
func (r *repository) UpdateDentistField(id int, p domain.Dentist) (*domain.Dentist, error) {
	return nil, nil
}