package newgame

import (
	"encoding/json"
	"httpgordle/internal/api"
	"httpgordle/internal/session"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	game, err := CreateGame()
	if err != nil {
		log.Printf("unable to create a new game: %s", err)
		http.Error(w, "failed to create a new game", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	apiGame := response(game)

	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}

func CreateGame() (session.Game, error) {
	return session.Game{}, nil
}

func response(game session.Game) api.GameResponse {
	return api.GameResponse{}
}
