package error

import (
	"errors"
)

var (
	ErrShortWrite      = errors.New("framework: short write")
	ErrShortBuffer     = errors.New("framework: short buffer")
	EOF                = errors.New("framework: EOF")
	ErrUnexpectedEOF   = errors.New("framework: unexpected EOF")
	ErrNoProgress      = errors.New("framework: multiple Read calls return no data or error")
	ErrServerClosed    = errors.New("framework: server was closed")
	ErrPoolLimit       = errors.New("framework: per-server connection limit reached")
	ErrNotFountInstant = errors.New("framework: not found Instantiated")
)
