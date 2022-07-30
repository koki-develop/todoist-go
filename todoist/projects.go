package todoist

import (
	"fmt"
)

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

// Gets list of all user projects.
func (cl *Client) GetProjects() (Projects, error) {
	projs := Projects{}
	if err := cl.get("/v1/projects", nil, &projs); err != nil {
		return nil, err
	}

	return projs, nil
}

// Gets a project.
func (cl *Client) GetProject(id int) (*Project, error) {
	proj := Project{}
	if err := cl.get(fmt.Sprintf("/v1/projects/%d", id), nil, &proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

// Options for creating a project.
type CreateProjectOptions struct {
	RequestID *string

	// Parent project ID.
	ParentID *int
	// A numeric ID representing the color of the project icon.
	// Refer to the id column in the Colors guide (https://developer.todoist.com/guides/#colors) for more info.
	Color *int
	// Whether the project is a favorite (a true or false value).
	Favorite *bool
}

// Creates a new project and returns it.
func (cl *Client) CreateProject(name string) (*Project, error) {
	return cl.CreateProjectWithOptions(name, nil)
}

// Creates a new project with options and returns it.
func (cl *Client) CreateProjectWithOptions(name string, opts *CreateProjectOptions) (*Project, error) {
	p := map[string]interface{}{"name": name}
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
		addOptionalIntToMap(p, "parent_id", opts.ParentID)
		addOptionalIntToMap(p, "color", opts.Color)
		addOptionalBoolToMap(p, "favorite", opts.Favorite)
	}

	proj := Project{}
	if err := cl.post("/v1/projects", p, reqID, &proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

// Options for updating a project.
type UpdateProjectOptions struct {
	RequestID *string

	// Name of the project.
	Name *string
	// A numeric ID representing the color of the project icon.
	// Refer to the id column in the Colors guide (https://developer.todoist.com/guides/#colors) for more info.
	Color *int
	// Whether the project is a favorite (a true or false value).
	Favorite *bool
}

// Updates a project.
func (cl *Client) UpdateProjectWithOptions(id int, opts *UpdateProjectOptions) error {
	p := map[string]interface{}{}
	var reqID *string = nil
	if opts != nil {
		reqID = opts.RequestID
		addOptionalStringToMap(p, "name", opts.Name)
		addOptionalIntToMap(p, "color", opts.Color)
		addOptionalBoolToMap(p, "favorite", opts.Favorite)
	}

	if err := cl.postWithoutBind(fmt.Sprintf("/v1/projects/%d", id), p, reqID); err != nil {
		return err
	}

	return nil
}

// Options for deleting a project.
type DeleteProjectOptions struct {
	RequestID *string
}

// Deletes a project.
func (cl *Client) DeleteProject(id int) error {
	return cl.DeleteProjectWithOptions(id, nil)
}

// Deletes a project with options.
func (cl *Client) DeleteProjectWithOptions(id int, opts *DeleteProjectOptions) error {
	var reqID *string = nil
	if opts != nil {
		reqID = opts.RequestID
	}

	if err := cl.delete(fmt.Sprintf("/v1/projects/%d", id), reqID); err != nil {
		return err
	}

	return nil
}

// Get list of all collaborators of a shared project.
func (cl *Client) GetCollaborators(projectID int) (Users, error) {
	users := Users{}
	if err := cl.get(fmt.Sprintf("/v1/projects/%d/collaborators", projectID), nil, &users); err != nil {
		return nil, err
	}

	return users, nil
}
