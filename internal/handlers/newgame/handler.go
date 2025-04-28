package newgame

import (
	"encoding/json"
	"httpgordle/internal/api"
	"httpgordle/internal/session"
	"log"
	"net/http"
)

type gameAdder interface {
	Add(game session.Game) error
}

// Handler returns the handler for the game creation endpoint.
func Handler(adder gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		game, err := CreateGame(adder)

		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(w, "failed to create a new game", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := api.ToGameResponse(game)

		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

func CreateGame(adder gameAdder) (session.Game, error) {
	return session.Game{}, nil
}
