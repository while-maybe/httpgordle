package gordle

import (
	"slices"
	"strings"
)

type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️" // grey square
	case wrongPosition:
		return "🟡" // yellow circle
	case correctPosition:
		return "💚" // green heart
	default:
		// we shouldn't be here
		return "💔" // red broken heart
	}
}

// feedback is a list of hints, one per character of the word.
type Feedback []hint

func (fb Feedback) Equal(other Feedback) bool {
	if len(fb) != len(other) || !slices.Equal(fb, other) {
		return false
	}

	return true
}

// StringConcat is a naive implementation to build feedback as a string.
// It is used only to benchmark it against the strings.Builder version.
func (fb Feedback) StringConcat() string {
	var output string
	for _, h := range fb {
		output += h.String()
	}
	return output
}

// String implements the Stringer interface for a slice of hints.
func (fb Feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

// GameWon returns true if the player won the game
func (fb Feedback) GameWon() bool {
	for _, c := range fb {
		if c != correctPosition {
			return false
		}
	}
	return true
}
