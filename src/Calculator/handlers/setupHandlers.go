package handlers

import (
	"net/http"
)

type EndpointRequirements struct {
	Handler          http.HandlerFunc
	Required         []string
	CustomMiddleware []func(http.Handler) http.Handler
}

var Requirements = map[string]EndpointRequirements{
	"/add": {
		Handler:  Add,
		Required: []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{
			TestMiddleware,
		},
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
