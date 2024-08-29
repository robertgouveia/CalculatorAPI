package handlers

import (
	"Calculator/structs"
	"net/http"
)

var Requirements = map[string]structs.EndpointRequirements{
	"/add": {
		Handler:          Add,
		Required:         []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
	"/subtract": {
		Handler:          Subtract,
		Required:         []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
	"/multiply": {
		Handler:          Multiply,
		Required:         []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
	"/divide": {
		Handler:          Divide,
		Required:         []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
	"/sum": {
		Handler:          Sum,
		Required:         []string{"Numbers"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
}

func GetEndpoints() []string {
	endpoints := make([]string, 0, len(Requirements))
	for key := range Requirements {
		endpoints = append(endpoints, key)
	}

	return endpoints
}

func GetRequirements(endpoint string) []string {
	return Requirements[endpoint].Required
}

func GetMiddleware(endpoint string) []func(http.Handler) http.Handler {
	return Requirements[endpoint].CustomMiddleware
}

func SetupHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	for path, handler := range Requirements {
		mux.HandleFunc(path, handler.Handler)
	}

	return mux
}
