package todoist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	token string
}

func New(token string) *Client {
	return &Client{token}
}

func (cl *Client) get(p string, params map[string]string, expectedStatus int, out interface{}) error {
	u, err := url.Parse("https://api.todoist.com/rest")
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, p)

	if params != nil {
		q := u.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatus {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(b))
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return err
	}

	return nil
}
