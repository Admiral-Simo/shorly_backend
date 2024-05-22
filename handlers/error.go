package handlers

import "github.com/gofiber/fiber/v2"

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func (e *Error) Error() string {
	return e.Err
}

func NewError(code int, err string) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	var (
		apiError *Error
		ok       bool
	)
	if apiError, ok = err.(*Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiError = NewError(fiber.StatusInternalServerError, "internal server error")
	return c.Status(apiError.Code).JSON(apiError)
}

func ErrUnAuthorized() *Error {
	return NewError(fiber.StatusUnauthorized, "unauthorized request")
}

func ErrInvalidId() *Error {
	return NewError(fiber.StatusBadRequest, "invalid id given")
}

func ErrBadRequest() *Error {
	return NewError(fiber.StatusBadRequest, "invalid JSON request")
}

func ErrNotFound(resource string) *Error {
	return NewError(fiber.StatusNotFound, resource+" not found")
}

func ErrUnavailable(resource string) *Error {
	return NewError(fiber.StatusNotFound, resource+" unavailable at the moment")
}

func ErrAlreadyExists(resource string) *Error {
	return NewError(fiber.StatusConflict, resource+" already exists")
}

func ErrInternalServerError() *Error {
	return NewError(fiber.StatusInternalServerError, "internal server error")
}

func NotFoundHandler(c *fiber.Ctx) error {
	errorResponse := ErrNotFound(c.OriginalURL())
	return c.Status(errorResponse.Code).JSON(errorResponse)
}
