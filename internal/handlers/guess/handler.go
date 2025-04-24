package guess

import (
	"encoding/json"
	"httpgordle/internal/api"
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

	apiGame := api.GameResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}
