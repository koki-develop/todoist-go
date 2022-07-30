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

func TestClient_GetTaskComments(t *testing.T) {
	type args struct {
		taskID int
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
			args: args{taskID: 1},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id": 1, "content": "COMMENT_1" }, { "id": 2, "content": "COMMENT_2" }]`),
			},
			want:    Comments{{ID: 1, Content: "COMMENT_1"}, {ID: 2, Content: "COMMENT_2"}},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{taskID: 1},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/comments?task_id=%d", tt.args.taskID),
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			cmts, err := cl.GetTaskComments(tt.args.taskID)

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

func TestClient_GetComment(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Comment
		wantErr bool
	}{
		{
			name: "should return a comment",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "content": "COMMENT" }`),
			},
			want:    &Comment{ID: 1, Content: "COMMENT"},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/comments/%d", tt.args.id),
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			cmt, err := cl.GetComment(tt.args.id)

			assert.Equal(t, tt.want, cmt)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CreateProjectComment(t *testing.T) {
	type args struct {
		projectID int
		content   string
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Comment
		wantErr bool
	}{
		{
			name: "should return a comment",
			args: args{projectID: 1, content: "COMMENT"},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "content": "COMMENT" }`),
			},
			want:    &Comment{ID: 1, Content: "COMMENT"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{projectID: 1, content: "COMMENT"},
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
				URL:     "https://api.todoist.com/rest/v1/comments",
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"project_id": tt.args.projectID, "content": tt.args.content},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json"},
			}).Return(tt.resp, nil)

			cmt, err := cl.CreateProjectComment(tt.args.projectID, tt.args.content)

			assert.Equal(t, tt.want, cmt)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CreateProjectCommentWithOptions(t *testing.T) {
	type args struct {
		projectID int
		content   string
		opts      *CreateProjectCommentOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Comment
		wantErr bool
	}{
		{
			name: "should return a comment",
			args: args{projectID: 1, content: "COMMENT", opts: &CreateProjectCommentOptions{
				RequestID: String("REQUEST_ID"),
				Attachment: &CreateAttachmentOptions{
					ResourceType: String("RESOURCE_TYPE"),
					FileName:     String("FILE_NAME"),
					FileURL:      String("FILE_URL"),
					FileType:     String("FILE_TYPE"),
				},
			}},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "content": "COMMENT" }`),
			},
			want:    &Comment{ID: 1, Content: "COMMENT"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{projectID: 1, content: "COMMENT", opts: &CreateProjectCommentOptions{
				RequestID: String("REQUEST_ID"),
				Attachment: &CreateAttachmentOptions{
					ResourceType: String("RESOURCE_TYPE"),
					FileName:     String("FILE_NAME"),
					FileURL:      String("FILE_URL"),
					FileType:     String("FILE_TYPE"),
				},
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
				URL:    "https://api.todoist.com/rest/v1/comments",
				Method: http.MethodPost,
				Payload: map[string]interface{}{
					"project_id": tt.args.projectID,
					"content":    tt.args.content,
					"attachment": map[string]string{
						"resource_type": *tt.args.opts.Attachment.ResourceType,
						"file_name":     *tt.args.opts.Attachment.FileName,
						"file_url":      *tt.args.opts.Attachment.FileURL,
						"file_type":     *tt.args.opts.Attachment.FileType,
					},
				},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			cmt, err := cl.CreateProjectCommentWithOptions(tt.args.projectID, tt.args.content, tt.args.opts)

			assert.Equal(t, tt.want, cmt)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CreateTaskComment(t *testing.T) {
	type args struct {
		taskID  int
		content string
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Comment
		wantErr bool
	}{
		{
			name: "should return a comment",
			args: args{taskID: 1, content: "COMMENT"},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "content": "COMMENT" }`),
			},
			want:    &Comment{ID: 1, Content: "COMMENT"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{taskID: 1, content: "COMMENT"},
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
				URL:     "https://api.todoist.com/rest/v1/comments",
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"task_id": tt.args.taskID, "content": tt.args.content},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json"},
			}).Return(tt.resp, nil)

			cmt, err := cl.CreateTaskComment(tt.args.taskID, tt.args.content)

			assert.Equal(t, tt.want, cmt)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CreateTaskCommentWithOptions(t *testing.T) {
	type args struct {
		taskID  int
		content string
		opts    *CreateTaskCommentOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Comment
		wantErr bool
	}{
		{
			name: "should return a comment",
			args: args{taskID: 1, content: "COMMENT", opts: &CreateTaskCommentOptions{
				RequestID: String("REQUEST_ID"),
				Attachment: &CreateAttachmentOptions{
					ResourceType: String("RESOURCE_TYPE"),
					FileName:     String("FILE_NAME"),
					FileURL:      String("FILE_URL"),
					FileType:     String("FILE_TYPE"),
				},
			}},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "content": "COMMENT" }`),
			},
			want:    &Comment{ID: 1, Content: "COMMENT"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{taskID: 1, content: "COMMENT", opts: &CreateTaskCommentOptions{
				RequestID: String("REQUEST_ID"),
				Attachment: &CreateAttachmentOptions{
					ResourceType: String("RESOURCE_TYPE"),
					FileName:     String("FILE_NAME"),
					FileURL:      String("FILE_URL"),
					FileType:     String("FILE_TYPE"),
				},
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
				URL:    "https://api.todoist.com/rest/v1/comments",
				Method: http.MethodPost,
				Payload: map[string]interface{}{
					"task_id": tt.args.taskID,
					"content": tt.args.content,
					"attachment": map[string]string{
						"resource_type": *tt.args.opts.Attachment.ResourceType,
						"file_name":     *tt.args.opts.Attachment.FileName,
						"file_url":      *tt.args.opts.Attachment.FileURL,
						"file_type":     *tt.args.opts.Attachment.FileType,
					},
				},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			cmt, err := cl.CreateTaskCommentWithOptions(tt.args.taskID, tt.args.content, tt.args.opts)

			assert.Equal(t, tt.want, cmt)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_UpdateComment(t *testing.T) {
	type args struct {
		id      int
		content string
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, content: "COMMENT"},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, content: "COMMENT"},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/comments/%d", tt.args.id),
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"content": tt.args.content},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json"},
			}).Return(tt.resp, nil)

			err := cl.UpdateComment(tt.args.id, tt.args.content)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_UpdateCommentWithOptions(t *testing.T) {
	type args struct {
		id      int
		content string
		opts    *UpdateCommentOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{
				id:      1,
				content: "COMMENT",
				opts:    &UpdateCommentOptions{RequestID: String("REQUEST_ID")},
			},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{
				id:      1,
				content: "COMMENT",
				opts:    &UpdateCommentOptions{RequestID: String("REQUEST_ID")},
			},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/comments/%d", tt.args.id),
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"content": tt.args.content},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.UpdateCommentWithOptions(tt.args.id, tt.args.content, tt.args.opts)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_DeleteComment(t *testing.T) {
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
				StatusCode: http.StatusBadGateway,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/comments/%d", tt.args.id),
				Method:  http.MethodDelete,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			err := cl.DeleteComment(tt.args.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_DeleteCommentWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *DeleteCommentOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &DeleteCommentOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &DeleteCommentOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusBadGateway,
				Body:       strings.NewReader("ERROR_RESPONSE"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl, api := newClientForTest()

			api.On("Do", &restRequest{
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/comments/%d", tt.args.id),
				Method:  http.MethodDelete,
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.DeleteCommentWithOptions(tt.args.id, tt.args.opts)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}
