package formatter

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(w http.ResponseWriter, status int, message string, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Status:  status,
		Message: message,
		Error:   err,
	})
}

func NoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func CreatedResponse(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusCreated,
		Message: message,
		Data:    data,
	})
}

func BadRequestResponse(w http.ResponseWriter, message string, err string) {
	ErrorResponse(w, http.StatusBadRequest, message, err)
}

func UnauthorizedResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusUnauthorized, message, "Unauthorized access")
}

func ForbiddenResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusForbidden, message, "Access forbidden")
}

func NotFoundResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusNotFound, message, "Resource not found")
}

func RedirectResponse(w http.ResponseWriter, url string, status int) {
	w.Header().Set("Location", url)
	w.WriteHeader(status)
}
