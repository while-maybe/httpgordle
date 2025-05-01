package gordle

import (
	_ "embed"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

const (
	ErrEmptyCorpus        = corpusError("corpus is empty")
	ErrInaccessibleCorpus = corpusError("corpus can't be opened")
)

// ReadCorpus reads the file located at the given path and returns a list of words.
func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading (%s): %w", path, err, ErrInaccessibleCorpus)
	}

	// we expect the corpus to be a line or space-separated list of words
	words := strings.Fields(string(data))

	if len(words) == 0 {
		return nil, ErrEmptyCorpus
	}

	return words, nil
}

// pickWord returns a random word from the corpus
func pickRandomWord(corpus []string) string {
	index := rand.IntN(len(corpus))

	return corpus[index]
}
