package guess

import (
	"encoding/json"
	"httpgordle/internal/api"
	"httpgordle/internal/session"
	"log"
	"net/http"
)

type gameGuesser interface {
	Find(ID session.GameID) (session.Game, error)
	Update(game session.Game) error
}

func Handler(guesser gameGuesser) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing GameID", http.StatusBadRequest)
			return
		}

		r := api.GuessRequest{}
		err := json.NewDecoder(req.Body).Decode(&r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		game := guess(id, r, guesser)

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

func guess(id string, r api.GuessRequest, db gameGuesser) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
