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

func NewPatientHandler(service patient.IService) *patientHandler {
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

		valid, err := validateNotEmptyPatient(&patient)
		if !valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
			return
		}

		patientResponse, err := h.service.CreatePatient(patient)
		if (err != nil) {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Ese DNI ya existe"))
			return
		}
		
		c.JSON(http.StatusCreated, gin.H{"paciente" : patientResponse})
	}
}

// validateNotEmpty valida que los campos no estén vacíos
func validateNotEmptyPatient(patient *domain.Patient) (bool, error) {
	switch {
	case patient.FirstName == "":
		return false, web.NewBadRequestApiError("el nombre del paciente no puede estar vacío")
	case patient.LastName == "":
		return false, web.NewBadRequestApiError("el apellido del paciente no puede estar vacío")
	case patient.DNI == "":
		return false, web.NewBadRequestApiError("el DNI del paciente no puede estar vacío")
	case patient.Address == "":
		return false, web.NewBadRequestApiError("la dirección del paciente no puede estar vacía")
	case patient.RegistrationDate == "":
		return false, web.NewBadRequestApiError("la fecha de alta del paciente no puede estar vacía")
	}
	return true, nil
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
		var patient domain.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("JSON invalido")))
			return
		}

		_, errNotFound := h.service.GetPatientById(patient.ID)
		if errNotFound != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el paciente con el id %d", patient.ID)))
			return
		}

		updatedPatient, err := h.service.UpdatePatient(patient)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"paciente" : updatedPatient})
	}
}

func (h *patientHandler) UpdatePatientField() gin.HandlerFunc {
	return func(c *gin.Context) { 
		var patient domain.Patient
		err := c.ShouldBindJSON(&patient)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("JSON invalido")))
			return
		}

		_, errNotFound := h.service.GetPatientById(patient.ID)
		if errNotFound != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el paciente con el id %d", patient.ID)))
			return
		}

		updatedPatient, err := h.service.UpdatePatientField(patient)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"paciente" : updatedPatient})
	}
}

func (h *patientHandler) DeletePatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		patientID, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("Ingrese un id valido")))
			return
		}

		if err := h.service.DeletePatient(patientID); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "paciente borrado"})
	}
}