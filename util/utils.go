package util

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
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
	var after_timestamp int64
	var before_timestamp int64

	after := r.FormValue("after")
	before := r.FormValue("before")
	log_level := r.FormValue("log_level")

	query := bson.M{}
	time_query := bson.M{}
	after_flag := false
	before_flag := false
	no_filters := true

	if after != "" {
		parsed_time, err := strconv.ParseInt(after, 10, 64)
		if err != nil {
			return nil, err
		}
		time_query["$gte"] = parsed_time
		after_timestamp = parsed_time
		after_flag = true
		no_filters = false
	}

	if before != "" {
		parsed_time, err := strconv.ParseInt(before, 10, 64)
		if err != nil {
			return nil, err
		}
		time_query["$lt"] = parsed_time
		before_timestamp = parsed_time
		before_flag = true
		no_filters = false
	}

	// check if agfter > before, and return with err
	if after_flag && before_flag && after_timestamp > before_timestamp {
		return nil, errors.New("After timestamp is more than before")
	}

	if log_level != "" {
		query["loglevel"] = log_level
		no_filters = false
	}

	if before_flag || after_flag { // time flag is present
		query["timestamp"] = time_query
	}

	if no_filters {
		return nil, nil
	}

	return query, nil
}

func IntValFromRequest(r *http.Request, key string) (int, error) {
	val := r.FormValue(key)

	if val == "" {
		return 0, nil
	}

	intval, err := strconv.Atoi(val)

	if err != nil {
		return 0, err
	}

	return intval, nil
}

func LimitCount(r *http.Request) (int, error) {
	limit, err := IntValFromRequest(r, "limit")

	return limit, err
}

func OffsetCount(r *http.Request) (int, error) {
	offset, err := IntValFromRequest(r, "offset")

	return offset, err
}

func AddValueToBSONMap(srcmap map[string]string, destmap bson.M,
	srckey string, destkey string) bson.M {

	if destmap == nil {
		destmap = bson.M{}
	}
	destmap[destkey] = srcmap[srckey]

	return destmap
}

func AddTwistIdToQuery(vars map[string]string, query bson.M) bson.M {
	query = AddValueToBSONMap(vars, query, "id", "twistdeviceid")

	return query
}

func AddMobileIdToQuery(vars map[string]string, query bson.M) bson.M {
	query = AddValueToBSONMap(vars, query, "id", "mobiledeviceid")

	return query
}

func URLVarsFromRequest(r *http.Request) map[string]string {
	vars := mux.Vars(r)

	return vars
}
