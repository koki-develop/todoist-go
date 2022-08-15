package todoist

import (
	"fmt"
)

type Section struct {
	// Section id.
	ID int `json:"id"`
	// ID of the project section belongs to.
	ProjectID int `json:"project_id"`
	// Section position among other sections from the same project.
	Order int `json:"order"`
	// Section name.
	Name string `json:"name"`
}

// List of sections.
type Sections []*Section

// Options for getting a sections.
type GetSectionsOptions struct {
	// Filter sections by project ID.
	ProjectID *int `url:"project_id,omitempty"`
}

// Gets list of all sections.
func (cl *Client) GetSections() (Sections, error) {
	return cl.GetSectionsWithOptions(nil)
}

// Gets list of all sections with options.
func (cl *Client) GetSectionsWithOptions(opts *GetSectionsOptions) (Sections, error) {
	secs := Sections{}
	if err := cl.get("/v1/sections", opts, &secs); err != nil {
		return nil, err
	}

	return secs, nil
}

// Gets a section.
func (cl *Client) GetSection(id int) (*Section, error) {
	sec := Section{}
	if err := cl.get(fmt.Sprintf("/v1/sections/%d", id), nil, &sec); err != nil {
		return nil, err
	}

	return &sec, nil
}

// Options for creating a section.
type CreateSectionOptions struct {
	RequestID *string

	// Order among other sections in a project.
	Order *int
}

// Creates a new section and returns it.
func (cl *Client) CreateSection(name string, projectID int) (*Section, error) {
	return cl.CreateSectionWithOptions(name, projectID, nil)
}

// Creates a new section with options and returns it.
func (cl *Client) CreateSectionWithOptions(name string, projectID int, opts *CreateSectionOptions) (*Section, error) {
	p := map[string]interface{}{"name": name, "project_id": projectID}
	var reqID *string
	if opts != nil {
		addOptionalIntToMap(p, "order", opts.Order)
		reqID = opts.RequestID
	}

	sec := Section{}
	if err := cl.post("/v1/sections", p, reqID, &sec); err != nil {
		return nil, err
	}

	return &sec, nil
}

// Options for updating a section.
type UpdateSectionOptions struct {
	RequestID *string
}

// Updates a section.
func (cl *Client) UpdateSection(id int, name string) error {
	return cl.UpdateSectionWithOptions(id, name, nil)
}

// Updates a section with options.
func (cl *Client) UpdateSectionWithOptions(id int, name string, opts *UpdateSectionOptions) error {
	p := map[string]interface{}{"name": name}
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	if err := cl.postWithoutBind(fmt.Sprintf("/v1/sections/%d", id), p, reqID); err != nil {
		return err
	}

	return nil
}

// Options for deleting a section.
type DeleteSectionOptions struct {
	RequestID *string
}

// Deletes a section.
func (cl *Client) DeleteSection(id int) error {
	return cl.DeleteSectionWithOptions(id, nil)
}

// Deletes a section with options.
func (cl *Client) DeleteSectionWithOptions(id int, opts *DeleteSectionOptions) error {
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	if err := cl.delete(fmt.Sprintf("/v1/sections/%d", id), reqID); err != nil {
		return err
	}

	return nil
}
