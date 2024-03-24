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

		createdAppointment, err := h.Service.CreateAppointment(appointment)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusCreated, gin.H{"turno" : createdAppointment})
	}
}

func (h *AppointmentHandler) GetAppointmentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		appointmentID, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("Ingrese un id valido")))
			return
		}

		appointment, err := h.Service.GetAppointmentById(appointmentID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"turno" : appointment})
	}
}

// UpdateAppointment
// UpdateAppointmentField
// DeleteAppointment
// CreateAppointmentByDNIAndLicense

func (h *AppointmentHandler) GetAppointmentsByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		DNI := c.Query("dni")

		appointments, err := h.Service.GetAppointmentsByDNI(DNI)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"turnos" : appointments})
	}
}
