package gordle

import "testing"

func TestFeedbackString(t *testing.T) {
	tt := map[string]struct {
		feedback feedback
		want     string
	}{
		"one hint absentCharacter": {
			feedback: feedback{absentCharacter},
			want:     "⬜️",
		},
		"one hint wrongPosition": {
			feedback: feedback{wrongPosition},
			want:     "🟡",
		},
		"one hint correctPosition": {
			feedback: feedback{correctPosition},
			want:     "💚",
		},
		"shouldn't be here": {
			feedback: feedback{101},
			want:     "💔",
		},
		"two hints": {
			feedback: feedback{wrongPosition, correctPosition},
			want:     "🟡💚",
		},
		"three hints": {
			feedback: feedback{correctPosition, absentCharacter, wrongPosition},
			want:     "💚⬜️🟡",
		},
		"four hints": {
			feedback: feedback{correctPosition, absentCharacter, correctPosition, wrongPosition},
			want:     "💚⬜️💚🟡",
		},
		"five hints": {
			feedback: feedback{100, wrongPosition, 120, wrongPosition, correctPosition},
			want:     "💔🟡💔🟡💚",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {

			if got := tc.feedback.String(); got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
