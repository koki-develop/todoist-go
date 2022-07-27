package todoist

import (
	"fmt"
	"strconv"
)

// Section.
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
