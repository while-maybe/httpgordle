package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a game variable, which can be used to Play!
func New(reader io.Reader, corpus []string, maxAttempts int) (*Game, error) {

	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	g := &Game{
		reader:      bufio.NewReader(reader),
		solution:    []rune(strings.ToUpper(pickWord(corpus))), // picks a random word from the corpus
		maxAttempts: maxAttempts,
	}

	return g, nil
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {

		// ask for a valid word
		guess := g.ask()

		// give user some feedback on current attempt
		feedback := computeFeedback(guess, g.solution)
		fmt.Println(feedback.String())

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! you found it in %d guess(es)! The word was: %s\n", currentAttempt, string(g.solution))
			return
		}
	}
	fmt.Printf("😞 You've lost! The solution was: %s. \n", string(g.solution))
}

// ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-digit character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

// errInvalidWordLength is returned when the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string into a list of characters.
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// computeFeedback verifies every character of the guess against the solution.
func computeFeedback(guess, solution []rune) feedback {

	// initialise holders for marks
	result := make(feedback, len(guess))
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
