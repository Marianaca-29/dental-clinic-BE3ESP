package store

import (
	"fmt"
	"strings"

	"DENTAL-CLINIC/internal/domain"
	"database/sql"
	"errors"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

// Métodos para dentistas:

func (s *sqlStore) CreateDentist(dentist domain.Dentist) (*domain.Dentist, error) {
	query := "INSERT INTO dentists (first_name, last_name, license) VALUES (?, ?, ?)"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(dentist.FirstName, dentist.LastName, dentist.License)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lastId, _ := res.LastInsertId()
	dentist.ID = int(lastId)

	return &dentist, nil
}

func (s *sqlStore) GetDentistById(id int) (*domain.Dentist, error) {
	query := "SELECT * FROM dentists WHERE id_dentist = ?"
	row := s.db.QueryRow(query, id)

	var dentist domain.Dentist
	err := row.Scan(&dentist.ID, &dentist.FirstName, &dentist.LastName, &dentist.License)
	if err != nil {
		return nil, err
	}

	return &dentist, nil
}

func (s *sqlStore) UpdateDentist(dentist domain.Dentist) (*domain.Dentist, error) {
	query := "UPDATE dentists SET firts_name = ?, last_name = ?, license = ? WHERE id_dentist = ?"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(dentist.FirstName, dentist.LastName, dentist.License, dentist.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, err
	}

	query = "SELECT * FROM dentists WHERE id_dentist = ?"
	row := s.db.QueryRow(query, dentist.ID)

	var updatedDentist domain.Dentist
	err = row.Scan(&updatedDentist.ID, &updatedDentist.FirstName, &updatedDentist.LastName, &updatedDentist.License)
	if err != nil {
		return nil, err
	}

	return &updatedDentist, nil
}

func (s *sqlStore) UpdateDentistField(id int, field []string, value []string) (*domain.Dentist, error) {

	// Verificar si los slices de campos y valores tienen la misma longitud
    if len(field) != len(value) {
        return nil, errors.New("la cantidad de campos y valores no coincide")
    }

	// Crear la parte de la consulta SQL para actualizar campos
    var setClause strings.Builder
    for i, field := range field {
        if i != 0 {
            setClause.WriteString(", ")
        }
        setClause.WriteString(field)
        setClause.WriteString(" = ?")
    }

    // Formar la consulta SQL completa
    query := fmt.Sprintf("UPDATE dentists SET %s WHERE id_dentist = ?", setClause.String())

    // Convertir los slices de strings en un slice de interfaces
	var interfaceValues []interface{}
	for _, v := range value {
		interfaceValues = append(interfaceValues, v)
	}

	// Concatenar el ID del dentista al final del slice de valores
	interfaceValues = append(interfaceValues, id)

	// Ejecutar la consulta SQL
	res, err := s.db.Exec(query, interfaceValues...)
	if err != nil {
		return nil, errors.New("No se ha podido ejecutar la consulta")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, errors.New("Ninguna fila fue cambiada")
	}

	updatedDentist, err := s.GetDentistById(id)
	if err != nil {
		return nil, err
	}

	return updatedDentist, nil
}

func (s *sqlStore) DeleteDentist(id int) error {
	query := "DELETE FROM dentists WHERE id_dentist = ?"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return err
	}

	return nil
}

func (s *sqlStore) GetDentistIdByLicense(license string) (int, error) {
	query := "SELECT id_dentist FROM dentists WHERE license = ?"
	row := s.db.QueryRow(query, license)

	var id int
	err := row.Scan(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Métodos para pacientes:

func (s *sqlStore) CreatePatient(patient domain.Patient) (*domain.Patient, error) {
	query := "INSERT INTO patients (first_name, last_name, address, dni, registration_date) VALUES (?, ?, ?, ?, ?)"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.RegistrationDate)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lastID, _ := res.LastInsertId()
	patient.ID = int(lastID)

	return &patient, nil
}

func (s *sqlStore) GetPatientById(id int) (*domain.Patient, error) {
	query := "SELECT * FROM patients WHERE id_patient = ?"
	row := s.db.QueryRow(query, id)

	var patient domain.Patient
	err := row.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.DNI, &patient.RegistrationDate)
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *sqlStore) UpdatePatient(patient domain.Patient) (*domain.Patient, error) {
	query := "UPDATE patients SET first_name = ?, last_name = ?, address = ?, dni = ?, registration_date = ? WHERE id_patient = ?"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.RegistrationDate, patient.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, errors.New("no se actualizaron filas")
	}

	query = "SELECT * FROM patients WHERE id_patient = ?"
	row := s.db.QueryRow(query, patient.ID)

	var updatedPatient domain.Patient
	err = row.Scan(&updatedPatient.ID, &updatedPatient.FirstName, &updatedPatient.LastName, &updatedPatient.Address, &updatedPatient.DNI, &updatedPatient.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return &updatedPatient, nil
}

