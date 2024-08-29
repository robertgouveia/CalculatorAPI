package middleware

import (
	"log/slog"
	"net/http"
	"os"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Info("API CALL: ", r.URL.Path+" ", r.Method)

		next.ServeHTTP(w, r)
	})
}
