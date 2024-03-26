package handler

import(
	"strconv"
	"net/http"
	"fmt"

	"DENTAL-CLINIC/internal/dentist"
	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	service dentist.IService
}

func NewDentistHandler(service dentist.IService) *dentistHandler {
	return &dentistHandler{
		service: service,
	}
}

// GetByDentistID godoc
// @Summary Get dentist by ID
// @Description Retrieve dentist's data by their ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Dentist ID"
// @Success 200 {object} domain.Dentist "Returns the requested dentist"
// @Failure 400 {object} web.ErrorApi "Bad request"
// @Failure 404 {object} web.ErrorApi "Not found"
// @Router /dentists/{id} [get]
func (h *dentistHandler) GetByDentistID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("Ingrese un id valido")))
			return
		}

		dentist, err := h.service.GetDentistById(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el dentista con el id %d", id)))
			return
		}

		c.JSON(http.StatusOK, gin.H{"dentista" : dentist})
	}
}

// validateNotEmpty valida que los campos no estén vacíos
func validateNotEmptyDentist(dentist *domain.Dentist) (bool, error) {
	switch {
	case dentist.FirstName == "":
		return false, web.NewBadRequestApiError("el nombre del dentista no puede estar vacío")
	case dentist.License == "":
		return false, web.NewBadRequestApiError("la matrícula del dentista no puede estar vacía")
	}
	return true, nil
}

// CreateDentist godoc
// @Summary Create a new dentist
// @Description Create a new dentist with the provided data
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param dentist body domain.Dentist true "Dentist data"
// @Success 201 {object} domain.Dentist "Returns the created dentist"
// @Failure 400 {object} web.ErrorApi "Bad request"
// @Router /dentists [post]
func (h *dentistHandler) CreateDentist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Datos del dentista mal ingresados"))
			return
		}
		valid, err := validateNotEmptyDentist(&dentist)
		if !valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
			return
		}

		d, err := h.service.CreateDentist(dentist)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
			return
		}
		c.JSON(http.StatusCreated, gin.H{"dentista" : d})
	}
}

// UpdateDentist godoc
// @Summary Update a dentist
// @Description Update a dentist's data with the provided data
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param dentist body domain.Dentist true "Updated dentist data"
// @Success 200 {object} domain.Dentist "Returns the updated dentist"
// @Failure 400 {object} web.ErrorApi "Bad request"
// @Failure 404 {object} web.ErrorApi "Not found"
// @Router /dentists [put]
func (h *dentistHandler) UpdateDentist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("JSON invalido")))
			return
		}

		_, errNotFound := h.service.GetDentistById(dentist.ID)
		if errNotFound != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el dentista con el id %d", dentist.ID)))
			return
		}

		updatedDentist, errUpdated := h.service.UpdateDentist(dentist)
		if errUpdated != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(errUpdated.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"dentista" : updatedDentist})
	}
}

// UpdateDentistField godoc
// @Summary Update a specific field of a dentist
// @Description Update a specific field of a dentist's data with the provided data
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param dentist body domain.Dentist true "Updated dentist field data"
// @Success 200 {object} domain.Dentist "Returns the updated dentist"
// @Failure 400 {object} web.ErrorApi "Bad request"
// @Failure 404 {object} web.ErrorApi "Not found"
// @Router /dentists [patch]
func (h *dentistHandler) UpdateDentistField() gin.HandlerFunc {
	return func(c *gin.Context) { 
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("JSON invalido")))
			return
		}

		_, errNotFound := h.service.GetDentistById(dentist.ID)
		if errNotFound != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el dentista con el id %d", dentist.ID)))
			return
		}

		updatedDentist, err := h.service.UpdateDentistField(dentist)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"dentista" : updatedDentist})
	}
}

// DeleteDentist godoc
// @Summary Delete a dentist
// @Description Delete a dentist by their ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Dentist ID"
// @Success 200 {object} gin.H "{'message': 'Dentist deleted'}"
// @Failure 400 {object} web.ErrorApi "Bad request"
// @Failure 500 {object} web.ErrorApi "Internal server error"
// @Router /dentists/{id} [delete]
func (h *dentistHandler) DeleteDentist() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		dentistId, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewNotFoundApiError(fmt.Sprintf("No existe el dentista con el id %d", dentistId)))
			return
		}
		errDelete := h.service.DeleteDentist(dentistId)
		if errDelete != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.NewInternalServerApiError(errDelete.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "dentista borrado"})
	}
}