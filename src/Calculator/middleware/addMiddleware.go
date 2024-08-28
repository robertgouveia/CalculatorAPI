package middleware

import "net/http"

func AddMiddleware(middleware http.Handler, additional func(http.Handler) http.Handler) http.Handler {
	return additional(middleware)
}
