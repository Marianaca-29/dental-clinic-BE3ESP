package handler

import (
	"DENTAL-CLINIC/internal/appointment"
	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/web"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	Service appointment.IService
}

func NewAppointmentHandler(service appointment.IService) *AppointmentHandler {
	return &AppointmentHandler{Service: service}
}

func (h *AppointmentHandler) CreateAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("JSON invalido")))
			return
		}

		values, err := validateNotEmptyAppointment(&appointment)
		if !values {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
			return
		}

		createdAppointment, err := h.Service.CreateAppointment(appointment)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusCreated, gin.H{"turno": createdAppointment})
	}
}

// validateNotEmpty valida que los campos no estén vacíos
func validateNotEmptyAppointment(appointment *domain.Appointment) (bool, error) {
	switch {
	case appointment.IdDentist == 0:
		return false, web.NewBadRequestApiError("el id del dentista no puede estar vacío")
	case appointment.IdPatient == 0:
		return false, web.NewBadRequestApiError("el id del paciente no puede estar vacío")
	case appointment.Date == "":
		return false, web.NewBadRequestApiError("la fecha no puede estar vacía")
	case appointment.Time == "":
		return false, web.NewBadRequestApiError("el horario no puede estar vacío")
	case appointment.Description == "":
		return false, web.NewBadRequestApiError("la descripción no puede estar vacía")
	}
	return true, nil
}

func (h *AppointmentHandler) GetAppointmentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		appointmentID, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(fmt.Sprintf("Ingrese un id valido")))
			return
		}

		appointment, err := h.Service.GetAppointmentById(appointmentID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"turno": appointment})
	}
}

func (h *AppointmentHandler) UpdateAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		if err := c.ShouldBindJSON(&appointment); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("JSON invalido")))
			return
		}

		_, errNotFound := h.Service.GetAppointmentById(appointment.ID)
		if errNotFound != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el turno con el id %d", appointment.ID)))
			return
		}

		updatedAppointment, err := h.Service.UpdateAppointment(appointment)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"turno" : updatedAppointment})
	}
}

func (h *AppointmentHandler) UpdateAppointmentField() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("JSON invalido")))
			return
		}

		_, errNotFound := h.Service.GetAppointmentById(appointment.ID)
		if errNotFound != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el turno con el id %d", appointment.ID)))
			return
		}

		updatedAppointment, err := h.Service.UpdateAppointmentField(appointment)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"turno" : updatedAppointment})
	}
}

func (h *AppointmentHandler) DeleteAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		appointmentId, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el turno con el id %d", appointmentId)))
			return
		}

		errDelete := h.Service.DeleteAppointment(appointmentId)
		if errDelete != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(errDelete.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "turno borrado"})
	}
}

func (h *AppointmentHandler) CreateAppointmentByDNIAndLicense() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointmentData domain.AppointmentData
		err := c.BindJSON(&appointmentData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("JSON invalido")))
			return
		}

		values, err := validateNotEmptyAppointmentDNILicense(&appointmentData)
		if !values {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
			return
		}

		appointments, err := h.Service.CreateAppointmentByDNIAndLicense(appointmentData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"turnos": appointments})
	}
}

// validateNotEmpty valida que los campos no estén vacíos
func validateNotEmptyAppointmentDNILicense(appointment *domain.AppointmentData) (bool, error) {
	switch {
	case appointment.DNI == "":
		return false, web.NewBadRequestApiError("el DNI del paciente no puede estar vacío")
	case appointment.License == "":
		return false, web.NewBadRequestApiError("la matricula del dentista no puede estar vacía")
	case appointment.Date == "":
		return false, web.NewBadRequestApiError("la fecha no puede estar vacía")
	case appointment.Time == "":
		return false, web.NewBadRequestApiError("el horario no puede estar vacío")
	case appointment.Description == "":
		return false, web.NewBadRequestApiError("la descripción no puede estar vacía")
	}
	return true, nil
}

func (h *AppointmentHandler) GetAppointmentsByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		DNI := c.Query("dni")
		if DNI == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Debe ingresar el DNI del paciente"))
			return
		}

		appointments, err := h.Service.GetAppointmentsByDNI(DNI)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"turnos": appointments})
	}
}
