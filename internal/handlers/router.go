package handlers

import (
	"httpgordle/internal/api"
	"httpgordle/internal/handlers/getstatus"
	"httpgordle/internal/handlers/guess"
	"httpgordle/internal/handlers/newgame"
	"net/http"
)

// NewRouter returns a new router that listens for requests to the following endpoints:
// - Create a new game
// - Get the status of a game

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle)
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handle)
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handle)
	return r
}
