package guess

import (
	"encoding/json"
	"httpgordle/internal/api"
	"httpgordle/internal/session"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
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

	game := guess(id, r)

	apiGame := api.ToGameResponse(game)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}

func guess(id string, r api.GuessRequest) session.Game {
	return session.Game{
		ID: session.GameID(id),
	}
}
