package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	apiBaseUrl string = "https://api.todoist.com/rest/v1"
)

type Client struct {
	token string
}

func New(token string) *Client {
	return &Client{token}
}

func (cl *Client) get(p string, params map[string]string, expectedStatus int, out interface{}) error {
	u, err := url.Parse(apiBaseUrl)
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

func (cl *Client) post(p string, params map[string]interface{}, expectedStatus int, reqID *string, out interface{}) error {
	u, err := url.Parse(apiBaseUrl)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, p)

	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))
	if reqID != nil {
		req.Header.Set("X-Request-Id", *reqID)
	}

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
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return err
	}

	return nil
}
