package gordle

import (
	"fmt"
	"os"
	"strings"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	solution []rune
}

// New returns a game variable, which can be used to Play!
func New(solution string) (*Game, error) {

	if len(englishCorpus) == 0 {
		return nil, ErrEmptyCorpus
	}

	return &Game{
		solution: splitToUppercaseCharacters(solution),
	}, nil

}

// Play runs the game.
func (g *Game) Play(guess string) (Feedback, error) {
	err := g.validateGuess(guess)

	if err != nil {
		return Feedback{}, fmt.Errorf("this guess is not the correct length: %w", err)
	}

	characters := splitToUppercaseCharacters(guess)
	feedback := computeFeedback(characters, g.solution)

	return feedback, nil
}

// ErrInvalidGuessLength is returned when the guess has the wrong number of characters.
const ErrInvalidGuessLength = GameError("invalid guess length")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess string) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected solution with %d characters, got %d, %w", len(g.solution), len(guess), ErrInvalidGuessLength)
	}
	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string into a list of characters.
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// computeFeedback verifies every character of the guess against the solution.
func computeFeedback(guess, solution []rune) Feedback {

	// initialise holders for marks
	result := make(Feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution have different lengths: %d vs %d", len(guess), len(solution))
		return result
	}

	// check for correct letters
	for charPos := range guess {
		if guess[charPos] == solution[charPos] {
			result[charPos] = correctPosition
			used[charPos] = true
		}
	}

	// look for letters in the wrong position
	for charPos := range guess {
		// if the character at this position has already been used move on
		if result[charPos] != absentCharacter {
			continue
		}

		for charPosInSolution := range solution {
			// check if the pos of the solution char has already been used
			if used[charPosInSolution] {
				continue
			}

			// if the guess character matches the one in solution (but not in right position)
			if guess[charPos] == solution[charPosInSolution] {
				result[charPos] = wrongPosition
				used[charPosInSolution] = true
				break
			}
		}
	}

	return result
}
