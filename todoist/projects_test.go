package todoist

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newClientForTest() (*Client, *mockRestAPI) {
	api := &mockRestAPI{}
	return &Client{token: "TOKEN", restAPI: api}, api
}

func TestClient_GetProjects(t *testing.T) {
	cl, api := newClientForTest()

	tests := []struct {
		name    string
		resp    *restResponse
		want    Projects
		wantErr bool
	}{
		{
			name: "return projects when succeeded.",
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id":1 , "name": "PROJECT_1" }, { "id":2 , "name": "PROJECT_2" }]`),
			},
			want: Projects{
				{ID: 1, Name: "PROJECT_1"},
				{ID: 2, Name: "PROJECT_2"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api.On("Do", &restRequest{
				URL:     "https://api.todoist.com/rest/v1/projects",
				Method:  http.MethodGet,
				Payload: nil,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			projs, err := cl.GetProjects()

			if tt.wantErr {
				assert.Nil(t, projs)
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, projs)
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}
