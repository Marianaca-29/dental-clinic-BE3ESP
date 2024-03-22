package web

import "net/http"

type ErrorApi struct {
	Status int `json:"status"`
	Code string `json:"code"`
	Message string `json:"message"`
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
