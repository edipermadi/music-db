package theory

import "errors"

// Error messages
var (
	ErrScaleNotFound        = errors.New("scale not found")
	ErrKeyNotFound          = errors.New("key not found")
	ErrChordNotFound        = errors.New("chord not found")
	ErrChordQualityNotFound = errors.New("chord quality not found")
)
