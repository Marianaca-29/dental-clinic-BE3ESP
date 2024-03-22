package main

import (
	"database/sql"
	"log"

	// "DENTAL-CLINIC/pkg/store"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos_odontologia")
	if err != nil {
		log.Fatal(err)
	} 

	// db.Ping() en Go se utiliza para verificar si la conexión con la base de datos está activa y disponible
	errPing := db.Ping()
	if errPing != nil {
		log.Fatal(errPing)
	}

	//storage := store.NewSqlStore(db)
}