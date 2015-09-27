package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	var logs LogMessages

	query, err := MakeQuery(r)
	if err != nil {
		w = SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		panic(err)
		return
	}

	count, err := GetLimitCount(r)

	if err != nil {
		w = SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		panic(err)
		return
	}

	if count < 0 {
		logs = DBGetLogs(query, -1)
	} else {
		logs = DBGetLogs(query, count)
	}

	w = SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := EncodeResponse(w , logs); err != nil {
		panic(err)
	}
}

func GetLogsForTwistDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var logs LogMessages

	query, err := MakeQuery(r)

	if err != nil {
		w = SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}

	count, err := GetLimitCount(r)

	if err != nil {
		w = SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		return
	}

	if query == nil {
		query = bson.M{}
	}

	query["twistdeviceid"] =  vars["id"]

	if count < 0 {
		logs = DBGetLogs(query, -1)
	} else {
		logs = DBGetLogs(query, count)
	}

	w = SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := EncodeResponse(w , logs); err != nil {
		panic(err)
	}
}

func GetLogsForMobileDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var logs LogMessages

	query, err := MakeQuery(r)

	if err != nil {
		w = SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}

	count, err := GetLimitCount(r)

	if err != nil {
		w = SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		return
	}

	if query == nil {
		query = bson.M{}
	}
	query["mobiledeviceid"] = vars["id"]

	if count < 0 {
		logs = DBGetLogs(query, -1)
	} else {
		logs = DBGetLogs(query, count)
	}

	w = SetContentTypeAndReturnCode(w, http.StatusOK)
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
		w = SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType) // cannot process
		err := json.NewEncoder(w).Encode(err)

		if err != nil {
			panic(err)
		}
		return
	}

	res := DBPostLog(logEntry)
	if res == 1 { // Error while inserting
		w = SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}

	w = SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := EncodeResponse(w, logEntry); err != nil {
		panic(err)
	}
}

func PostLogsBatch(w http.ResponseWriter, r *http.Request) {
	var logEntries LogMessages
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 20971520))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &logEntries); err != nil {
		w = SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType) // cannot process

		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			panic(err)
		}
		return
	}

	res := DBPostLogsBatch(logEntries)
	if res == 1 { // Error while inserting
		w = SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}

	w = SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := EncodeResponse(w, logEntries); err != nil {
		panic(err)
	}
}
