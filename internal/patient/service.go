package patient

import (
	"DENTAL-CLINIC/internal/domain"
)

type IService interface {
	CreatePatient (patient domain.Patient) (*domain.Patient, error)
	GetPatientById (id int) (*domain.Patient, error)
	UpdatePatient (patient domain.Patient) (*domain.Patient, error)
	UpdatePatientField (id int, field string, value string) (*domain.Patient, error)
	DeletePatient (id int) (error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) IService {
	return &service{r}
}

func (s *service) CreatePatient (patient domain.Patient) (*domain.Patient, error) {
	patientResponse, err := s.repository.CreatePatient(patient)
	if (err != nil) {
		return nil, err
	}

	return patientResponse, nil
}

func (s *service) GetPatientById (id int) (*domain.Patient, error) {
	patientResponse, err := s.repository.GetPatientById(id)
	if (err != nil) {
		return nil, err
	}

	return patientResponse, nil
}

func (s *service) UpdatePatient (patient domain.Patient) (*domain.Patient, error) {
	return nil, nil
}

func (s *service) UpdatePatientField (id int, field string, value string) (*domain.Patient, error) {
	return nil, nil
}

func (s *service) DeletePatient (id int) (error) {
	return nil
}

