package main

import (
	"httpgordle/internal/handlers"
	"httpgordle/internal/repository"
	"net/http"
)

func main() {
	db := repository.New()

	err := http.ListenAndServe(":8123", handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}
