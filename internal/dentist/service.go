package dentist

import (
	"DENTAL-CLINIC/internal/domain"
)

type IService interface {
	CreateDentist(d domain.Dentist) (*domain.Dentist, error)
	GetDentistById(id int) (*domain.Dentist, error)
	DeleteDentist(id int) error
	UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error)
	UpdateDentistField(d domain.Dentist) (*domain.Dentist, error)
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

func (s *service) CreateDentist(d domain.Dentist) (*domain.Dentist, error) {
	dentist, err := s.Repository.CreateDentist(d)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

func (s *service) UpdateDentist(d domain.Dentist) (*domain.Dentist, error){
	dentistFound, err := s.Repository.GetDentistById(d.ID)
	if err != nil {
		return nil, err
	}

	if d.FirstName == "" {
		d.FirstName = dentistFound.FirstName
	}
	if d.LastName == "" {
		d.LastName = dentistFound.LastName
	}
	if d.License == "" {
		d.License = dentistFound.License
	}
	
	dentist, err := s.Repository.UpdateDentist(d)
	if err != nil {
		return nil, err
	}
	
	return dentist, nil
}

func (s *service) UpdateDentistField(d domain.Dentist) (*domain.Dentist, error) {
	var field []string
	var values []string
	if d.FirstName != "" {
		field = append(field, "first_name")
		values = append(values, d.FirstName)
	}
	if d.LastName != "" {
		field = append(field, "last_name")
		values = append(values, d.LastName)
	}
	if d.License != "" {
		field = append(field, "license")
		values = append(values, d.License)
	}
	dentist, err := s.Repository.UpdateDentistField(d.ID, field, values)
	if err != nil {
		return nil, err
	}
	return dentist, nil
}

func (s *service) DeleteDentist(id int) error {
	err := s.Repository.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}