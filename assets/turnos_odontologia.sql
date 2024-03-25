create database turnos_odontologia;

use turnos_odontologia;

CREATE TABLE dentists (
    id_dentist INTEGER PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    license VARCHAR(255)
);

CREATE TABLE patients (
    id_patient INTEGER PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    address VARCHAR(255),
    dni VARCHAR(255),
    registration_date VARCHAR(10)
);

CREATE TABLE appointments (
    id_appointment INTEGER PRIMARY KEY AUTO_INCREMENT,
    id_patient INTEGER,
    id_dentist INTEGER,
    date VARCHAR(10),
    time VARCHAR(8),
    description VARCHAR(255),
    FOREIGN KEY (id_patient) REFERENCES patients(id_patient),
    FOREIGN KEY (id_dentist) REFERENCES dentists(id_dentist)
);


