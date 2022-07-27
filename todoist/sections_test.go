package todoist

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetSections(t *testing.T) {
	tests := []struct {
		name    string
		resp    *restResponse
		want    Sections
		wantErr bool
	}{
		{
			name: "should return sections",
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id": 1, "project_id": 1, "order": 1, "name": "SECTION_1" }, { "id": 2, "project_id": 2, "order": 2, "name": "SECTION_2" }]`),
			},
			want: Sections{
				{ID: 1, ProjectID: 1, Order: 1, Name: "SECTION_1"},
				{ID: 2, ProjectID: 2, Order: 2, Name: "SECTION_2"},
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
				URL:     "https://api.todoist.com/rest/v1/sections",
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			secs, err := cl.GetSections()

			assert.Equal(t, tt.want, secs)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_GetSectionsWithOptions(t *testing.T) {
	type args struct {
		opts *GetSectionsOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    Sections
		wantErr bool
	}{
		{
			name: "should return sections",
			args: args{opts: &GetSectionsOptions{ProjectID: Int(1)}},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id": 1, "project_id": 1, "order": 1, "name": "SECTION_1" }, { "id": 2, "project_id": 2, "order": 2, "name": "SECTION_2" }]`),
			},
			want: Sections{
				{ID: 1, ProjectID: 1, Order: 1, Name: "SECTION_1"},
				{ID: 2, ProjectID: 2, Order: 2, Name: "SECTION_2"},
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{opts: &GetSectionsOptions{ProjectID: Int(1)}},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/sections?project_id=%d", *tt.args.opts.ProjectID),
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			secs, err := cl.GetSectionsWithOptions(tt.args.opts)

			assert.Equal(t, tt.want, secs)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}

func TestClient_CreateSectionWithOptions(t *testing.T) {
	type args struct {
		name      string
		projectID int
		opts      *CreateSectionOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Section
		wantErr bool
	}{
		{
			name: "should return section",
			args: args{name: "SECTION", projectID: 1, opts: &CreateSectionOptions{RequestID: String("REQUEST_ID"), Order: Int(1)}},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "name": "SECTION" }`),
			},
			want:    &Section{ID: 1, Name: "SECTION"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{name: "SECTION", projectID: 1, opts: &CreateSectionOptions{RequestID: String("REQUEST_ID"), Order: Int(1)}},
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
				URL:     "https://api.todoist.com/rest/v1/sections",
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"name": tt.args.name, "project_id": tt.args.projectID, "order": *tt.args.opts.Order},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			sec, err := cl.CreateSectionWithOptions(tt.args.name, tt.args.projectID, tt.args.opts)

			assert.Equal(t, tt.want, sec)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}
