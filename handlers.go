package main

import (
	"io"
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


func Welcome(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	w = SetReturnCode(w, http.StatusOK)

	fmt.Fprintf(w, "Hi there")
}

func GetLogs(w http.ResponseWriter, r *http.Request) {
	logs := DBGetAllLogs()

	w = SetContentType(w)
	w = SetReturnCode(w, http.StatusOK)

	if err := EncodeResponse(w , logs); err != nil {
		panic(err)
	}
}

func GetLogsWithTimestamp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkpoint := vars["timestamp"]
	timestamp, parserr := strconv.ParseInt(checkpoint, 10, 64)

	if parserr != nil {
		panic(parserr)
	}

	logs := DBGetLogs(timestamp)

	w = SetContentType(w)
	w = SetReturnCode(w, http.StatusOK)

	if err := EncodeResponse(w , logs); err != nil {
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

	logs := DBGetLogsAfter(timestamp)

	w = SetContentType(w)
	w = SetReturnCode(w, http.StatusOK)

	if err := EncodeResponse(w , logs); err != nil {
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

	logs := DBGetLogsBefore(timestamp)

	w = SetContentType(w)
	w = SetReturnCode(w, http.StatusOK)

	if err := EncodeResponse(w , logs); err != nil {
		panic(err)
	}
}

func PostLog(w http.ResponseWriter, r *http.Request) {
	var logEntry LogMessage
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &logEntry); err != nil {
		w = SetContentType(w)
		w = SetReturnCode(w, http.StatusUnsupportedMediaType) // cannot process
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	res := DBPostLog(logEntry)
	if res == 1 { // Error while inserting
		w = SetContentType(w)
		w = SetReturnCode(w, http.StatusInternalServerError)
		return
	}

	w = SetContentType(w)
	w = SetReturnCode(w, http.StatusOK)

	if err := EncodeResponse(w, logEntry); err != nil {
		panic(err)
	}
}

func PostLogsInBath(w http.ResponseWriter, r *http.Request) {
	var logEntries LogMessages
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 20971520))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &logEntries); err != nil {
		w = SetContentType(w)
		w = SetReturnCode(w, http.StatusUnsupportedMediaType) // cannot process

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	res := DBPostLogsBatch(logEntries)
	if res == 1 { // Error while inserting
		w = SetContentType(w)
		w = SetReturnCode(w, http.StatusInternalServerError)
		return
	}

	w = SetContentType(w)
	w = SetReturnCode(w, http.StatusOK)

	if err := EncodeResponse(w, logEntries); err != nil {
		panic(err)
	}
}
