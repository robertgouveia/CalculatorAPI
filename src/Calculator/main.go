package main

import (
	"Calculator/handlers"
	"Calculator/middleware"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", handlers.Add)

	middlewares := [...]func(http.Handler) http.Handler{
		middleware.JsonParseMiddleware,
		middleware.CheckMethodMiddleware,
	}

	middlewareResult := http.Handler(mux)

	for _, middlewareToAdd := range middlewares {
		fmt.Println("Adding: " + runtime.FuncForPC(reflect.ValueOf(middlewareToAdd).Pointer()).Name())
		middlewareResult = middleware.AddMiddleware(middlewareResult, middlewareToAdd)
	}

	http.ListenAndServe(":8080", middlewareResult)
}
