package todoist

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetTasks(t *testing.T) {
	tests := []struct {
		name    string
		resp    *restResponse
		want    Tasks
		wantErr bool
	}{
		{
			name: "should return tasks",
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id": 1, "content": "TASK_1" }, { "id": 2, "content": "TASK_2" }]`),
			},
			want:    Tasks{{ID: 1, Content: "TASK_1"}, {ID: 2, Content: "TASK_2"}},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
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
				URL:     "https://api.todoist.com/rest/v1/tasks",
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			tasks, err := cl.GetTasks()

			assert.Equal(t, tt.want, tasks)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_GetTasksWithOptions(t *testing.T) {
	type args struct {
		opts *GetTasksOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    Tasks
		wantErr bool
	}{
		{
			name: "should return tasks",
			args: args{opts: &GetTasksOptions{
				ProjectID: Int(1),
				SectionID: Int(2),
				LabelID:   Int(3),
				Filter:    String("FILTER"),
				Lang:      String("LANG"),
				IDs:       []int{4, 5, 6},
			}},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id": 1, "content": "TASK_1" }, { "id": 2, "content": "TASK_2" }]`),
			},
			want:    Tasks{{ID: 1, Content: "TASK_1"}, {ID: 2, Content: "TASK_2"}},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{opts: &GetTasksOptions{
				ProjectID: Int(1),
				SectionID: Int(2),
				LabelID:   Int(3),
				Filter:    String("FILTER"),
				Lang:      String("LANG"),
				IDs:       []int{4, 5, 6},
			}},
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
				URL: fmt.Sprintf(
					"https://api.todoist.com/rest/v1/tasks?filter=%s&ids=%s&label_id=%d&lang=%s&project_id=%d&section_id=%d",
					*tt.args.opts.Filter,
					strings.Join(intsToStrings(tt.args.opts.IDs), "%2C"),
					*tt.args.opts.LabelID,
					*tt.args.opts.Lang,
					*tt.args.opts.ProjectID,
					*tt.args.opts.SectionID,
				),
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			tasks, err := cl.GetTasksWithOptions(tt.args.opts)

			assert.Equal(t, tt.want, tasks)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}
