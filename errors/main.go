package errors

import (
	"fmt"
	"net/http"
)

type ResponseError struct {
	Response *http.Response
}

func (e ResponseError) Error() string {
	return fmt.Sprintf(
		"%s: %d",
		"unexpected response",
		e.Response.StatusCode,
	)
}
