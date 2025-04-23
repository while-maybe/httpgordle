package main

import (
	"httpgordle/internal/handlers"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8123", handlers.NewRouter())
	if err != nil {
		panic(err)
	}
}
