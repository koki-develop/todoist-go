package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Project.
type Project struct {
	// Project ID.
	ID int `json:"id"`
	// Project name.
	Name string `json:"name"`
	// A numeric ID representing the color of the project icon.
	// Refer to the id column in the Colors guide (https://developer.todoist.com/guides/#colors) for more info.
	Color int `json:"color"`
	// ID of parent project (read-only, absent for top-level projects).
	ParentID *int `json:"parent_id"`
	// Project position under the same parent (read-only).
	Order int `json:"order"`
	// Number of project comments.
	CommentCount int `json:"comment_count"`
	// Whether the project is shared (read-only, a true or false value)
	Shared bool `json:"shared"`
	// Whether the project is a favorite (a true or false value).
	Favorite bool `json:"favorite"`
	// Whether the project is Inbox (read-only, true or otherwise this property is not sent).
	InboxProject bool `json:"inbox_project"`
	// Whether the project is TeamInbox (read-only, true or otherwise this property is not sent).
	TeamInbox bool `json:"team_inbox"`
	// Identifier to find the match between different copies of shared projects.
	// When you share a project, its copy has a different ID for your collaborators.
	// To find a project in a different account that matches yours, you can use the "sync_id" attribute.
	// For non-shared projects the attribute is set to 0.
	SyncID int `json:"sync_id"`
	// URL to access this project in the Todoist web or mobile applications.
	URL string `json:"url"`
}

// List of Projects.
type Projects []*Project

// Returns slice containing all user projects.
func (cl *Client) GetProjects() (Projects, error) {
	ep := "https://api.todoist.com/rest/v1/projects"
	req, err := http.NewRequest(http.MethodGet, ep, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}

	projs := Projects{}
	if err := json.NewDecoder(resp.Body).Decode(&projs); err != nil {
		return nil, err
	}

	return projs, nil
}

func (cl *Client) CreateProject(name string) (*Project, error) {
	return cl.CreateProjectWithOptions(name, nil)
}

type CreateProjectOptions struct {
	RequestID *string
	ParentID  *int
	Color     *int
	Favorite  *bool
}

func (cl *Client) CreateProjectWithOptions(name string, opts *CreateProjectOptions) (*Project, error) {
	ep := "https://api.todoist.com/rest/v1/projects"
	j := map[string]interface{}{"name": name}
	if opts != nil {
		if opts.ParentID != nil {
			j["parent_id"] = *opts.ParentID
		}
		if opts.Color != nil {
			j["color"] = *opts.Color
		}
		if opts.Favorite != nil {
			j["favorite"] = *opts.Favorite
		}
	}
	p, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, ep, bytes.NewBuffer(p))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))
	req.Header.Set("Content-Type", "application/json")
	if opts != nil && opts.RequestID != nil {
		req.Header.Set("X-Request-Id", *opts.RequestID)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}

	var proj Project
	if err := json.NewDecoder(resp.Body).Decode(&proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

type UpdateProjectOptions struct {
	RequestID *string
	Name      *string
	Color     *int
	Favorite  *bool
}

func (cl *Client) UpdateProject(id int, opts *UpdateProjectOptions) error {
	ep := fmt.Sprintf("https://api.todoist.com/rest/v1/projects/%d", id)
	j := map[string]interface{}{}
	if opts.Name != nil {
		j["name"] = *opts.Name
	}
	if opts.Color != nil {
		j["color"] = *opts.Color
	}
	if opts.Favorite != nil {
		j["favorite"] = *opts.Favorite
	}
	p, err := json.Marshal(j)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, ep, bytes.NewBuffer(p))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))
	req.Header.Set("Content-Type", "application/json")
	if opts != nil && opts.RequestID != nil {
		req.Header.Set("X-Request-Id", *opts.RequestID)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(b))
	}

	return nil
}
