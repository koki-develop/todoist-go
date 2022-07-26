package todoist

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type httpAPI interface {
	Do(req *http.Request) (*http.Response, error)
}

type restClient struct {
	httpAPI httpAPI
}

type restRequest struct {
	URL     string
	Method  string
	Payload map[string]interface{}
	Headers map[string]string
}

type restResponse struct {
	StatusCode int
	Body       io.Reader
}

func newRESTClient() *restClient {
	return &restClient{httpAPI: new(http.Client)}
}

func (cl *restClient) Do(req *restRequest) (*restResponse, error) {
	var p io.Reader
	if req.Payload != nil {
		j, err := json.Marshal(req.Payload)
		if err != nil {
			return nil, err
		}
		p = bytes.NewBuffer(j)
	}
	httpreq, err := http.NewRequest(req.Method, req.URL, p)
	if err != nil {
		return nil, err
	}
	for k, v := range req.Headers {
		httpreq.Header.Set(k, v)
	}

	resp, err := cl.httpAPI.Do(httpreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return nil, err
	}

	return &restResponse{
		StatusCode: resp.StatusCode,
		Body:       buf,
	}, nil
}
