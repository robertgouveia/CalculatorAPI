package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func JsonParseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling: JsonParseMiddleware")

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Return the io.ReadCloser to its original state

		var jsonData map[string]interface{} // Define a map to store the JSON data
		if err := json.Unmarshal(bodyBytes, &jsonData); err != nil {
			http.Error(w, "Empty JSON Payload", http.StatusBadRequest)
			return
		}

		if jsonData["Number1"] == nil || jsonData["Number2"] == nil {
			http.Error(w, "No Number1, Number2 supplied", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
