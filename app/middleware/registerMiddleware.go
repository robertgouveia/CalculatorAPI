package middleware

import (
	"net/http"
)

func RegisterMiddleware(mux *http.ServeMux) http.Handler {
	middlewares := [...]func(http.Handler) http.Handler{
		Logger,
		JsonParseMiddleware,
		CheckMethodMiddleware,
	}

	wrappedHandler := http.Handler(mux)

	for _, middlewareToAdd := range middlewares {
		wrappedHandler = middlewareToAdd(wrappedHandler)
	}

	return CustomMiddleware(wrappedHandler)
}
