package todoist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/google/go-querystring/query"
)

const (
	apiBaseUrl string = "https://api.todoist.com/rest"
)

// Client for Todoist REST API.
type Client struct {
	token string

	restAPI restAPI
}

// Returns new client.
func New(token string) *Client {
	return &Client{
		token:   token,
		restAPI: newRESTClient(),
	}
}

func (cl *Client) get(p string, params interface{}, out interface{}) error {
	body, err := cl.sendRequest(p, params, http.MethodGet, nil, nil)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(body).Decode(out); err != nil {
		return err
	}

	return nil
}

func (cl *Client) post(p string, payload map[string]interface{}, reqID *string, out interface{}) error {
	body, err := cl.sendRequest(p, nil, http.MethodPost, payload, reqID)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(body).Decode(out); err != nil {
		return err
	}

	return nil
}

func (cl *Client) postWithoutBind(p string, payload map[string]interface{}, reqID *string) error {
	if _, err := cl.sendRequest(p, nil, http.MethodPost, payload, reqID); err != nil {
		return err
	}

	return nil
}

func (cl *Client) delete(p string, reqID *string) error {
	if _, err := cl.sendRequest(p, nil, http.MethodDelete, nil, reqID); err != nil {
		return err
	}

	return nil
}

func (cl *Client) buildEndpoint(p string, params interface{}) (string, error) {
	u, err := url.Parse(apiBaseUrl)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, p)

	if params != nil {
		v, err := query.Values(params)
		if err != nil {
			return "", err
		}
		u.RawQuery = v.Encode()
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

func (cl *Client) sendRequest(p string, params interface{}, method string, payload map[string]interface{}, reqID *string) (io.Reader, error) {
	ep, err := cl.buildEndpoint(p, params)
	if err != nil {
		return nil, err
	}

	req := cl.buildRequest(ep, method, payload, reqID)
	resp, err := cl.restAPI.Do(req)
	if err != nil {
		return nil, err
	}

	if 200 <= resp.StatusCode && resp.StatusCode <= 299 {
		return resp.Body, nil
	}

	reqerr, err := newRequestError(resp)
	if err != nil {
		return nil, err
	}
	return nil, reqerr
}
