package ali

import (
	"errors"
	"fmt"
)

var (
	ErrNotImplemented = errors.New("feature is not implemented")
	ErrBadResponse    = errors.New("request produced a bad response")
)

type ResponseError struct {
	StatusCode int
	Body       []byte
}

func (e ResponseError) Unwrap() error {
	return ErrBadResponse
}

func (e ResponseError) Error() string {
	return fmt.Sprintf(
		"%s (status:%d\nbody:%s\n)",
		"bad response",
		e.StatusCode,
		string(e.Body),
	)
}
