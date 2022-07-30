package todoist

import "fmt"

type Label struct {
	// Label ID.
	ID int `json:"id"`
	// Label name.
	Name string `json:"name"`
	// A numeric ID representing the color of the label icon.
	// Refer to the id column in the Colors (https://developer.todoist.com/guides/#colors) guide for more info.
	Color int `json:"color"`
	// Number used by clients to sort list of labels.
	Order int `json:"order"`
	// Whether the label is a favorite (a true or false value).
	Favorite bool `json:"favorite"`
}

// List of labels.
type Labels []*Label

// Gets list of all user labels.
func (cl *Client) GetLabels() (Labels, error) {
	labels := Labels{}
	if err := cl.get("/v1/labels", nil, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}

// Gets a label.
func (cl *Client) GetLabel(id int) (*Label, error) {
	label := Label{}
	if err := cl.get(fmt.Sprintf("/v1/labels/%d", id), nil, &label); err != nil {
		return nil, err
	}

	return &label, nil
}

// Options for creating a label.
type CreateLabelOptions struct {
	RequestID *string

	// Label order.
	Order *int
	// A numeric ID representing the color of the label icon.
	// Refer to the id column in the Colors (https://developer.todoist.com/guides/#colors) guide for more info.
	Color *int
	// Whether the label is a favorite (a true or false value).
	Favorite *bool
}

// Creates a label.
func (cl *Client) CreateLabel(name string) (*Label, error) {
	return cl.CreateLabelWithOptions(name, nil)
}

// Creates a label with options.
func (cl *Client) CreateLabelWithOptions(name string, opts *CreateLabelOptions) (*Label, error) {
	p := map[string]interface{}{"name": name}
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
		addOptionalIntToMap(p, "order", opts.Order)
		addOptionalIntToMap(p, "color", opts.Color)
		addOptionalBoolToMap(p, "favorite", opts.Favorite)
	}

	label := Label{}
	if err := cl.post("/v1/labels", p, reqID, &label); err != nil {
		return nil, err
	}

	return &label, nil
}
