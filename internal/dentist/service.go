package dentist

import "DENTAL-CLINIC/internal/domain"

type IService interface {
	CreateDentist(p domain.Dentist) (*domain.Dentist, error)
	GetDentistById(id int) (*domain.Dentist, error)
	DeleteDentist(id int) error
	UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentistField(id int, p domain.Dentist) (*domain.Dentist, error)
}

type Service struct {
	Repository Repository
}

func NewService(repository Repository) IService {
	return &Service{repository}
}

func (s *Service) GetDentistById(id int) (*domain.Dentist, error) {
	p, err := s.Repository.GetDentistById(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) CreateDentist(p domain.Dentist) (*domain.Dentist, error) {
	dentist, err := s.Repository.CreateDentist(p)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

// completar
func (s *Service) DeleteDentist(id int) error {
	return nil
}

func (s *Service) UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error){
	return nil, nil
}
func (s *Service) UpdateDentistField(id int, p domain.Dentist) (*domain.Dentist, error) {
	return nil, nil
}