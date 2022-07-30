package todoist

import (
	"bytes"
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
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return RequestError{}, err
	}

	return RequestError{StatusCode: resp.StatusCode, Body: buf}, nil
}
