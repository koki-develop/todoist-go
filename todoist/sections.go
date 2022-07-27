package todoist

import (
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
