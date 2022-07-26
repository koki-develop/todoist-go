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

const (
	apiBaseUrl string = "https://api.todoist.com/rest"
)

type Client struct {
	token string

	restAPI restAPI
}

func New(token string) *Client {
	return &Client{
		token:   token,
		restAPI: newRESTClient(),
	}
}

func (cl *Client) get(p string, params map[string]string, expectedStatus int, out interface{}) error {
	ep, err := cl.buildEndpoint(p, params)
	if err != nil {
		return err
	}

	body, err := cl.sendRequest(ep, http.MethodGet, nil, nil, expectedStatus)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(body).Decode(out); err != nil {
		return err
	}

	return nil
}

func (cl *Client) post(p string, payload map[string]interface{}, expectedStatus int, reqID *string, out interface{}) error {
	ep, err := cl.buildEndpoint(p, nil)
	if err != nil {
		return err
	}

	body, err := cl.sendRequest(ep, http.MethodPost, payload, reqID, expectedStatus)
	if err != nil {
		return err
	}
	if expectedStatus == http.StatusNoContent {
		return nil
	}

	if err := json.NewDecoder(body).Decode(out); err != nil {
		return err
	}

	return nil
}

func (cl *Client) delete(p string, expectedStatus int, reqID *string) error {
	ep, err := cl.buildEndpoint(p, nil)
	if err != nil {
		return err
	}

	if _, err := cl.sendRequest(ep, http.MethodDelete, nil, reqID, expectedStatus); err != nil {
		return err
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

func (cl *Client) buildRequest(ep, method string, payload map[string]interface{}, reqID *string) *restRequest {
	h := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", cl.token),
	}
	if reqID != nil {
		h["X-Request-Id"] = *reqID
	}
	if payload != nil {
		h["Content-Type"] = "application/json"
	}

	return &restRequest{
		URL:     ep,
		Method:  method,
		Payload: payload,
		Headers: h,
	}
}

func (cl *Client) sendRequest(ep, method string, payload map[string]interface{}, reqID *string, expectedStatusCode int) (io.Reader, error) {
	req := cl.buildRequest(ep, method, payload, reqID)

	resp, err := cl.restAPI.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != expectedStatusCode {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}

	return resp.Body, nil
}
