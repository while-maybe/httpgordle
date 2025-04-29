package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in English": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in Arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in Japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in Japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			// we're not testing the error so I'll ignore it here
			g, _ := New(strings.NewReader(tc.input), []string{string(tc.want)}, 0)
			// g := New(bufio.NewReader(os.Stdin), corpus, maxAttempts)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word []rune
		want error
	}{
		"nominal": {
			word: []rune("GUESS"),
			want: nil,
		},
		"too short": {
			word: []rune("HI"),
			want: ErrInvalidGuessLength,
		},
		"too long": {
			word: []rune("SHOULDFAIL"),
			want: ErrInvalidGuessLength,
		},
		"empty string": {
			word: []rune(""),
			want: ErrInvalidGuessLength,
		},
		"empty slice": {
			word: []rune{},
			want: ErrInvalidGuessLength,
		},
		"nil": {
			word: nil,
			want: ErrInvalidGuessLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			// only the validateGuess is being tested so New take nil/0 args
			// g := New(nil, "XXXXX", 0)
			g, _ := New(nil, []string{"XXXXX"}, 0)

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.want) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.want, err)
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
