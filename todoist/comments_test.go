package todoist

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetProjectComments(t *testing.T) {
	type args struct {
		projectID int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    Comments
		wantErr bool
	}{
		{
			name: "should return comments",
			args: args{projectID: 1},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id": 1, "content": "COMMENT_1" }, { "id": 2, "content": "COMMENT_2" }]`),
			},
			want:    Comments{{ID: 1, Content: "COMMENT_1"}, {ID: 2, Content: "COMMENT_2"}},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{projectID: 1},
			resp: &restResponse{
				StatusCode: http.StatusBadRequest,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/comments?project_id=%d", tt.args.projectID),
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			cmts, err := cl.GetProjectComments(tt.args.projectID)

			assert.Equal(t, tt.want, cmts)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}
