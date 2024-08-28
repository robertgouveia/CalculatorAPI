package middleware

import "net/http"

func RegisterMiddleware(mux *http.ServeMux) http.Handler {
	middlewares := [...]func(http.Handler) http.Handler{
		JsonParseMiddleware,
		CheckMethodMiddleware,
	}

	wrappedHandler := http.Handler(mux)

	for _, middlewareToAdd := range middlewares {
		wrappedHandler = middlewareToAdd(wrappedHandler)
	}

	return wrappedHandler
}
