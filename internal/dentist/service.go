package dentist

import "DENTAL-CLINIC/internal/domain"

type IService interface {
	CreateDentist(p domain.Dentist) (*domain.Dentist, error)
	GetDentistById(id int) (*domain.Dentist, error)
	DeleteDentist(id int) error
	UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentistField(id int, p domain.Dentist) (*domain.Dentist, error)
}

type service struct {
	Repository Repository
}

func NewService(repository Repository) IService {
	return &service{repository}
}

func (s *service) GetDentistById(id int) (*domain.Dentist, error) {
	p, err := s.Repository.GetDentistById(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *service) CreateDentist(p domain.Dentist) (*domain.Dentist, error) {
	dentist, err := s.Repository.CreateDentist(p)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

// completar
func (s *service) UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error){
	return nil, nil
}
func (s *service) UpdateDentistField(id int, p domain.Dentist) (*domain.Dentist, error) {
	return nil, nil
}

func (s *service) DeleteDentist(id int) error {
	err := s.Repository.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}