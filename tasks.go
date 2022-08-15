package todoist

import (
	"fmt"
)

type Task struct {
	// Task ID.
	ID int `json:"id"`
	// Task's project ID (read-only).
	ProjectID int `json:"project_id"`
	// ID of section task belongs to.
	SectionID int `json:"section_id"`
	// Task content.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article (https://todoist.com/help/articles/text-formatting) in the Help Center.
	Content string `json:"content"`
	// A description for the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article (https://todoist.com/help/articles/text-formatting) in the Help Center.
	Description string `json:"description"`
	// Flag to mark completed tasks.
	Completed bool `json:"completed"`
	// Array of label IDs, associated with a task.
	LabelIDs []int `json:"label_ids"`
	// ID of parent task (read-only, absent for top-level tasks).
	ParentID *int `json:"parent_id"`
	// Position under the same parent or project for top-level tasks (read-only).
	Order int `json:"order"`
	// Task priority from 1 (normal, default value) to 4 (urgent).
	Priority int `json:"priority"`
	// object representing task due date/time.
	Due *Due `json:"due"`
	// URL to access this task in the Todoist web or mobile applications.
	URL string `json:"url"`
	// Number of task comments.
	CommentCount int `json:"comment_count"`
	// The responsible user ID (if set, and only for shared tasks).
	Assignee *int `json:"assignee"`
	// The ID of the user who assigned the task. 0 if the task is unassigned.
	Assigner int `json:"assigner"`
}

// List of tasks.
type Tasks []*Task

type Due struct {
	// Human defined date in arbitrary format.
	String string `json:"string"`
	// Date in format YYYY-MM-DD corrected to user's timezone.
	Date string `json:"date"`
	// Whether the task has a recurring due date (https://todoist.com/help/articles/set-a-recurring-due-date).
	Recurring bool `json:"recurring"`
	// Only returned if exact due time set (i.e. it's not a whole-day task), date and time in RFC3339 (https://www.ietf.org/rfc/rfc3339.txt) format in UTC.
	Datetime *string `json:"datetime"`
	// Only returned if exact due time set, user's timezone definition either in tzdata-compatible format ("Europe/Berlin") or as a string specifying east of UTC offset as "UTC±HH:MM" (i.e. "UTC-01:00").
	Timezone *string `json:"timezone"`
}

// Options for getting a tasks.
type GetTasksOptions struct {
	// Filter tasks by project ID.
	ProjectID *int `url:"project_id,omitempty"`
	// Filter tasks by section ID.
	SectionID *int `url:"section_id,omitempty"`
	// Filter tasks by label.
	LabelID *int `url:"label_id,omitempty"`
	// Filter by any supported filter (https://todoist.com/help/articles/205248842).
	Filter *string `url:"filter,omitempty"`
	// IETF language tag defining what language filter is written in, if differs from default English.
	Lang *string `url:"lang,omitempty"`
	// A list of the task IDs to retrieve, this should be a comma separated list.
	IDs *[]int `url:"ids,comma,omitempty"`
}

// Gets list of all active tasks.
func (cl *Client) GetTasks() (Tasks, error) {
	return cl.GetTasksWithOptions(nil)
}

