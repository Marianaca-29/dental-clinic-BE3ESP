package web

import ( 
	"net/http"
	"github.com/gin-gonic/gin"
)

type ErrorApi struct {
	Status int `json:"status"`
	Code string `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}
// retorna el mensaje de nuestro error
func (e *ErrorApi) Error() string {
	return e.Message
}

// si no encuentra el objeto
func NewNotFoundApiError(message string) error {
	return &ErrorApi {http.StatusNotFound, "not_found", message}
}

// si es una mal request, no existe el endpoint
func NewBadRequestApiError(message string) error {
	return &ErrorApi {http.StatusBadRequest, "bad_request", message}
}

// si hay un error del lado del servidor, error genérico
func NewInternalServerApiError(message string) error {
	return &ErrorApi {http.StatusInternalServerError, "internal_server_error", message}
}

// Success escribe una respuesta exitosa
func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data: data,
	})
}

// Failure escribe una respuesta fallida
func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, ErrorApi{
		Message: err.Error(),
		Status:  status,
		Code:    http.StatusText(status),
	})
}