func (s *sqlStore) UpdatePatientField(id int, field []string, value []string) (*domain.Patient, error) {
	// Verificar si los slices de campos y valores tienen la misma longitud
    if len(field) != len(value) {
        return nil, errors.New("la cantidad de campos y valores no coincide")
    }

	// Crear la parte de la consulta SQL para actualizar campos
    var setClause strings.Builder
    for i, field := range field {
        if i != 0 {
            setClause.WriteString(", ")
        }
        setClause.WriteString(field)
        setClause.WriteString(" = ?")
    }

    // Formar la consulta SQL completa
    query := fmt.Sprintf("UPDATE patients SET %s WHERE id_patient = ?", setClause.String())

    // Convertir los slices de strings en un slice de interfaces
	var interfaceValues []interface{}
	for _, v := range value {
		interfaceValues = append(interfaceValues, v)
	}

	// Concatenar el ID del dentista al final del slice de valores
	interfaceValues = append(interfaceValues, id)

	// Ejecutar la consulta SQL
	res, err := s.db.Exec(query, interfaceValues...)
	if err != nil {
		return nil, errors.New("No se ha podido ejecutar la consulta")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, errors.New("Ninguna fila fue cambiada")
	}

	updatedPatient, err := s.GetPatientById(id)
	if err != nil {
		return nil, err
	}

	return updatedPatient, nil
}

func (s *sqlStore) DeletePatient(id int) error {
	query := "DELETE FROM patients WHERE id_patient = ?"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return err
	}

	return nil
}

func (s *sqlStore) GetPatientIdByDNI(DNI string) (int, error) {
	query := "SELECT id_patient FROM patients WHERE dni = ?"
	row := s.db.QueryRow(query, DNI)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Métodos para turnos:

func (s *sqlStore) CreateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
	query := "INSERT INTO appointments (id_dentist, id_patient, date, time, description) VALUES (?, ?, ?, ?, ?)"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(appointment.IdDentist, appointment.IdPatient, appointment.Date, appointment.Time, appointment.Description)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lastID, _ := res.LastInsertId()
	appointment.ID = int(lastID)

	return &appointment, nil
}

func (s *sqlStore) GetAppointmentById(id int) (*domain.Appointment, error) {
	query := "SELECT * FROM appointments WHERE id_appointment = ?"
	row := s.db.QueryRow(query, id)

	var appointment domain.Appointment
	err := row.Scan(&appointment.ID, &appointment.IdDentist, &appointment.IdPatient, &appointment.Date, &appointment.Time, &appointment.Description)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (s *sqlStore) UpdateAppointment(appointment domain.Appointment) (*domain.Appointment, error) {
	query := "UPDATE appointments SET id_dentist = ?, id_patient = ?, date = ?, time = ?, description = ? WHERE id_appointment = ?"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(appointment.IdDentist, appointment.IdPatient, appointment.Date, appointment.Time, appointment.Description, appointment.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, errors.New("ninguna fila fue actualizada")
	}

	updatedAppointment, err := s.GetAppointmentById(appointment.ID)
	if err != nil {
		return nil, err
	}

	return updatedAppointment, nil
}

func (s *sqlStore) UpdateAppointmentField(id int, field []string, value []string) (*domain.Appointment, error) {
	// Verificar si los slices de campos y valores tienen la misma longitud
    if len(field) != len(value) {
        return nil, errors.New("la cantidad de campos y valores no coincide")
    }

	// Crear la parte de la consulta SQL para actualizar campos
    var setClause strings.Builder
    for i, field := range field {
        if i != 0 {
            setClause.WriteString(", ")
        }
        setClause.WriteString(field)
        setClause.WriteString(" = ?")
    }

    // Formar la consulta SQL completa
    query := fmt.Sprintf("UPDATE appointments SET %s WHERE id_appointment = ?", setClause.String())

    // Convertir los slices de strings en un slice de interfaces
	var interfaceValues []interface{}
	for _, v := range value {
		interfaceValues = append(interfaceValues, v)
	}

	// Concatenar el ID del dentista al final del slice de valores
	interfaceValues = append(interfaceValues, id)

	// Ejecutar la consulta SQL
	res, err := s.db.Exec(query, interfaceValues...)
	if err != nil {
		return nil, errors.New("No se ha podido ejecutar la consulta")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, errors.New("Ninguna fila fue cambiada")
	}

	updatedAppointment, err := s.GetAppointmentById(id)
	if err != nil {
		return nil, err
	}

	return updatedAppointment, nil
}

func (s *sqlStore) DeleteAppointment(id int) error {
	query := "DELETE FROM appointments WHERE id_appointment = ?"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return err
	}
	return nil
}

func (s *sqlStore) CreateAppointmentByDNIAndLicense(appointmentData domain.AppointmentData) (*domain.Appointment, error) {

	idDentist, err := s.GetDentistIdByLicense(appointmentData.License)
	if err != nil || idDentist == 0 {
		return nil, err
	}

	idPatient, err := s.GetPatientIdByDNI(appointmentData.DNI)
	if err != nil || idPatient == 0 {
		return nil, err
	}

	query := "INSERT INTO appointments (id_patient, id_dentist, date, time, description) VALUES (?, ?, ?, ?)"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(idDentist, idPatient, appointmentData.Date, appointmentData.Time, appointmentData.Description)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	appointmentID := int(lastID)

	appointment, err := s.GetAppointmentById(appointmentID)
	if err != nil {
		return nil, err
	}
	
	return appointment, nil
}

func (s *sqlStore) GetAppointmentsByDNI(DNI string) ([]domain.Appointment, error) {

	idPatient, err := s.GetPatientIdByDNI(DNI)
	if err != nil || idPatient == 0 {
		return nil, err
	}

	query := "SELECT * FROM appointments WHERE id_patient = ?"
	rows, err := s.db.Query(query, idPatient)
	if err != nil {
		return nil, err
	}

	listAppointments := []domain.Appointment{}

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(&appointment.ID, appointment.IdDentist, appointment.IdPatient, appointment.Date, appointment.Time, appointment.Description)
		if err != nil {
			return nil, err
		}
		listAppointments = append(listAppointments, appointment)
	}

	return listAppointments, nil
}
