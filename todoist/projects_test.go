package todoist

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetProjects(t *testing.T) {
	tests := []struct {
		name    string
		resp    *restResponse
		want    Projects
		wantErr bool
	}{
		{
			name: "should return projects",
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
				URL:     "https://api.todoist.com/rest/v1/projects",
				Method:  http.MethodGet,
				Payload: nil,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			projs, err := cl.GetProjects()

			assert.Equal(t, tt.want, projs)
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

func TestClient_GetProject(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Project
		wantErr bool
	}{
		{
			name: "should return a project",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id":1 , "name": "PROJECT_1" }`),
			},
			want:    &Project{ID: 1, Name: "PROJECT_1"},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/projects/%d", tt.args.id),
				Method:  http.MethodGet,
				Payload: nil,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			proj, err := cl.GetProject(tt.args.id)

			assert.Equal(t, tt.want, proj)
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

func TestClient_CreateProject(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Project
		wantErr bool
	}{
		{
			name: "should return a project",
			args: args{name: "NEW_PROJECT"},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "name": "NEW_PROJECT" }`),
			},
			want:    &Project{ID: 1, Name: "NEW_PROJECT"},
			wantErr: false,
		},
		{
			name: "return an error if the request fails",
			args: args{name: "NEW_PROJECT"},
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
				URL:     "https://api.todoist.com/rest/v1/projects",
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"name": tt.args.name},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json"},
			}).Return(tt.resp, nil)

			proj, err := cl.CreateProject(tt.args.name)

			assert.Equal(t, tt.want, proj)
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, RequestError{}, err)
			} else {
				assert.Equal(t, tt.want, proj)
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CreateProjectWithOptions(t *testing.T) {
	type args struct {
		name string
		opts *CreateProjectOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Project
		wantErr bool
	}{
		{
			name: "should return a project",
			args: args{name: "NEW_PROJECT", opts: &CreateProjectOptions{ParentID: Int(2), Color: Int(30), Favorite: Bool(true), RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "name": "NEW_PROJECT" }`),
			},
			want:    &Project{ID: 1, Name: "NEW_PROJECT"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails.",
			args: args{name: "NEW_PROJECT", opts: &CreateProjectOptions{ParentID: Int(2), Color: Int(30), Favorite: Bool(true), RequestID: String("REQUEST_ID")}},
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
				URL:     "https://api.todoist.com/rest/v1/projects",
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"name": tt.args.name, "parent_id": *tt.args.opts.ParentID, "color": *tt.args.opts.Color, "favorite": *tt.args.opts.Favorite},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			proj, err := cl.CreateProjectWithOptions(tt.args.name, tt.args.opts)

			assert.Equal(t, tt.want, proj)
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

func TestClient_UpdateProjectWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *UpdateProjectOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &UpdateProjectOptions{Name: String("UPDATED_PROJECT"), Color: Int(99), Favorite: Bool(true), RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &UpdateProjectOptions{Name: String("UPDATED_PROJECT"), Color: Int(99), Favorite: Bool(true), RequestID: String("REQUEST_ID")}},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/projects/%d", tt.args.id),
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"name": *tt.args.opts.Name, "color": *tt.args.opts.Color, "favorite": *tt.args.opts.Favorite},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.UpdateProjectWithOptions(tt.args.id, tt.args.opts)

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

func TestClient_DeleteProject(t *testing.T) {
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/projects/%d", tt.args.id),
				Method:  http.MethodDelete,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			err := cl.DeleteProject(tt.args.id)

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

func TestClient_DeleteProjectWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *DeleteProjectOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &DeleteProjectOptions{RequestID: String("REQUEST_ID")}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &DeleteProjectOptions{RequestID: String("REQUEST_ID")}},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/projects/%d", tt.args.id),
				Method:  http.MethodDelete,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			err := cl.DeleteProject(tt.args.id)

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

func TestClient_GetCollaborators(t *testing.T) {
	type args struct {
		projectID int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    Users
		wantErr bool
	}{
		{
			name: "should return users",
			args: args{projectID: 1},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id": 1, "name": "USER_1", "email": "user1@example.com" }, { "id": 2, "name": "USER_2", "email": "user2@example.com" }]`),
			},
			want:    Users{{ID: 1, Name: "USER_1", Email: "user1@example.com"}, {ID: 2, Name: "USER_2", Email: "user2@example.com"}},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/projects/%d/collaborators", tt.args.projectID),
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			users, err := cl.GetCollaborators(tt.args.projectID)

			assert.Equal(t, tt.want, users)
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
