package gordle_test

import (
	"httpgordle/internal/gordle"
	"testing"
)

const relativePath = "./../.."

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file   string
		length int
		err    error
	}{
		"English corpus": {
			file:   relativePath + "/corpus/english.txt",
			length: 34,
			err:    nil,
		},
		"empty corpus": {
			file:   relativePath + "/corpus/empty.txt",
			length: 0,
			err:    gordle.ErrEmptyCorpus,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			words, err := gordle.ReadCorpus(tc.file)
			if tc.err != err {
				t.Errorf("expected err %v, got%v", tc.err, err)
			}
			if tc.length != len(words) {
				t.Errorf("expected %d, got %d", tc.length, len(words))
			}
		})
	}
}
