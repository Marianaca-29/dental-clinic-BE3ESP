package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/internal/patient"
	"DENTAL-CLINIC/pkg/web"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	service patient.IService
}

func NewProductHandler(service patient.IService) *patientHandler {
	return &patientHandler{
		service: service,
	}
}

func (h *patientHandler) CreatePatient() gin.HandlerFunc {
	return func(c *gin.Context) { 
		var patient domain.Patient
		err := c.ShouldBindJSON(&patient)
		if (err != nil) {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Datos del paciente mal ingresados"))
			return
		}

		patientResponse, err := h.service.CreatePatient(patient)
		if (err != nil) {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("No se ha podido crear el paciente"))
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"paciente" : patientResponse})
	}
}

func (h *patientHandler) GetPatientById() gin.HandlerFunc {
	return func(c *gin.Context) { 
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if (err != nil) {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Id invalido"))
			return
		}

		patientResponse, err := h.service.GetPatientById(id)
		if (err != nil) {
			var errAPI web.ErrorApi
			if (err == &errAPI) {
				c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el paciente con el id %d", id)))
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"paciente" : patientResponse})
	}
}

func (h *patientHandler) UpdatePatient() gin.HandlerFunc {
	return func(c *gin.Context) { 

	}
}

func (h *patientHandler) UpdatePatientField() gin.HandlerFunc {
	return func(c *gin.Context) { 

	}
}

func (h *patientHandler) DeletePatient() gin.HandlerFunc {
	return func(c *gin.Context) { 

	}
}