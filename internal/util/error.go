package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"repositorie/internal/entities"

	"github.com/sirupsen/logrus"
)

func ParseErrorToHTTPErrorCode(err error) int {
	switch {
	case errors.Is(err, entities.ErrForbidden):
		return http.StatusForbidden
	case errors.Is(err, entities.ErrBadRequest):
		return http.StatusBadRequest
	case errors.Is(err, entities.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, entities.ErrUnauthorized):
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}

func NewErrorResponse(logger *logrus.Entry, w http.ResponseWriter, status int, message string) {
	switch status {
	case http.StatusInternalServerError:
		fallthrough
	default:
		logger.Error(message)
		status = http.StatusInternalServerError
	}

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(&entities.Response{
		Status:  status,
		Message: message,
	})
}
