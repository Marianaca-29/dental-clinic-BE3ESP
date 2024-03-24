package main

import (
	"DENTAL-CLINIC/cmd/server/handler"
	"DENTAL-CLINIC/internal/appointment"
	"DENTAL-CLINIC/internal/dentist"
	"DENTAL-CLINIC/internal/patient"
	"DENTAL-CLINIC/pkg/store"
	"DENTAL-CLINIC/pkg/middleware"

	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error al cargar el archivo .env")
	}

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
		dentists.GET("getDentist/:id", dentistHandler.GetByDentistID())
		dentists.POST("postDentist", middleware.Authentication(), dentistHandler.CreateDentist())
		dentists.PUT("updateDentist", middleware.Authentication(), dentistHandler.UpdateDentist())
		dentists.PATCH("updateDentistField", middleware.Authentication(), dentistHandler.UpdateDentistField())
		dentists.DELETE("deleteDentist/:id", middleware.Authentication(), dentistHandler.DeleteDentist())
	}

	patients := r.Group("/patients")
	{
		patients.GET("getPatient/:id", patientHandler.GetPatientById())
		patients.POST("postPatient", middleware.Authentication(), patientHandler.CreatePatient())
		patients.PATCH("updatePatientField", middleware.Authentication(), patientHandler.UpdatePatient())
		patients.PUT("updatePatient", middleware.Authentication(), patientHandler.UpdatePatientField())
		patients.DELETE("deletePatient/:id", middleware.Authentication(), patientHandler.DeletePatient())
	}

	appointments := r.Group("/appointments")
	{
		appointments.GET("getAppointment/:id", appointmentHandler.GetAppointmentById())
		appointments.POST("postAppointment", appointmentHandler.CreateAppointment())
		appointments.PUT("updateAppointment", appointmentHandler.UpdateAppointment())
		appointments.PATCH("updateAppointmentField", appointmentHandler.UpdateAppointmentField())
		appointments.DELETE("deleteAppointment/:id", appointmentHandler.DeleteAppointment())
		appointments.POST("postAppointmentDNILicense", appointmentHandler.CreateAppointmentByDNIAndLicense())
		appointments.GET("getByDNI", appointmentHandler.GetAppointmentsByDNI())
	}
	

	r.Run(":8080")
}