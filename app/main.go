package main

import (
	"Calculator/handlers"
	"Calculator/middleware"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", middleware.RegisterMiddleware(handlers.SetupHandlers()))
}
