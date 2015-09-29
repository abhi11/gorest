package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/abhi11/gorest/util"
	"github.com/abhi11/gorest/mongo"
	"github.com/abhi11/gorest/model"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	var logs model.LogMessages
	caps := map[string]int{}

	query, err := util.MakeQuery(r)
	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}

	limit, err := util.LimitCount(r)

	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		return
	}

	offset, err := util.OffsetCount(r)

	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		return
	}

	caps["limit"] = limit
	caps["offset"] = offset

	logs = mongo.DBGetLogs(query, caps)


	w = util.SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := util.EncodeResponse(w , logs); err != nil {
		panic(err)
	}
}

func GetLogsForTwistDevice(w http.ResponseWriter, r *http.Request) {
	var logs model.LogMessages
	caps := map[string]int{}

	query, err := util.MakeQuery(r)

	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}


	limit, err := util.LimitCount(r)

	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		return
	}

	offset, err := util.OffsetCount(r)

	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		return
	}

	caps["limit"] = limit
	caps["offset"] = offset

	vars := util.URLVarsFromRequest(r)
	query = util.AddTwistIdToQuery(vars, query)

	logs = mongo.DBGetLogs(query, caps)

	w = util.SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := util.EncodeResponse(w , logs); err != nil {
		panic(err)
	}
}

func GetLogsForMobileDevice(w http.ResponseWriter, r *http.Request) {
	var logs model.LogMessages
	caps := map[string]int{}

	query, err := util.MakeQuery(r)

	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}

	limit, err := util.LimitCount(r)

	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		return
	}

	offset, err := util.OffsetCount(r)

	if err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType)
		return
	}

	caps["limit"] = limit
	caps["offset"] = offset


	vars := util.URLVarsFromRequest(r)
	query = util.AddMobileIdToQuery(vars, query)

	logs = mongo.DBGetLogs(query, caps)

	w = util.SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := util.EncodeResponse(w , logs); err != nil {
		panic(err)
	}
}

func PostLog(w http.ResponseWriter, r *http.Request) {
	var logEntry model.LogMessage
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &logEntry); err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType) // cannot process
		err := json.NewEncoder(w).Encode(err)

		if err != nil {
			panic(err)
		}
		return
	}

	res := mongo.DBPostLog(logEntry)
	if res == 1 { // Error while inserting
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}

	w = util.SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := util.EncodeResponse(w, logEntry); err != nil {
		panic(err)
	}
}

func PostLogsBatch(w http.ResponseWriter, r *http.Request) {
	var logEntries model.LogMessages
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 20971520))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &logEntries); err != nil {
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusUnsupportedMediaType) // cannot process

		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			panic(err)
		}
		return
	}

	res := mongo.DBPostLogsBatch(logEntries)
	if res == 1 { // Error while inserting
		w = util.SetContentTypeAndReturnCode(w,
			http.StatusInternalServerError)
		return
	}

	w = util.SetContentTypeAndReturnCode(w, http.StatusOK)
	if err := util.EncodeResponse(w, logEntries); err != nil {
		panic(err)
	}
}
