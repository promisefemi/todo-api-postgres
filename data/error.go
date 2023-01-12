package data

import (
	"encoding/json"
	"fmt"
	"log"
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
	Cause   error  `json:"cause,omitempty"`
}

func (e *ErrorResponse) ToJSON() ([]byte, error) {
	log.Println(e.Cause.Error())
	data, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("error while parsing error response: %v", err)
	}
	return data, nil
}

// NewErrorResponse returns a new instance of ErrorResponse
func NewErrorResponse(cause error, status int, message string) *ErrorResponse {
	log.Println(cause)
	return &ErrorResponse{
		Cause:   cause,
		Status:  status,
		Message: message,
	}
}
