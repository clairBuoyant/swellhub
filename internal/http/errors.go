package http

import (
	"encoding/json"
	"errors"
	"net/http"

	ie "github.com/clairBuoyant/swellhub/internal/errors"
)

// ErrorResponse represents the structure of error responses.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandleError handles application errors and writes appropriate HTTP responses.
func HandleError(w http.ResponseWriter, err error) {
	var appErr *ie.AppError
	if ok := errors.As(err, &appErr); ok {
		respondWithError(w, appErr.Code, appErr.Message)
		return
	}
	respondWithError(w, http.StatusInternalServerError, "internal server error")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Code: code, Message: message})
}
