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
	RequestID *string `json:"-"`

	// Label order.
	Order *int `json:"order,omitempty"`
	// A numeric ID representing the color of the label icon.
	// Refer to the id column in the Colors (https://developer.todoist.com/guides/#colors) guide for more info.
	Color *int `json:"color,omitempty"`
	// Whether the label is a favorite (a true or false value).
	Favorite *bool `json:"favorite,omitempty"`
}

// Creates a label.
func (cl *Client) CreateLabel(name string) (*Label, error) {
	return cl.CreateLabelWithOptions(name, nil)
}

// Creates a label with options.
func (cl *Client) CreateLabelWithOptions(name string, opts *CreateLabelOptions) (*Label, error) {
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	p := map[string]interface{}{"name": name}
	if err := toMap(opts, p); err != nil {
		return nil, err
	}

	label := Label{}
	if err := cl.post("/v1/labels", p, reqID, &label); err != nil {
		return nil, err
	}

	return &label, nil
}

// Options for updating a label.
type UpdateLabelOptions struct {
	RequestID *string `json:"-"`

	// New name of the label.
	Name *string `json:"name,omitempty"`
	// Number that is used by clients to sort list of labels.
	Order *int `json:"order,omitempty"`
	//	A numeric ID representing the color of the label icon.
	// Refer to the id column in the Colors (https://developer.todoist.com/guides/#colors) guide for more info.
	Color *int `json:"color,omitempty"`
	// Whether the label is a favorite (a true or false value).
	Favorite *bool `json:"favorite,omitempty"`
}

// Updates a label with options.
func (cl *Client) UpdateLabelWithOptions(id int, opts *UpdateLabelOptions) error {
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	p := map[string]interface{}{}
	if err := toMap(opts, p); err != nil {
		return err
	}

	if err := cl.postWithoutBind(fmt.Sprintf("/v1/labels/%d", id), p, reqID); err != nil {
		return err
	}

	return nil
}

// Options for deleting a label.
type DeleteLabelOptions struct {
	RequestID *string `json:"-"`
}

// Deletes a label.
func (cl *Client) DeleteLabel(id int) error {
	return cl.DeleteLabelWithOptions(id, nil)
}

// Deletes a label with options.
func (cl *Client) DeleteLabelWithOptions(id int, opts *DeleteLabelOptions) error {
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	if err := cl.delete(fmt.Sprintf("/v1/labels/%d", id), reqID); err != nil {
		return err
	}
	return nil
}
