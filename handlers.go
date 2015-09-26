package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


func Welcome(w http.ResponseWriter, r *http.Request) {
	w = SetHeaders(w)
	fmt.Fprintf(w, "Hi there")
}

func GetLogs(w http.ResponseWriter, r *http.Request) {
	logs := LogMessages{
		LogMessage{
			Name: "After",
			Severity: "Debug",
			Timestamp: time.Now().Unix(),
		},
		LogMessage{
			Name: "Before",
			Severity : "Debug",
			Timestamp: time.Now().Unix(),
		},
	}
	w = SetHeaders(w)
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		panic(err)
	}
}

func GetLogsAfter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkpoint := vars["timestamp"]
	timestamp, parserr := strconv.ParseInt(checkpoint, 10, 64)

	if parserr != nil {
		panic(parserr)
	}

	logs := LogMessages{
		LogMessage{
			Name: "After",
			Severity: "Debug",
			Timestamp: timestamp,
		},
	}

	fmt.Printf("Logs after %s", checkpoint)

	SetHeaders(w)
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		panic(err)
	}
}

func GetLogsBefore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkpoint := vars["timestamp"]
	timestamp, parserr := strconv.ParseInt(checkpoint, 10, 64)
	if parserr != nil {
		panic(parserr)
	}

	logs := LogMessages{
		LogMessage{
			Name: "Before",
			Severity : "Debug",
			Timestamp: timestamp,
		},
	}

	fmt.Printf("Logs before %s", checkpoint)

	w = SetHeaders(w)
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		panic(err)
	}
}

func SetHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return w
}
