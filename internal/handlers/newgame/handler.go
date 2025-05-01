package newgame

import (
	"encoding/json"
	"fmt"
	"httpgordle/internal/api"
	"httpgordle/internal/gordle"
	"httpgordle/internal/session"
	"log"
	"net/http"

	"github.com/oklog/ulid/v2"
)

type gameAdder interface {
	Add(game session.Game) error
}

// Handler returns the handler for the game creation endpoint.
func Handler(adder gameAdder, corpusPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {

		game, err := createGame(adder, corpusPath)

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

const maxAttempts = 5

func createGame(db gameAdder, corpusPath string) (session.Game, error) {
	corpus, err := gordle.ReadCorpus(corpusPath) // should come from config
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to read corpus: %w", err)
	}

	if len(corpus) == 0 {
		return session.Game{}, gordle.ErrEmptyCorpus
	}

	game, err := gordle.New(corpus)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to create a new Gordle game")
	}

	g := session.Game{
		ID:           session.GameID(ulid.Make().String()),
		Gordle:       *game,
		AttemptsLeft: maxAttempts,
		Guesses:      []session.Guess{},
		Status:       session.StatusPlaying,
	}

	err = db.Add(g)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to save the new game")
	}

	return g, nil
}
