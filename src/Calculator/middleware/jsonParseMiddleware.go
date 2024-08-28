package middleware

import (
	"Calculator/handlers"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func JsonParseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var jsonData map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &jsonData); err != nil {
			http.Error(w, "Empty JSON Payload", http.StatusBadRequest)
			return
		}

		if handlers.GetRequirements(r.URL.Path) != nil {
			for _, field := range handlers.GetRequirements(r.URL.Path) {
				if jsonData[field] == nil {
					http.Error(w, "Missing field: "+field, http.StatusBadRequest)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
