package main

import (
	"net/http"
	"strconv"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

func SetContentTypeAndReturnCode(w http.ResponseWriter, status int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	return w
}

func EncodeResponse(w http.ResponseWriter, data interface{}) error {
	err := json.NewEncoder(w).Encode(data)
	return err
}

func MakeQuery(r *http.Request) (bson.M, error) {
	after := r.FormValue("after")
	before := r.FormValue("before")
	log_level := r.FormValue("log_level")
	query := bson.M{}
	time_query := bson.M{}
	time_flag := false
	no_filters := true

	if after != "" {
		after_timestamp, err := strconv.ParseInt(after, 10, 64)
		if err != nil {
			return nil, err
		}
		time_query["$gte"] = after_timestamp
		time_flag = true
		no_filters = false
	}

	if before != "" {
		before_timestamp, err := strconv.ParseInt(before, 10, 64)
		if err != nil {
			return nil, err
		}
		time_query["$lt"] = before_timestamp
		time_flag = true
		no_filters = false
	}

	if log_level != "" {
		query["loglevel"] = log_level
		no_filters = false
	}

	if time_flag {
		query["timestamp"] = time_query
	}

	if no_filters {
		return nil, nil
	}

	return query, nil
}

func GetLimitCount(r *http.Request) (int, error) {
	limit := r.FormValue("limit")

	if limit == "" {
		return -1, nil
	}

	count, err := strconv.Atoi(limit)

	if err != nil {
		return -1, err
	}

	return count, nil
}
