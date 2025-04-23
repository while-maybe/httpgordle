package api

const (
	// NewGameRoute is the path to create a new game.
	NewGameRoute = "/games"
)

// GameResponse constains the information about a game.
type GameResponse struct {
	ID           string  `json:"id"`
	AttemptsLeft int     `json:"attempts_left"`
	Guesses      []Guess `json:"guesses"`
	WordLength   byte    `json:"word_length"`
	Solution     string  `json:"solution,omitempty"`
	Status       string  `json:"status"`
}

// Guess is a pair of a word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}
