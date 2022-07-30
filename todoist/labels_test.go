package todoist

import (
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
			} else {
				assert.NoError(t, err)
			}
			api.AssertExpectations(t)
		})
	}
}
