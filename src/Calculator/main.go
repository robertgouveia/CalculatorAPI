package main

import (
	"Calculator/handlers"
	"Calculator/middleware"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", middleware.RegisterMiddleware(handlers.SetupHandlers()))
}
