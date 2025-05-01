package api

import "httpgordle/internal/session"

// ToGameResponse converts a session.Game into a GameResponse.
func ToGameResponse(g session.Game) GameResponse {
	solution := g.Gordle.ShowAnswer()

	apiGame := GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]Guess, len(g.Guesses)),
		Status:       string(g.Status),
		WordLength:   byte(len(solution)),
	}

	for i := range g.Guesses {
		apiGame.Guesses[i].Word = g.Guesses[i].Word
		apiGame.Guesses[i].Feedback = g.Guesses[i].Feedback
	}

	if g.AttemptsLeft == 0 {
		apiGame.Solution = solution
	}

	return apiGame
}
