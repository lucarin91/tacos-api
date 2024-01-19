package models

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func WriteInternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	err = json.NewEncoder(w).Encode(Error{
		Type:    "internal_error",
		Message: err.Error(),
	})
	if err != nil {
		panic(err)
	}
}

func WriteClientError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	err = json.NewEncoder(w).Encode(Error{
		Type:    "client_error",
		Message: err.Error(),
	})
	if err != nil {
		panic(err)
	}
}
