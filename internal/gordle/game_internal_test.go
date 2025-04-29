package gordle

import (
	"errors"
	"slices"
	"testing"
)

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word string
		want error
	}{
		"nominal": {
			word: "GUESS",
			want: nil,
		},
		"too short": {
			word: "HI",
			want: ErrInvalidGuessLength,
		},
		"too long": {
			word: "SHOULDFAIL",
			want: ErrInvalidGuessLength,
		},
		"empty string": {
			word: "",
			want: ErrInvalidGuessLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			// only the validateGuess is being tested so New take nil/0 args
			// g := New(nil, "XXXXX", 0)
			g, _ := New("XXXXX")

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.want) {
				t.Errorf("%s, expected %q, got %q", tc.word, tc.want, err)
			}
		})
	}
}

func TestGameSplitToUppercaseCharacters(t *testing.T) {
	tt := map[string]struct {
		word string
		want []rune
	}{
		"lower": {
			word: "lower",
			want: []rune{'L', 'O', 'W', 'E', 'R'},
		},
		"Title": {
			word: "Title",
			want: []rune{'T', 'I', 'T', 'L', 'E'},
		},
		"mIxEd": {
			word: "mIxEd",
			want: []rune{'M', 'I', 'X', 'E', 'D'},
		},
		"CAPITALS": {
			word: "CAPITALS",
			want: []rune{'C', 'A', 'P', 'I', 'T', 'A', 'L', 'S'},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := splitToUppercaseCharacters(tc.word)

			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess    string
		solution string
		fb       Feedback
	}{
		"good guess": {
			solution: "hello",
			guess:    "hello",
			fb:       Feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"1 wrong char guess": {
			solution: "hello",
			guess:    "hlllo",
			fb:       Feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"2 wrong char guess": {
			solution: "hello",
			guess:    "shall",
			fb:       Feedback{absentCharacter, wrongPosition, absentCharacter, correctPosition, wrongPosition},
		},
		"3 wrong char guess": {
			solution: "hello",
			guess:    "shall",
			fb:       Feedback{absentCharacter, wrongPosition, absentCharacter, correctPosition, wrongPosition},
		},
		"4 wrong char guess": {
			solution: "hello",
			guess:    "hleol",
			fb:       Feedback{correctPosition, wrongPosition, wrongPosition, wrongPosition, wrongPosition},
		},
		"5 wrong char guess": {
			solution: "hello",
			guess:    "lloeh",
			fb:       Feedback{wrongPosition, wrongPosition, wrongPosition, wrongPosition, wrongPosition},
		},
		"no guess": {
			solution: "hello",
			guess:    "xxxxx",
			fb:       Feedback{absentCharacter, absentCharacter, absentCharacter, absentCharacter, absentCharacter},
		},
		"empty": {
			solution: "",
			guess:    "",
			fb:       Feedback{},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := computeFeedback([]rune(tc.guess), []rune(tc.solution))

			if !got.Equal(tc.fb) {
				t.Errorf("guess: %q, got the wrong feedback, expected %v, got %v", tc.guess, tc.fb, got)
			}
		})
	}
}
