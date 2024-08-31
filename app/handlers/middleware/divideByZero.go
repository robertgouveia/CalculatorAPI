package customMiddleware

import (
	"Calculator/helpers"
	"net/http"
)

func DivideByZero(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var jsonData, err = helpers.GetBody(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		num1, ok1 := jsonData["Number1"].(float64)
		num2, ok2 := jsonData["Number2"].(float64)

		if !ok1 || !ok2 {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if num1 == 0 || num2 == 0 {
			http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
