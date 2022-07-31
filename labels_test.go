package todoist

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetLabels(t *testing.T) {
	tests := []struct {
		name    string
		resp    *restResponse
		want    Labels
		wantErr bool
	}{
		{
			name: "should return labels",
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`[{ "id": 1, "name": "LABEL_1" }, { "id": 2, "name": "LABEL_2" }]`),
			},
			want:    Labels{{ID: 1, Name: "LABEL_1"}, {ID: 2, Name: "LABEL_2"}},
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
				URL:     "https://api.todoist.com/rest/v1/labels",
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			labels, err := cl.GetLabels()

			assert.Equal(t, tt.want, labels)
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

func TestClient_GetLabel(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Label
		wantErr bool
	}{
		{
			name: "should return a label",
			args: args{id: 1},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "name": "LABEL" }`),
			},
			want:    &Label{ID: 1, Name: "LABEL"},
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/labels/%d", tt.args.id),
				Method:  http.MethodGet,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			label, err := cl.GetLabel(tt.args.id)

			assert.Equal(t, tt.want, label)
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

func TestClient_CreateLabel(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Label
		wantErr bool
	}{
		{
			name: "should return a label",
			args: args{name: "LABEL"},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "name": "LABEL" }`),
			},
			want:    &Label{ID: 1, Name: "LABEL"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{name: "LABEL"},
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
				URL:     "https://api.todoist.com/rest/v1/labels",
				Method:  http.MethodPost,
				Payload: map[string]interface{}{"name": tt.args.name},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json"},
			}).Return(tt.resp, nil)

			label, err := cl.CreateLabel(tt.args.name)

			assert.Equal(t, tt.want, label)
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

func TestClient_CreateLabelWithOptions(t *testing.T) {
	type args struct {
		name string
		opts *CreateLabelOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		want    *Label
		wantErr bool
	}{
		{
			name: "should return a label",
			args: args{name: "LABEL", opts: &CreateLabelOptions{
				RequestID: String("REQUEST_ID"),
				Order:     Int(1),
				Color:     Int(2),
				Favorite:  Bool(true),
			}},
			resp: &restResponse{
				StatusCode: http.StatusOK,
				Body:       strings.NewReader(`{ "id": 1, "name": "LABEL" }`),
			},
			want:    &Label{ID: 1, Name: "LABEL"},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{name: "LABEL", opts: &CreateLabelOptions{
				RequestID: String("REQUEST_ID"),
				Order:     Int(1),
				Color:     Int(2),
				Favorite:  Bool(true),
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
				URL:    "https://api.todoist.com/rest/v1/labels",
				Method: http.MethodPost,
				Payload: map[string]interface{}{
					"name":     tt.args.name,
					"order":    *tt.args.opts.Order,
					"color":    *tt.args.opts.Color,
					"favorite": *tt.args.opts.Favorite,
				},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			label, err := cl.CreateLabelWithOptions(tt.args.name, tt.args.opts)

			assert.Equal(t, tt.want, label)
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

func TestClient_UpdateLabelWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *UpdateLabelOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &UpdateLabelOptions{
				RequestID: String("REQUEST_ID"),
				Name:      String("NAME"),
				Order:     Int(1),
				Color:     Int(2),
				Favorite:  Bool(true),
			}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &UpdateLabelOptions{
				RequestID: String("REQUEST_ID"),
				Name:      String("NAME"),
				Order:     Int(1),
				Color:     Int(2),
				Favorite:  Bool(true),
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
				URL:    fmt.Sprintf("https://api.todoist.com/rest/v1/labels/%d", tt.args.id),
				Method: http.MethodPost,
				Payload: map[string]interface{}{
					"name":     *tt.args.opts.Name,
					"order":    *tt.args.opts.Order,
					"color":    *tt.args.opts.Color,
					"favorite": *tt.args.opts.Favorite,
				},
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "Content-Type": "application/json", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.UpdateLabelWithOptions(tt.args.id, tt.args.opts)

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

func TestClient_DeleteLabel(t *testing.T) {
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/labels/%d", tt.args.id),
				Method:  http.MethodDelete,
				Headers: map[string]string{"Authorization": "Bearer TOKEN"},
			}).Return(tt.resp, nil)

			err := cl.DeleteLabel(tt.args.id)

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

func TestClient_DeleteLabelWithOptions(t *testing.T) {
	type args struct {
		id   int
		opts *DeleteLabelOptions
	}
	tests := []struct {
		name    string
		args    args
		resp    *restResponse
		wantErr bool
	}{
		{
			name: "should return nil",
			args: args{id: 1, opts: &DeleteLabelOptions{
				RequestID: String("REQUEST_ID"),
			}},
			resp: &restResponse{
				StatusCode: http.StatusNoContent,
				Body:       strings.NewReader(""),
			},
			wantErr: false,
		},
		{
			name: "should return an error if the request fails",
			args: args{id: 1, opts: &DeleteLabelOptions{
				RequestID: String("REQUEST_ID"),
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
				URL:     fmt.Sprintf("https://api.todoist.com/rest/v1/labels/%d", tt.args.id),
				Method:  http.MethodDelete,
				Headers: map[string]string{"Authorization": "Bearer TOKEN", "X-Request-Id": *tt.args.opts.RequestID},
			}).Return(tt.resp, nil)

			err := cl.DeleteLabelWithOptions(tt.args.id, tt.args.opts)

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
