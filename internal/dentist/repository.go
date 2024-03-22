package dentist

import (
	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/store"
	"DENTAL-CLINIC/pkg/web"
)

type IRepository interface {
	GetById(id int) (domain.Dentist, error)
	Create(d domain.Dentist) (domain.Dentist, error)
}

type Repository struct {
	Store store.StoreInterface
}

func NewRepository(Store store.StoreInterface) IRepository {
	return &Repository{Store}
}

func (r *Repository) GetById(id int) (domain.Dentist, error) {
	dent, err := r.Store.GetDentistById(id)
	if err != nil {
		return domain.Dentist{}, web.NewNotFoundApiError("/nDentista inexistente")
	}

	return *dent, nil
}

func (r *Repository) Create(d domain.Dentist) (domain.Dentist, error) {
	exists, err := r.Store.GetDentistIdByLicense(d.License)
	if err != nil {
		return domain.Dentist{}, err
	}
	if exists != 0 {
		return domain.Dentist{}, web.NewBadRequestApiError("/nEsa matrícula ya existe")
	}
	dent, err := r.Store.CreateDentist(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return *dent, nil
}

// Completar
func (r *Repository) Delete(id int) error {

	return nil
}

func (r *Repository) Update(id int, p domain.Dentist) (domain.Dentist, error) {
	return p, nil
}