package getstatus

import (
	"httpgordle/internal/api"
	"httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type gameFinderStub struct {
	game session.Game
	err  error
}

func (g gameFinderStub) Find(ID session.GameID) (session.Game, error) {
	return g.game, g.err
}

func TestHandle(t *testing.T) {
	handleFunc := Handler(gameFinderStub{session.Game{ID: "123456"}, nil})

	req, err := http.NewRequest(http.MethodGet, "/games/", nil)
	require.NoError(t, err)

	// add a path parameter
	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()

	handleFunc(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}
