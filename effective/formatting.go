package effective

import (
	"sync"

	"github.com/pkg/errors"
)

type T struct {
	name         string
	value        int
	withLongName float32
}

// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*string, error) {
	return nil, nil
}

// Error codes returned by failures to parse an expression
var (
	ErrInternal      = errors.New("regexp: internal error")
	ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
	ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
)

var (
	countLock   sync.Mutex
	inputCount  uint32
	outputCount uint32
	errorCount  uint32
)
