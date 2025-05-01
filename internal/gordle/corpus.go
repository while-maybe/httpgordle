package gordle

import (
	_ "embed"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
	"sync"
)

const (
	ErrEmptyCorpus        = corpusError("corpus is empty")
	ErrInaccessibleCorpus = corpusError("corpus can't be opened")
)

type CorpusCache struct {
	wordMap map[string][]string
	mu      sync.RWMutex
}

// NewCorpusCache creates an empty mapBuffer
func NewCorpusCache() *CorpusCache {
	return &CorpusCache{
		wordMap: make(map[string][]string),
	}
}

// Set writes to CorpusCache using the provided path as key, and file contents as value in a map
func (cc *CorpusCache) Set(path string) ([]string, error) {
	log.Printf("Reading words from %s", path)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading (%s): %w", path, err, ErrInaccessibleCorpus)
	}

	words := strings.Fields(string(data))

	cc.mu.Lock()
	defer cc.mu.Unlock()

	cc.wordMap[path] = words

	return words, nil
}

// Get reads the CorpusCache for a given path and returns the stored contents if they exist or errors otherwise
func (cc *CorpusCache) Get(path string) ([]string, bool) {
	cc.mu.RLock()
	defer cc.mu.RUnlock()

	value, ok := cc.wordMap[path]

	if !ok {
		return nil, false
	}

	return value, true
}

// ReadCorpus reads the file located at the given path and returns a list of words.
func ReadCorpus(cc *CorpusCache, path string) ([]string, error) {
	words, ok := cc.Get(path)

	if ok {
		return words, nil
	}

	words, err := cc.Set(path)
	if err != nil {
		return nil, fmt.Errorf("Couldn't access %s: %w", path, err)
	}

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
