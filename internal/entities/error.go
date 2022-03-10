package entities

import "errors"

var (
	ErrForbidden    = errors.New("access denied")
	ErrBadRequest   = errors.New("bad request")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
