package middleware

import (
	"encoding/json"
	"net/http"
)

func JsonParseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var jsonData map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
			http.Error(w, "Empty JSON Payload", http.StatusBadRequest)
			return
		}

		if jsonData["Number1"] == nil || jsonData["Number2"] == nil {
			http.Error(w, "Number1 and Number2 Required", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
