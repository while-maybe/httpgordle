package newgame

import (
	"httpgordle/internal/api"
	"httpgordle/internal/gordle"
	"httpgordle/internal/session"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type gameCreatorStub struct {
	err error
}

func (g gameCreatorStub) Add(_ session.Game) error {
	return g.err
}

func TestHandler(t *testing.T) {
	idFinderRegExp := regexp.MustCompile(`.+"id":"([0-9A-Z]+)".+`)
	corpusPath := "./../../../corpus/english.txt"

	tt := map[string]struct {
		wantStatusCode int
		wantBody       string
		creator        gameAdder
	}{
		"nominal": {
			wantStatusCode: http.StatusCreated,
			wantBody:       `{"id":"123456","attempts_left":5,"guesses":[],"word_length":4,"status":"Playing"}`,
			creator:        gameCreatorStub{err: nil},
		},
	}

	cc := gordle.NewCorpusCache()

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			f := Handler(cc, tc.creator, corpusPath)

			req, err := http.NewRequest(http.MethodPost, api.NewGameRoute, nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			f.ServeHTTP(recorder, req)

			assert.Equal(t, tc.wantStatusCode, recorder.Code)

			if tc.wantBody == "" {
				return
			}

			body := recorder.Body.String()
			id := idFinderRegExp.FindStringSubmatch(body)

			if len(id) != 2 {
				t.Fatal("cannot find id in json output")
			}
			body = strings.Replace(body, id[1], "123456", 1)

			assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
			assert.JSONEq(t, tc.wantBody, body)
		})
	}

}

func Test_createGame(t *testing.T) {
	cc := gordle.NewCorpusCache()
	corpusPath := "testdata/corpus.txt"

	g, err := createGame(cc, gameCreatorStub{nil}, corpusPath)
	require.NoError(t, err)

	assert.Regexp(t, "[A-Z0-9]+", g.ID)
	assert.Equal(t, byte(5), g.AttemptsLeft)
	assert.Equal(t, 0, len(g.Guesses))
}
