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
				assert.IsType(t, RequestError{}, err)
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
				IDs:       Ints(4, 5, 6),
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
				IDs:       Ints(4, 5, 6),
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
					strings.Join(intsToStrings(*tt.args.opts.IDs), "%2C"),
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
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_GetTask(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Task
		wantErr bool
	}{
		{
			name: "should return a task",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "content": "TASK" }`),
			},
			want:    &Task{ID: 1, Content: "TASK"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%d", tt.args.id),
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			task, err := cl.GetTask(tt.args.id)

			assert.Equal(t, tt.want, task)
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CreateTask(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Task
		wantErr bool
	}{
		{
			name: "should return a task",
			args: args{content: "TASK"},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "content": "TASK" }`),
			},
			want:    &Task{ID: 1, Content: "TASK"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{content: "TASK"},
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
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"content": tt.args.content},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json"},
			}).Return(tt.resp, nil)

			task, err := cl.CreateTask(tt.args.content)

			assert.Equal(t, tt.want, task)
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CreateTaskWithOptions(t *testing.T) {
	type args struct {
		content string
		opts    *CreateTaskOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Task
		wantErr bool
	}{
		{
			name: "should return a task",
			args: args{
				content: "TASK",
				opts: &CreateTaskOptions{
					RequestID:   String("REQUEST_ID"),
					Description: String("DESCRIPTION"),
					ProjectID:   Int(1),
					SectionID:   Int(2),
					ParentID:    Int(3),
					Order:       Int(4),
					LabelIDs:    Ints(5, 6, 7),
					Priority:    Int(8),
					DueString:   String("DUE_STRING"),
					DueDate:     String("DUE_DATE"),
					DueDatetime: String("DUE_DATETIME"),
					DueLang:     String("DUE_LANG"),
					Assignee:    Int(9),
				},
			},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "content": "TASK" }`),
			},
			want:    &Task{ID: 1, Content: "TASK"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{
				content: "TASK",
				opts: &CreateTaskOptions{
					RequestID:   String("REQUEST_ID"),
					Description: String("DESCRIPTION"),
					ProjectID:   Int(1),
					SectionID:   Int(2),
					ParentID:    Int(3),
					Order:       Int(4),
					LabelIDs:    Ints(5, 6, 7),
					Priority:    Int(8),
					DueString:   String("DUE_STRING"),
					DueDate:     String("DUE_DATE"),
					DueDatetime: String("DUE_DATETIME"),
					DueLang:     String("DUE_LANG"),
					Assignee:    Int(9),
				},
			},
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
				URL:    "https://api.todoist.com/rest/v1/tasks",
				Method: http.MethodPost,
				Payload: map[string]interface{}{
					"content":      tt.args.content,
					"description":  *tt.args.opts.Description,
					"project_id":   *tt.args.opts.ProjectID,
					"section_id":   *tt.args.opts.SectionID,
					"parent_id":    *tt.args.opts.ParentID,
					"order":        *tt.args.opts.Order,
					"label_ids":    strings.Join(intsToStrings(*tt.args.opts.LabelIDs), ","),
					"priority":     *tt.args.opts.Priority,
					"due_string":   *tt.args.opts.DueString,
					"due_date":     *tt.args.opts.DueDate,
					"due_datetime": *tt.args.opts.DueDatetime,
					"due_lang":     *tt.args.opts.DueLang,
					"assignee":     *tt.args.opts.Assignee,
				},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			task, err := cl.CreateTaskWithOptions(tt.args.content, tt.args.opts)

			assert.Equal(t, tt.want, task)
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_UpdateTaskWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *UpdateTaskOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &UpdateTaskOptions{
				RequestID:   String("REQUEST_ID"),
				Content:     String("TASK"),
				Description: String("DESCRIPTION"),
				LabelIDs:    Ints(1, 2, 3),
				Priority:    Int(4),
				DueString:   String("DUE_STRING"),
				DueDate:     String("DUE_DATE"),
				DueDatetime: String("DUE_DATETIME"),
				DueLang:     String("DUE_LANG"),
				Assignee:    Int(5),
			}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &UpdateTaskOptions{
				RequestID:   String("REQUEST_ID"),
				Content:     String("TASK"),
				Description: String("DESCRIPTION"),
				LabelIDs:    Ints(1, 2, 3),
				Priority:    Int(4),
				DueString:   String("DUE_STRING"),
				DueDate:     String("DUE_DATE"),
				DueDatetime: String("DUE_DATETIME"),
				DueLang:     String("DUE_LANG"),
				Assignee:    Int(5),
			}},
			resp: &restResponse{
				StatusCode: http.StatusBadRequest,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:    fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%d", tt.args.id),
				Method: http.MethodPost,
				Payload: map[string]interface{}{
					"content":      *tt.args.opts.Content,
					"description":  *tt.args.opts.Description,
					"label_ids":    tt.args.opts.LabelIDs,
					"priority":     *tt.args.opts.Priority,
					"due_string":   *tt.args.opts.DueString,
					"due_date":     *tt.args.opts.DueDate,
					"due_datetime": *tt.args.opts.DueDatetime,
					"due_lang":     *tt.args.opts.DueLang,
					"assignee":     *tt.args.opts.Assignee,
				},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.UpdateTaskWithOptions(tt.args.id, tt.args.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CloseTask(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusBadRequest,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%d/close", tt.args.id),
				Method:  http.MethodPost,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			err := cl.CloseTask(tt.args.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CloseTaskWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *CloseTaskOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &CloseTaskOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &CloseTaskOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusBadRequest,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%d/close", tt.args.id),
				Method:  http.MethodPost,
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.CloseTaskWithOptions(tt.args.id, tt.args.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_ReopenTask(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusBadRequest,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%d/reopen", tt.args.id),
				Method:  http.MethodPost,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			err := cl.ReopenTask(tt.args.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_ReopenTaskWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *ReopenTaskOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &ReopenTaskOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &ReopenTaskOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusBadRequest,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%d/reopen", tt.args.id),
				Method:  http.MethodPost,
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.ReopenTaskWithOptions(tt.args.id, tt.args.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_DeleteTask(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusBadRequest,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%d", tt.args.id),
				Method:  http.MethodDelete,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			err := cl.DeleteTask(tt.args.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_DeleteTaskWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *DeleteTaskOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &DeleteTaskOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &DeleteTaskOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusBadRequest,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/tasks/%d", tt.args.id),
				Method:  http.MethodDelete,
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.DeleteTaskWithOptions(tt.args.id, tt.args.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}
