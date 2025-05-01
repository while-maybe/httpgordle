package getstatus

import (
	"encoding/json"
	"errors"
	"httpgordle/internal/api"
	"httpgordle/internal/repository"
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
			http.Error(w, "missing GameID", http.StatusNotFound)
			return
		}
		log.Printf("retrieve status of game with id: %v", id)

		game, err := finder.Find(session.GameID(id))
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				http.Error(w, "this game does not exist", http.StatusNotFound)
				return
			}

			log.Printf("cannot fetch game %s: %s", id, err)
			http.Error(w, "failed to fetch game", http.StatusInternalServerError)
			return
		}

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
			return
		}
	}
}
