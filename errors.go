package todoist

import (
	"fmt"
	"io"
)

type RequestError struct {
	// Error response status code.
	StatusCode int
	// Error response body.
	Body io.Reader
}

func (err RequestError) Error() string {
	return fmt.Sprintf("request error: %d", err.StatusCode)
}

func newRequestError(resp *restResponse) (RequestError, error) {
	return RequestError{StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
