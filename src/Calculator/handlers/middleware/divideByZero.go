package customMiddleware

import (
	"Calculator/helpers"
	"log/slog"
	"net/http"
	"os"
)

func DivideByZero(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Info("Divide by zero middleware")

		var jsonData, err = helpers.GetBody(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.Info("Divide operation",
			"Number1", jsonData["Number1"],
			"Number2", jsonData["Number2"],
		)

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
