package data

import (
	"log"
	"net/http"
)

type ClientResponse interface {
	ToJSON() ([]byte, error)
}

func Respond(w http.ResponseWriter, payload ClientResponse) {
	w.Header().Set("Content-Type", "appliction/json")
	if errorResponse, ok := payload.(*ErrorResponse); ok {
		w.WriteHeader(errorResponse.Status)
	}
	data, err := payload.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}
