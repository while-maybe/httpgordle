package session

import (
	"errors"
	"httpgordle/internal/gordle"
)

// Game contains the information about a game.
type Game struct {
	ID           GameID
	Gordle       gordle.Game
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}

// GameID represents the ID of a game
type GameID string

// Status is the current status of the game and tells what operations can be made on it
type Status string

const (
	StatusPlaying = "Playing"
	StatusWon     = "Won"
	StatusLost    = "Lost"
)

// A Guess is a (user-submitted) word and (Gordle generated) feedback pair.
type Guess struct {
	Word     string
	Feedback string
}

// ErrGameOver is returned whan a play is made but the game is already over.
var ErrGameOver = errors.New("game over")
