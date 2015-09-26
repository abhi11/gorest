package main

import (
	"net/http"
	"encoding/json"
)

func SetContentType(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	return w
}

func SetReturnCode(w http.ResponseWriter, status int) http.ResponseWriter {
	w.WriteHeader(status)
	return w
}

func EncodeResponse(w http.ResponseWriter, data interface{}) error {
	err := json.NewEncoder(w).Encode(data)
	return err
}

/*
func UnmarshalBody(body []byte, node *interface{}) error {
	err := json.Unmarshal(body, node)
	return err
}*/
