package store

import (
	"DENTAL-CLINIC/internal/domain"
	"database/sql"
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

func (s *sqlStore) CreateDentist (dentist domain.Dentist) (*domain.Dentist, error) {
	query := "INSERT INTO dentists (first_name, last_name, license) VALUES (?, ?, ?)"

	stmt, err := s.db.Prepare(query)
	if (err != nil) {
		return nil, err
	}

	res, err := stmt.Exec(dentist.FirstName, dentist.LastName, dentist.License)
	if (err != nil) {
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

func (s *sqlStore) GetDentistById (id int) (*domain.Dentist, error) {
	query:= "SELECT * FROM dentists WHERE id_dentist = ?"
	row := s.db.QueryRow(query, id)

	var dentist domain.Dentist
	err := row.Scan(&dentist.ID, &dentist.FirstName, &dentist.LastName, &dentist.License)
	if err != nil {
		return nil, err
	}

	return &dentist, nil
}

func (s *sqlStore) UpdateDentist (dentist domain.Dentist) (*domain.Dentist, error) {
	query := "UPDATE dentists SET firts_name = ?, last_name = ?, license = ? WHERE id_dentist = ?"

	stmt, err := s.db.Prepare(query)
	if (err != nil) {
		return nil, err
	}

	res, err := stmt.Exec(dentist.FirstName, dentist.LastName, dentist.License, dentist.ID)
	if (err != nil) {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if (rowsAffected == 0) {
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

func (s *sqlStore) UpdateDentistField(dentist domain.Dentist) (*domain.Dentist, error) {
    // Definir la consulta base y los argumentos
    query := "UPDATE dentists SET "
    args := []interface{}{}

    // Actualizar el campo FirstName si está presente en el dentista proporcionado
    if dentist.FirstName != "" {
        query += "first_name = ? "
        args = append(args, dentist.FirstName)
    }

    // Actualizar el campo LastName si está presente en el dentista proporcionado
    if dentist.LastName != "" {
        if len(args) > 0 {
            query += ", "
        }
        query += "last_name = ? "
        args = append(args, dentist.LastName)
    }

    // Agregar la condición WHERE
    query += "WHERE id_dentist = ?"
    args = append(args, dentist.ID)

    // Preparar la consulta
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return nil, err
    }

    // Ejecutar la consulta
    _, err = stmt.Exec(args...)
    if err != nil {
        return nil, err
    }

    // Consultar el dentista actualizado
    query = "SELECT id_dentist, first_name, last_name, license FROM dentists WHERE id_dentist = ?"
    row := s.db.QueryRow(query, dentist.ID)

    // Escanear los resultados en una estructura de dentista
    var updatedDentist domain.Dentist
    err = row.Scan(&updatedDentist.ID, &updatedDentist.FirstName, &updatedDentist.LastName, &updatedDentist.License)
    if err != nil {
        return nil, err
    }

    return &updatedDentist, nil
}

func (s *sqlStore) DeleteDentist (id int) (error) {
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
	} else if (rowsAffected == 0) {
		return err
	}

	return nil
}

func (s *sqlStore) GetDentistIdByLicense (license string) (int, error) {
	query:= "SELECT id_dentist FROM dentists WHERE license = ?"
	row := s.db.QueryRow(query, license)

	var id int
	err := row.Scan(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}


// Métodos para pacientes:

func (s *sqlStore) CreatePatient (patient domain.Patient) (*domain.Patient, error) {
	return nil, nil
}
func (s *sqlStore) GetPatientById (id int) (*domain.Patient, error) {
	return nil, nil
}
func (s *sqlStore) UpdatePatient (patient domain.Patient) (*domain.Patient, error) {
	return nil, nil
}
func (s *sqlStore) UpdatePatientField (patient domain.Patient) (*domain.Patient, error) {
	return nil, nil
} 
func (s *sqlStore) DeletePatient (id int) (error) {
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
	} else if (rowsAffected == 0) {
		return err
	}

	return nil
}

func (s *sqlStore) GetPatientIdByDNI (DNI string) (int, error) {
	query:= "SELECT id_patient FROM patients WHERE dni = ?"
	row := s.db.QueryRow(query, DNI)

	var id int
	err := row.Scan(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Métodos para turnos:

func (s *sqlStore) CreateAppointment (appointment domain.Appointment) (*domain.Appointment, error) {
	return nil, nil
}
func (s *sqlStore) GetAppointmentById (id int) (*domain.Appointment, error) {
	return nil, nil
}
func (s *sqlStore) UpdateAppointment (appointment domain.Appointment) (*domain.Appointment, error) {
	return nil, nil
}
func (s *sqlStore) UpdateAppointmentField (appointment domain.Appointment) (*domain.Appointment, error) {
	return nil, nil
}

func (s *sqlStore) DeleteAppointment (id int) (error) {
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
	} else if (rowsAffected == 0) {
		return err
	}

	return nil
}

func (s *sqlStore) CreateAppointmentByDNIAndLicense (DNI string, license string, appointment domain.Appointment) (*domain.Appointment, error) {
	
	idDentist, err := s.GetDentistIdByLicense(license)
	if err != nil || idDentist == 0 {
		return nil, err
	} 

	idPatient, err := s.GetPatientIdByDNI(DNI)
	if err != nil || idPatient == 0 {
		return nil, err
	} 

	query := "INSERT INTO appointments (id_patient, id_dentist, date, time) VALUES (?, ?, ?, ?)"

	stmt, err := s.db.Prepare(query)
	if (err != nil) {
		return nil, err
	}

	res, err := stmt.Exec(idDentist, idPatient, appointment.Date, appointment.Time)
	if (err != nil) {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lastId, _ := res.LastInsertId()
	appointment.ID = int(lastId)

	return &appointment, nil
}

func (s *sqlStore) GetAppointmentsByDNI (DNI string) ([]domain.Appointment, error) {

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

