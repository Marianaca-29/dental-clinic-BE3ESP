package dentist

import "DENTAL-CLINIC/internal/domain"

type IService interface {
	GetByID(id int) (domain.Dentist, error)
	Create(p domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
	Update(id int, p domain.Dentist) (domain.Dentist, error)
}

type Service struct {
	Repository IRepository
}

func NewService(repository IRepository) IService {
	return &Service{Repository: repository}
}

func (s *Service) GetByID(id int) (domain.Dentist, error) {
	p, err := s.Repository.GetById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

func (s *Service) Create(p domain.Dentist) (domain.Dentist, error) {
	p, err := s.Repository.Create(p)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

// completar
func (s *Service) Delete(id int) error {
	return nil
}

func (s *Service) Update(id int, p domain.Dentist) (domain.Dentist, error) {
	return domain.Dentist{}, nil
}