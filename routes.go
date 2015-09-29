package main

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"AllLogs",
		"GET",
		"/logs",
		GetLogs,
	},
	Route{
		"LogsForTwistDevice",
		"GET",
		"/logs/twist_device/{id}",
		GetLogsForTwistDevice,
	},
	Route{
		"LogsForMobileDevice",
		"GET",
		"/logs/mobile_device/{id}",
		GetLogsForMobileDevice,
	},
	Route{
		"PostLog",
		"POST",
		"/logs",
		PostLog,
	},
	Route{
		"PostMultipleLogs",
		"POST",
		"/logs/batch_insert",
		PostLogsBatch,
	},
}
