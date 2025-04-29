package gordle

// corpusError defines a sentinel error
type corpusError string

// Error is the implementation of the error interface by corpusError
func (e corpusError) Error() string {
	return string(e)
}

// GameError defines an error that happens during a game.
type GameError string

// Error is the implementation of the error interface by GameError
func (e GameError) Error() string {
	return string(e)
}
