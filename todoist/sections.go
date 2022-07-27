package todoist

import (
	"fmt"
	"strconv"
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

// Options for getting sections.
type GetSectionsOptions struct {
	ProjectID *int
}

// Get list of all sections.
func (cl *Client) GetSections() (Sections, error) {
	return cl.GetSectionsWithOptions(nil)
}

// Get list of all sections with options.
func (cl *Client) GetSectionsWithOptions(opts *GetSectionsOptions) (Sections, error) {
	p := map[string]string{}
	if opts != nil {
		if opts.ProjectID != nil {
			p["project_id"] = strconv.Itoa(*opts.ProjectID)
		}
	}

	secs := Sections{}
	if err := cl.get("/v1/sections", p, &secs); err != nil {
		return nil, err
	}

	return secs, nil
}

// Gets the section related to the given ID.
func (cl *Client) GetSection(id int) (*Section, error) {
	sec := Section{}
	if err := cl.get(fmt.Sprintf("/v1/sections/%d", id), nil, &sec); err != nil {
		return nil, err
	}
	return &sec, nil
}

// Options for creating section.
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
	j := map[string]interface{}{"name": name, "project_id": projectID}
	var reqID *string
	if opts != nil {
		addOptionalIntToMap(j, "order", opts.Order)
		reqID = opts.RequestID
	}

	sec := Section{}
	if err := cl.post("/v1/sections", j, reqID, &sec); err != nil {
		return nil, err
	}

	return &sec, nil
}

// Options for updating section.
type UpdateSectionOptions struct {
	RequestID *string
}

// Updates the section for the given ID.
func (cl *Client) UpdateSection(id int, name string) error {
	return cl.UpdateSectionWithOptions(id, name, nil)
}

// Updates the section for the given ID with options.
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

// Deletes the section for the given ID.
func (cl *Client) DeleteSection(id int) error {
	return cl.DeleteSectionWithOptions(id, nil)
}

// Deletes the section for the given ID with options.
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
