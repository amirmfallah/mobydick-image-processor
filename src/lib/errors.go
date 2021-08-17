package lib

import (
	"encoding/json"
	"errors"
)

var (
	ErrDatabaseConnection = errors.New("Couldn't connect to database.")
	ErrNotFound           = errors.New("Not found.")
	ErrAlreadyExists      = errors.New("Already exists.")
	ErrForbidden          = errors.New("Forbidden.")
	ErrInternal           = errors.New("Internal error.")
	ErrBadRequest         = errors.New("Bad Request.")
)

type ErrorResponse struct {
	Type        string `json:"@type"`
	Code        int    `json:"statusCode"`
	Title       string `json:"title"`
	Description string `json:"message"`
}

func NewErrorResponse(code int, title, description string) *ErrorResponse {
	err := new(ErrorResponse)
	err.Type = "Error"
	err.Code = code
	err.Title = title
	err.Description = description
	return err
}

func (errResponse ErrorResponse) ToJsonString() (string, error) {
	jsonBytes, err := json.Marshal(errResponse)
	return string(jsonBytes), err
}
