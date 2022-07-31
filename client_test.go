package todoist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newClientForTest() (*Client, *mockRestAPI) {
	api := &mockRestAPI{}
	return &Client{token: "TOKEN", restAPI: api}, api
}

func TestNew(t *testing.T) {
	t.Run("should return a client", func(t *testing.T) {
		tkn := "TOKEN"
		cl := New(tkn)

		assert.NotNil(t, cl)
		assert.NotNil(t, cl.restAPI)
		assert.IsType(t, &Client{}, cl)
		assert.Equal(t, tkn, cl.token)
	})
}
