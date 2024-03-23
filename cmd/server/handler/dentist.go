package handler

import(
	"DENTAL-CLINIC/internal/dentist"
	"DENTAL-CLINIC/internal/domain"
	"DENTAL-CLINIC/pkg/web"
	"github.com/gin-gonic/gin"
	"errors"
	"strconv"
)
type dentistHandler struct {
	service dentist.IService
}

func NewDentistHandler(service dentist.IService) *dentistHandler {
	return &dentistHandler{
		service: service,
	}
}

func (h *dentistHandler) GetByDentistID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id no válido"))
			return
		}
		product, err := h.service.GetDentistById(id)
		if err != nil {
			web.Failure(c, 404, errors.New("product no encontrado"))
			return
		}
		web.Success(c, 200, product)
	}
}
// validateNotEmpty valida que los campos no estén vacíos
func validateNotEmpty(dentist *domain.Dentist) (bool, error) {
	switch {
	case dentist.FirstName == "":
		return false, web.NewBadRequestApiError("el nombre del dentista no puede estar vacío")
	case dentist.License == "":
		return false, web.NewBadRequestApiError("la matrícula del dentista no puede estar vacía")
	}
	return true, nil
}

// Post crea un nuevo producto
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, web.NewBadRequestApiError("invalid json"))
			return
		}
		valid, err := validateNotEmpty(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.service.CreateDentist(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Completar
// Delete elimina un producto
func (h *dentistHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// Put actualiza un producto
func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// Patch actualiza un producto o alguno de sus campos
func (h *dentistHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}