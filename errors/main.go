package errors

import (
	"fmt"
	"io"
	"net/http"
)

type ResponseError struct {
	Response *http.Response
}

func (e ResponseError) Error() string {
	body, _ := io.ReadAll(e.Response.Body)
	return fmt.Sprintf(
		"%s: %d\n%s",
		"unexpected response",
		e.Response.StatusCode,
		body,
	)
}
