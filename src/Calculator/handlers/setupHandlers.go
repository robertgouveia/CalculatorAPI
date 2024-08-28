package handlers

import "net/http"

func SetupHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	for path, handler := range map[string]http.HandlerFunc{
		"/add":      Add,
		"/subtract": Subtract,
		"/multiply": Multiply,
		"/divide":   Divide,
	} {
		mux.HandleFunc(path, handler)
	}

	return mux
}
