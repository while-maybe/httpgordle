package handlers

import (
	"httpgordle/internal/api"
	"httpgordle/internal/handlers/newgame"
	"net/http"
)

// NewRouter returns a new router that listens for requests to the following endpoints:
// - Create a new game
// -
func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle)
	return r
}
