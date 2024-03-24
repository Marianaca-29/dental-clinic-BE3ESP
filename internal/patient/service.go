package patient

import (
	"DENTAL-CLINIC/internal/domain"
)

type IService interface {
	CreatePatient (patient domain.Patient) (*domain.Patient, error)
	GetPatientById (id int) (*domain.Patient, error)
	UpdatePatient (patient domain.Patient) (*domain.Patient, error)
	UpdatePatientField (p domain.Patient) (*domain.Patient, error)
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
	patientFound, err := s.repository.GetPatientById(patient.ID)
	if err != nil {
		return nil, err
	}

	if patient.FirstName == "" {
		patient.FirstName = patientFound.FirstName
	}
	if patient.LastName == "" {
		patient.LastName = patientFound.LastName
	}
	if patient.Address == "" {
		patient.Address = patientFound.Address
	}
	if patient.DNI == "" {
		patient.DNI = patientFound.DNI
	}
	if patient.RegistrationDate == "" {
		patient.RegistrationDate = patientFound.RegistrationDate
	}
	
	patientUpdated, err := s.repository.UpdatePatient(patient)
	if err != nil {
		return nil, err
	}
	
	return patientUpdated, nil
}

func (s *service) UpdatePatientField (p domain.Patient) (*domain.Patient, error) {
	var field []string
	var values []string
	if p.FirstName != "" {
		field = append(field, "first_name")
		values = append(values, p.FirstName)
	}
	if p.LastName != "" {
		field = append(field, "last_name")
		values = append(values, p.LastName)
	}
	if p.Address != "" {
		field = append(field, "address")
		values = append(values, p.Address)
	}
	if p.DNI != "" {
		field = append(field, "dni")
		values = append(values, p.DNI)
	}
	if p.RegistrationDate != "" {
		field = append(field, "registration_date")
		values = append(values, p.RegistrationDate)
	}

	patient, err := s.repository.UpdatePatientField(p.ID, field, values)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (s *service) DeletePatient (id int) (error) {
	err := s.repository.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}

