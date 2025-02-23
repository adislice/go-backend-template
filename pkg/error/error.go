package pkg_error

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ClientError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *ClientError) Error() string {
	return e.Message
}

func BadRequest(err string) *ClientError {
	return &ClientError{
		Status:  http.StatusBadRequest,
		Message: err,
	}
}

func InternalServerError(err string) *ClientError {
	return &ClientError{
		Status:  http.StatusInternalServerError,
		Message: err,
	}
}

func NotFound(err string) *ClientError {
	return &ClientError{
		Status:  http.StatusNotFound,
		Message: err,
	}
}

func Forbidden(err string) *ClientError {
	return &ClientError{
		Status:  http.StatusForbidden,
		Message: err,
	}
}

func Unauthorized(err string) *ClientError {
	return &ClientError{
		Status:  http.StatusUnauthorized,
		Message: err,
	}
}

func HandleServiceError(c *fiber.Ctx, err error) error {
	var clientErr *ClientError
	if errors.As(err, &clientErr) {
		return c.Status(clientErr.Status).JSON(clientErr)
	}

	return c.Status(fiber.StatusInternalServerError).JSON(InternalServerError(err.Error()))
}
