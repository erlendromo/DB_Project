package utils

import (
	"errors"
	"log"
	"net/http"
)

// Validate errors
type ValidateErrors struct {
	Errors []struct {
		Error string `json:"error"`
	} `json:"errors"`
}

// NewValidateErrors function for creating new ValidateErrors
func NewValidateErrors(errs map[string]error) *ValidateErrors {
	ve := &ValidateErrors{}

	for _, err := range errs {
		ve.Errors = append(ve.Errors, struct {
			Error string `json:"error"`
		}{Error: err.Error()})
	}

	return ve
}

// UnauthorizedError function for logging actual error and returning generic error
func NewUnauthorizedError(err error) error {
	if err != nil {
		log.Printf("unauthorized access: %v", err)
	}

	return errors.New("unauthorized access")
}

// InternalServerError function for logging actual error and returning generic error
func NewInternalServerError(err error) error {
	if err != nil {
		log.Printf("internal server error: %v", err)
	}

	return errors.New("internal server error")
}

// ErrorResponse struct
//
//	@title			ErrorResponse
//	@description	Error Response with message and statuscode
type ErrorResponse struct {
	Message    string `json:"message"`
	Statuscode int    `json:"statuscode"`
}

// NewErrorResponse function for creating new ErrorResponse
func NewErrorResponse(statuscode int, err error) *ErrorResponse {
	if err == nil {
		err = errors.New("internal server error")
	}

	return &ErrorResponse{
		Message:    err.Error(),
		Statuscode: statuscode,
	}
}

// ERROR function for writing error response as json
func ERROR(w http.ResponseWriter, statuscode int, err error) {
	JSON(w, statuscode, NewErrorResponse(statuscode, err))
}
