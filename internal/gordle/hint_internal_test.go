package gordle

import "testing"

func TestFeedbackString(t *testing.T) {
	tt := map[string]struct {
		feedback Feedback
		want     string
	}{
		"one hint absentCharacter": {
			feedback: Feedback{absentCharacter},
			want:     "⬜️",
		},
		"one hint wrongPosition": {
			feedback: Feedback{wrongPosition},
			want:     "🟡",
		},
		"one hint correctPosition": {
			feedback: Feedback{correctPosition},
			want:     "💚",
		},
		"shouldn't be here": {
			feedback: Feedback{101},
			want:     "💔",
		},
		"two hints": {
			feedback: Feedback{wrongPosition, correctPosition},
			want:     "🟡💚",
		},
		"three hints": {
			feedback: Feedback{correctPosition, absentCharacter, wrongPosition},
			want:     "💚⬜️🟡",
		},
		"four hints": {
			feedback: Feedback{correctPosition, absentCharacter, correctPosition, wrongPosition},
			want:     "💚⬜️💚🟡",
		},
		"five hints": {
			feedback: Feedback{100, wrongPosition, 120, wrongPosition, correctPosition},
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
