package middleware

import (
	"Calculator/handlers"
	"net/http"
)

func CustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware := handlers.GetMiddleware(r.URL.Path)

		if middleware == nil {
			return
		}

		for _, m := range middleware {
			next = m(next)
		}

		next.ServeHTTP(w, r)
	})
}
