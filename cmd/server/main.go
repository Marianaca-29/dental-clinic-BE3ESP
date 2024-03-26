package main

import (
	"DENTAL-CLINIC/cmd/server/docs"
	"DENTAL-CLINIC/cmd/server/handler"
	"DENTAL-CLINIC/internal/appointment"
	"DENTAL-CLINIC/internal/dentist"
	"DENTAL-CLINIC/internal/patient"
	"DENTAL-CLINIC/pkg/middleware"
	"DENTAL-CLINIC/pkg/store"

	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title DENTAL-CLINIC PAULANA'S OFFICE -  API DOCUMENTATION
// @version 1.0
// @description This API handles appointments with patients and dentists data.
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones
// @contact.name Paulina Oberti Busso (paulinaobertibusso@gmail.com)- Mariana Cañas (mariana.famaf@gmail.com)
// @contact.url https://developers.ctd.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error al cargar el archivo .env")
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/turnos_odontologia")
	if err != nil {
		log.Fatal(err)
	} 

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
	docs.SwaggerInfo.Host =  os.Getenv("HOST") 	
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		appointments.POST("postAppointment", middleware.Authentication(), appointmentHandler.CreateAppointment())
		appointments.PUT("updateAppointment", middleware.Authentication(), appointmentHandler.UpdateAppointment())
		appointments.PATCH("updateAppointmentField", middleware.Authentication(), appointmentHandler.UpdateAppointmentField())
		appointments.DELETE("deleteAppointment/:id", middleware.Authentication(), appointmentHandler.DeleteAppointment())
		appointments.POST("postAppointmentDNILicense", middleware.Authentication(), appointmentHandler.CreateAppointmentByDNIAndLicense())
		appointments.GET("getByDNI", appointmentHandler.GetAppointmentsByDNI())
	}

	r.Run(":8080")
}