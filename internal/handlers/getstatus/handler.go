package getstatus

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
	log.Printf("retrieve status of game with id: %v", id)

	game := session.Game{
		ID: session.GameID(id),
	}

	apiGame := api.ToGameResponse(game)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}
