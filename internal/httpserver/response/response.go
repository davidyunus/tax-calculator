package response

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse ...
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// JSON writes json http response
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Error writes Error HTTP Response
func Error(w http.ResponseWriter, status int, data string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var errorCode string
	switch status {
	case http.StatusBadRequest:
		errorCode = "BadRequest"
	case http.StatusInternalServerError:
		errorCode = "InternalServerError"
	}

	if status == http.StatusInternalServerError {
		json.NewEncoder(w).Encode(ErrorResponse{
			Code:    errorCode,
			Message: "Internal Server Error",
		})

	} else {
		json.NewEncoder(w).Encode(ErrorResponse{
			Code:    errorCode,
			Message: data,
		})
	}
}
