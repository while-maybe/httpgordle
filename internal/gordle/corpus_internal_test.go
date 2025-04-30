package gordle

import "testing"

// inCorpus checks if a word is contained in a provided word slice returning true/false
func inCorpus(corpus []string, word string) bool {
	// we could have simply used slices.Contains() but no harm done in extra practice
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}
	return false
}

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}
	word := PickRandomWord(corpus)

	if !inCorpus(corpus, word) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}
