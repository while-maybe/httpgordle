package repository

import (
	"fmt"
	"httpgordle/internal/session"
	"log"
)

// GameRepository holds all the current games.
type GameRepository struct {
	storage map[session.GameID]session.Game
}

// New creates an empty game repository.
func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

// Add inserts for the first time a game in memory.
func (gr *GameRepository) Add(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("%w (%s)", ErrConflictingID, game.ID)
	}

	gr.storage[game.ID] = game

	return nil
}

// Find returns an existing game given its ID, errors if it doesn't exist
func (gr *GameRepository) Find(ID session.GameID) (session.Game, error) {
	log.Printf("Looking for game %s...", ID)

	game, ok := gr.storage[ID]
	if !ok {
		return session.Game{}, fmt.Errorf("Can't find game %s: %w", ID, ErrNotFound)
	}

	return game, nil
}

// Update modifies an existing game, errors if it doesn't exist
func (gr *GameRepository) Update(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if !ok {
		return fmt.Errorf("Can't find game %s: %w", game.ID, ErrNotFound)
	}

	gr.storage[game.ID] = game

	return nil
}
