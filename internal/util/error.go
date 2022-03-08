package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	ErrForbidden    = errors.New("access denied")
	ErrBadRequest   = errors.New("bad request")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

func ParseErrorToHTTPErrorCode(err error) int {
	switch {
	case errors.Is(err, ErrForbidden):
		return http.StatusForbidden
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}

func NewErrorResponse(logger *logrus.Entry, c *gin.Context, status int, message string) {
	switch status {
	case http.StatusInternalServerError:
		fallthrough
	default:
		logger.Error(message)
		status = http.StatusInternalServerError
	}

	c.AbortWithStatusJSON(status, map[string]interface{}{
		"status":  status,
		"message": message,
	})
}
