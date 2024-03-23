package main

import (
	"DENTAL-CLINIC/cmd/server/handler"
	"DENTAL-CLINIC/internal/appointment"
	"DENTAL-CLINIC/internal/dentist"
	"DENTAL-CLINIC/internal/patient"
	"DENTAL-CLINIC/pkg/store"
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

	dentistRepo := dentist.NewRepository(storage)
	patientRepo := patient.NewRepository(storage)
	appointmentRepo := appointment.NewRepository(storage)

	dentistService := dentist.NewService(dentistRepo)
	patientService := patient.NewService(patientRepo)
	appointmentService := appointment.NewService(appointmentRepo)

	dentistHandler := handler.NewDentistHandler(dentistService)
	patientHandler := handler.NewPatientHandler(patientService)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService)
	
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetByDentistID())
		dentists.POST("", dentistHandler.Post())
		//Completar otros 
	}
	patients := r.Group("/patients")
	{
		
		patients.GET(":id", patientHandler.GetPatientById())
		patients.POST("", patientHandler.Post())
		patients.PATCH(":id", patientHandler.Put())
		patients.PUT(":id", patientHandler.Patch())
		patients.DELETE(":id", patientHandler.Delete())
		
	}

	appointments := r.Group("/appointments")
	{
		appointments.GET(":id", appointmentHandler.GetById())
		appointments.GET("DNI", appointmentHandler.GetByDNI())
		appointments.POST("", appointmentHandler.POST())
	}
	

	r.Run(":8080")
}