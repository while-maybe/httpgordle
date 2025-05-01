package handlers

import (
	"httpgordle/internal/api"
	"httpgordle/internal/handlers/getstatus"
	"httpgordle/internal/handlers/guess"
	"httpgordle/internal/handlers/newgame"
	"httpgordle/internal/repository"
	"net/http"
)

// NewRouter returns a new router that listens for requests to the following endpoints:
// - Create a new game
// - Get the status of a game
// - Make a guess in the game

const corpusPath = "corpus/english.txt"

func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handler(db, corpusPath))
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handler(db))
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handler(db))
	return r
}