// Gets list of all active tasks with options.
func (cl *Client) GetTasksWithOptions(opts *GetTasksOptions) (Tasks, error) {
	tasks := Tasks{}
	if err := cl.get("/v1/tasks", opts, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Get a single active task.
func (cl *Client) GetTask(id int) (*Task, error) {
	task := Task{}
	if err := cl.get(fmt.Sprintf("/v1/tasks/%d", id), nil, &task); err != nil {
		return nil, err
	}

	return &task, nil
}

// Options for creating a task.
type CreateTaskOptions struct {
	RequestID *string `json:"-"`

	// A description for the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article (https://todoist.com/help/articles/text-formatting) in the Help Center.
	Description *string `json:"description,omitempty"`
	// Task project ID.
	// If not set, task is put to user's Inbox.
	ProjectID *int `json:"project_id,omitempty"`
	// ID of section to put task into.
	SectionID *int `json:"section_id,omitempty"`
	// Parent task ID.
	ParentID *int `json:"parent_id,omitempty"`
	// Non-zero integer value used by clients to sort tasks under the same parent.
	Order *int `json:"order,omitempty"`
	// IDs of labels associated with the task.
	LabelIDs *[]int `json:"label_ids,omitempty"`
	// Task priority from 1 (normal) to 4 (urgent).
	Priority *int `json:"priority,omitempty"`
	// Human defined (https://todoist.com/help/articles/due-dates-and-times) task due date (ex.: "next Monday", "Tomorrow"). Value is set using local (not UTC) time.
	DueString *string `json:"due_string,omitempty"`
	// Specific date in YYYY-MM-DD format relative to user’s timezone.
	DueDate *string `json:"due_date,omitempty"`
	// Specific date and time in RFC3339 (https://www.ietf.org/rfc/rfc3339.txt) format in UTC.
	DueDatetime *string `json:"due_datetime,omitempty"`
	// 2-letter code specifying language in case due_string is not written in English.
	DueLang *string `json:"due_lang,omitempty"`
	// The responsible user ID (if set, and only for shared tasks).
	Assignee *int `json:"assignee,omitempty"`
}

// Creates a new task and returns it.
func (cl *Client) CreateTask(content string) (*Task, error) {
	return cl.CreateTaskWithOptions(content, nil)
}

// Creates a new task with options and returns it.
func (cl *Client) CreateTaskWithOptions(content string, opts *CreateTaskOptions) (*Task, error) {
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	p := map[string]interface{}{"content": content}
	if err := toMap(opts, p); err != nil {
		return nil, err
	}

	task := Task{}
	if err := cl.post("/v1/tasks", p, reqID, &task); err != nil {
		return nil, err
	}

	return &task, nil
}

// Options for updating a task.
type UpdateTaskOptions struct {
	RequestID *string

	// Task content.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article (https://todoist.com/help/articles/text-formatting) in the Help Center.
	Content *string
	// A description for the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article (https://todoist.com/help/articles/text-formatting) in the Help Center.
	Description *string
	// IDs of labels associated with the task.
	LabelIDs *[]int
	// Task priority from 1 (normal) to 4 (urgent).
	Priority *int
	// Human defined (https://todoist.com/help/articles/due-dates-and-times) task due date (ex.: "next Monday", "Tomorrow"). Value is set using local (not UTC) time.
	DueString *string
	// Specific date in YYYY-MM-DD format relative to user’s timezone.
	DueDate *string
	// Specific date and time in RFC3339 (https://www.ietf.org/rfc/rfc3339.txt) format in UTC.
	DueDatetime *string
	// 2-letter code specifying language in case due_string is not written in English.
	DueLang *string
	// The responsible user ID or 0 to unset (for shared tasks).
	Assignee *int
}

// Updates a task.
func (cl *Client) UpdateTaskWithOptions(id int, opts *UpdateTaskOptions) error {
	var reqID *string
	p := map[string]interface{}{}
	if opts != nil {
		reqID = opts.RequestID
		addOptionalStringToMap(p, "content", opts.Content)
		addOptionalStringToMap(p, "description", opts.Description)
		if opts.LabelIDs != nil {
			p["label_ids"] = opts.LabelIDs
		}
		addOptionalIntToMap(p, "priority", opts.Priority)
		addOptionalStringToMap(p, "due_string", opts.DueString)
		addOptionalStringToMap(p, "due_date", opts.DueDate)
		addOptionalStringToMap(p, "due_datetime", opts.DueDatetime)
		addOptionalStringToMap(p, "due_lang", opts.DueLang)
		addOptionalIntToMap(p, "assignee", opts.Assignee)
	}

	if err := cl.postWithoutBind(fmt.Sprintf("/v1/tasks/%d", id), p, reqID); err != nil {
		return err
	}

	return nil
}

// Options for closing a task.
type CloseTaskOptions struct {
	RequestID *string
}

// Closes a task.
func (cl *Client) CloseTask(id int) error {
	return cl.CloseTaskWithOptions(id, nil)
}

// Closes a task with options.
func (cl *Client) CloseTaskWithOptions(id int, opts *CloseTaskOptions) error {
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	if err := cl.postWithoutBind(fmt.Sprintf("/v1/tasks/%d/close", id), nil, reqID); err != nil {
		return err
	}

	return nil
}

// Options for reopening a task.
type ReopenTaskOptions struct {
	RequestID *string
}

// Reopens a task.
func (cl *Client) ReopenTask(id int) error {
	return cl.ReopenTaskWithOptions(id, nil)
}

// Reopens a task with options.
func (cl *Client) ReopenTaskWithOptions(id int, opts *ReopenTaskOptions) error {
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	if err := cl.postWithoutBind(fmt.Sprintf("/v1/tasks/%d/reopen", id), nil, reqID); err != nil {
		return err
	}

	return nil
}

// Options for deleting a task.
type DeleteTaskOptions struct {
	RequestID *string
}

// Deletes a task.
func (cl *Client) DeleteTask(id int) error {
	return cl.DeleteTaskWithOptions(id, nil)
}

// Deletes a task with options.
func (cl *Client) DeleteTaskWithOptions(id int, opts *DeleteTaskOptions) error {
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	if err := cl.delete(fmt.Sprintf("/v1/tasks/%d", id), reqID); err != nil {
		return err
	}

	return nil
}
