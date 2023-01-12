package data

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func (r *Response) ToJSON() ([]byte, error) {

	_, ok := r.Data.([]Todo)
	_, okTodo := r.Data.(Todo)

	if !ok && !okTodo {
		return nil, fmt.Errorf("invalid data type")
	}

	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("error while parsing response: %v", err)
	}
	return data, nil
}
