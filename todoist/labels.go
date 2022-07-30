package todoist

type Label struct {
	// Label ID.
	ID int `json:"id"`
	// Label name.
	Name string `json:"name"`
	// A numeric ID representing the color of the label icon. Refer to the id column in the Colors (https://developer.todoist.com/guides/#colors) guide for more info.
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
