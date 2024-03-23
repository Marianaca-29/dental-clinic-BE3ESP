package handler

import (
	"DENTAL-CLINIC/internal/appointment"
	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/web"

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

func (h *AppointmentHandler) POST() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		if err := c.ShouldBindJSON(&appointment); err != nil {
			web.Failure(c, http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
			return
		}

		createdAppointment, err := h.Service.CreateAppointment(appointment)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		web.Success(c, http.StatusCreated, createdAppointment)
	}
}

func (h *AppointmentHandler) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		appointmentID, err := strconv.Atoi(id)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, web.NewBadRequestApiError("Invalid appointment ID"))
			return
		}

		appointment, err := h.Service.GetAppointmentById(appointmentID)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		web.Success(c, http.StatusOK, appointment)
	}
}

func (h *AppointmentHandler) GetByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		DNI := c.Query("DNI")

		appointments, err := h.Service.GetAppointmentsByDNI(DNI)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		web.Success(c, http.StatusOK, appointments)
	}
}

//completar