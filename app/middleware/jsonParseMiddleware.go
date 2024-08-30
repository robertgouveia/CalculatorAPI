package middleware

import (
	"Calculator/handlers"
	"Calculator/helpers"
	"net/http"
)

func JsonParseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if handlers.GetRequirements(r.URL.Path) != nil {
			for _, field := range handlers.GetRequirements(r.URL.Path) {

				var jsonData, err = helpers.GetBody(r)

				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				if jsonData[field] == nil {
					http.Error(w, "Missing field: "+field, http.StatusBadRequest)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
