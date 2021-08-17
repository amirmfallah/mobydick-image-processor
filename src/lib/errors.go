package lib

import "encoding/json"

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
