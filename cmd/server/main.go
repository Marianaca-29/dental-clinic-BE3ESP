package main

import (
	"DENTAL-CLINIC/internal/dentist"
	"DENTAL-CLINIC/pkg/store"
	"DENTAL-CLINIC/cmd/server/handler"
	"database/sql"
	"log"
	"github.com/gin-gonic/gin"
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

	storage := store.NewSqlStore(db)
	repo := dentist.NewRepository(storage)
	service := dentist.NewService(repo)
	dentistHandler := handler.NewDentistHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentist := r.Group("/dentists")
	{
		dentist.GET(":id", dentistHandler.GetByDentistID())
		dentist.POST("", dentistHandler.Post())
		//Completar otros 
	}

	r.Run(":8080")
}