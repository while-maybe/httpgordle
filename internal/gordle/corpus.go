package gordle

import (
	_ "embed"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
)

const (
	ErrEmptyCorpus        = corpusError("corpus is empty")
	ErrInaccessibleCorpus = corpusError("corpus can't be opened")
)

var wordMap = make(map[string][]string)

// ReadCorpus reads the file located at the given path and returns a list of words.
func ReadCorpus(path string) ([]string, error) {
	if wordMap[path] != nil {
		return wordMap[path], nil
	}

	log.Printf("Reading words from %s", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading (%s): %w", path, err, ErrInaccessibleCorpus)
	}

	// we expect the corpus to be a line or space-separated list of words
	wordMap[path] = strings.Fields(string(data))
	// words = strings.Fields(string(data))

	if len(wordMap[path]) == 0 {
		return nil, ErrEmptyCorpus
	}

	return wordMap[path], nil
}

// pickWord returns a random word from the corpus
func pickRandomWord(corpus []string) string {
	index := rand.IntN(len(corpus))

	return corpus[index]
}
