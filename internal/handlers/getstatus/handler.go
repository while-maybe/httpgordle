package getstatus

import (
	"encoding/json"
	"httpgordle/internal/api"
	"httpgordle/internal/session"
	"log"
	"net/http"
)

type gameFinder interface {
	Find(ID session.GameID) (session.Game, error)
}

func Handler(finder gameFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing GameID", http.StatusBadRequest)
			return
		}
		log.Printf("retrieve status of game with id: %v", id)

		game := getGame(id, finder)

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

func getGame(id string, db gameFinder) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
