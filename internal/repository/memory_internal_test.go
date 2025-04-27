package repository

import (
	"errors"
	"httpgordle/internal/session"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	want := &GameRepository{
		storage: map[session.GameID]session.Game{},
	}

	dontWant := &GameRepository{
		storage: map[session.GameID]session.Game{
			session.GameID("10"): {
				ID:           session.GameID("10"),
				AttemptsLeft: 5,
				Guesses:      []session.Guess{},
				Status:       session.StatusPlaying,
			},
		},
	}

	got := New()

	assert.Equal(t, want, got)
	assert.NotEqual(t, dontWant, got)
}

func TestAdd(t *testing.T) {
	tt := map[string]struct {
		starting, want *GameRepository
		game           session.Game
		err            error
	}{
		"New Element": {
			starting: New(),
			game: session.Game{
				ID:           session.GameID("10"),
				AttemptsLeft: 5,
				Guesses:      []session.Guess{},
				Status:       session.StatusPlaying,
			},
			want: &GameRepository{
				storage: map[session.GameID]session.Game{
					session.GameID("10"): {
						ID:           session.GameID("10"),
						AttemptsLeft: 5,
						Guesses:      []session.Guess{},
						Status:       session.StatusPlaying,
					},
				},
			},
			err: nil,
		},
		"Existing Element": {
			starting: &GameRepository{
				storage: map[session.GameID]session.Game{
					session.GameID("10"): {
						ID:           session.GameID("10"),
						AttemptsLeft: 5,
						Guesses:      []session.Guess{},
						Status:       session.StatusPlaying,
					},
				},
			},
			game: session.Game{
				ID:           session.GameID("10"),
				AttemptsLeft: 5,
				Guesses:      []session.Guess{},
				Status:       session.StatusPlaying,
			},
			want: &GameRepository{
				storage: map[session.GameID]session.Game{
					session.GameID("10"): {
						ID:           session.GameID("10"),
						AttemptsLeft: 5,
						Guesses:      []session.Guess{},
						Status:       session.StatusPlaying,
					},
				},
			},
			err: ErrConflictingID,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {

			got := tc.starting
			err := got.Add(tc.game)

			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tt := map[string]struct {
		existing *GameRepository
		ID       session.GameID
		want     session.Game
		err      error
	}{
		"finds existing": {
			existing: &GameRepository{
				storage: map[session.GameID]session.Game{
					session.GameID("10"): {
						ID:           session.GameID("10"),
						AttemptsLeft: 5,
						Guesses:      []session.Guess{},
						Status:       session.StatusPlaying,
					},
				},
			},
			ID: session.GameID("10"),
			want: session.Game{
				ID:           session.GameID("10"),
				AttemptsLeft: 5,
				Guesses:      []session.Guess{},
				Status:       session.StatusPlaying,
			},
			err: nil,
		},

		"finds nothing": {
			existing: &GameRepository{
				storage: map[session.GameID]session.Game{
					session.GameID("10"): {
						ID:           session.GameID("10"),
						AttemptsLeft: 5,
						Guesses:      []session.Guess{},
						Status:       session.StatusPlaying,
					},
				},
			},
			ID:   session.GameID("15"),
			want: session.Game{},
			err:  ErrNotFound,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := tc.existing.Find(tc.ID)

			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, got)
			}

		})
	}
}
