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
	ep, err := cl.buildEndpoint(p, params)
	if err != nil {
		return err
	}

	req, err := cl.buildRequest(ep, http.MethodGet, nil, nil)
	if err != nil {
		return err
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

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return err
	}

	return nil
}

func (cl *Client) post(p string, payload map[string]interface{}, expectedStatus int, reqID *string, out interface{}) error {
	ep, err := cl.buildEndpoint(p, nil)
	if err != nil {
		return err
	}

	req, err := cl.buildRequest(ep, http.MethodPost, payload, reqID)
	if err != nil {
		return err
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

func (cl *Client) delete(p string, expectedStatus int, reqID *string) error {
	ep, err := cl.buildEndpoint(p, nil)
	if err != nil {
		return err
	}

	req, err := cl.buildRequest(ep, http.MethodDelete, nil, reqID)
	if err != nil {
		return err
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

	return nil
}

func (cl *Client) buildEndpoint(p string, params map[string]string) (string, error) {
	u, err := url.Parse(apiBaseUrl)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, p)

	if params != nil {
		q := u.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		u.RawQuery = q.Encode()
	}

	return u.String(), nil
}

func (cl *Client) buildRequest(ep string, method string, payload map[string]interface{}, reqID *string) (*http.Request, error) {
	var b io.Reader
	if payload != nil {
		j, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		b = bytes.NewBuffer(j)
	}

	req, err := http.NewRequest(method, ep, b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))
	if reqID != nil {
		req.Header.Set("X-Request-Id", *reqID)
	}

	return req, nil
}
