package main

import (
	"Calculator/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/add", handlers.Add)

	http.ListenAndServe(":8080", nil)
}
