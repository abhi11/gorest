package main

import (
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Welcome,
	},
	Route{
		"AllLogs",
		"GET",
		"/logs",
		GetLogs,
	},
	Route{
		"LogsBefore",
		"GET",
		"/logs/before/{timestamp}",
		GetLogsBefore,
	},
	Route{
		"LogsAfter",
		"GET",
		"/logs/after/{timestamp}",
		GetLogsAfter,
	},
}