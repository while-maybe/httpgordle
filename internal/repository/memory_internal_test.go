package repository

import (
	"errors"
	"httpgordle/internal/session"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var startingGamesRepo = &GameRepository{
	storage: map[session.GameID]session.Game{
		session.GameID("10"): {
			ID:           session.GameID("10"),
			AttemptsLeft: 5,
			Guesses:      []session.Guess{},
			Status:       session.StatusPlaying,
		},
	},
}

func TestNew(t *testing.T) {
	want := &GameRepository{
		storage: map[session.GameID]session.Game{},
	}

	dontWant := startingGamesRepo

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
			want: startingGamesRepo,
			err:  nil,
		},
		"Existing Element": {
			starting: startingGamesRepo,
			game: session.Game{
				ID:           session.GameID("10"),
				AttemptsLeft: 5,
				Guesses:      []session.Guess{},
				Status:       session.StatusPlaying,
			},
			want: startingGamesRepo,
			err:  ErrConflictingID,
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
			existing: startingGamesRepo,
			ID:       session.GameID("10"),
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

func TestUpdate(t *testing.T) {
	tt := map[string]struct {
		existing, want *GameRepository
		game           session.Game
		err            error
	}{
		"update existing": {
			existing: startingGamesRepo,
			game: session.Game{
				ID:           session.GameID("10"),
				AttemptsLeft: 2,
				Guesses:      []session.Guess{},
				Status:       session.StatusPlaying,
			},
			want: &GameRepository{
				storage: map[session.GameID]session.Game{
					session.GameID("10"): {
						ID:           session.GameID("10"),
						AttemptsLeft: 2,
						Guesses:      []session.Guess{},
						Status:       session.StatusPlaying,
					},
				},
			},
			err: nil,
		},
		"update missing": {
			existing: startingGamesRepo,
			game: session.Game{
				ID:           session.GameID("20"),
				AttemptsLeft: 2,
				Guesses:      []session.Guess{},
				Status:       session.StatusPlaying,
			},
			want: startingGamesRepo,
			err:  ErrNotFound,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.existing
			err := got.Update(tc.game)

			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}
