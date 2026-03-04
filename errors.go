package ali

import (
	"errors"
	"fmt"
)

var (
	ErrResponse = errors.New("response indicates an error")
)

type ResponseError struct {
	StatusCode int
	Body       []byte
}

func (e ResponseError) Unwrap() error {
	return ErrResponse
}

func (e ResponseError) Error() string {
	return fmt.Sprintf(
		"%s (status:%d\nbody:%s\n)",
		"bad response",
		e.StatusCode,
		string(e.Body),
	)
}
