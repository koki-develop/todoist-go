package todoist

import "strings"

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

// Options for getting tasks.
type GetTasksOptions struct {
	// Filter tasks by project ID.
	ProjectID *int
	// Filter tasks by section ID.
	SectionID *int
	// Filter tasks by label.
	LabelID *int
	// Filter by any supported filter (https://todoist.com/help/articles/205248842).
	Filter *string
	// IETF language tag defining what language filter is written in, if differs from default English.
	Lang *string
	// A list of the task IDs to retrieve, this should be a comma separated list.
	IDs []int
}

// Gets list of all active tasks.
func (cl *Client) GetTasks() (Tasks, error) {
	return cl.GetTasksWithOptions(nil)
}

// Gets list of all active tasks with options.
func (cl *Client) GetTasksWithOptions(opts *GetTasksOptions) (Tasks, error) {
	p := map[string]string{}
	if opts != nil {
		addOptionalIntToStringMap(p, "project_id", opts.ProjectID)
		addOptionalIntToStringMap(p, "section_id", opts.SectionID)
		addOptionalIntToStringMap(p, "label_id", opts.LabelID)
		addOptionalStringToStringMap(p, "filter", opts.Filter)
		addOptionalStringToStringMap(p, "lang", opts.Lang)
		if len(opts.IDs) > 0 {
			ids := strings.Join(intsToStrings(opts.IDs), ",")
			addOptionalStringToStringMap(p, "ids", &ids)
		}
	}

	tasks := Tasks{}
	if err := cl.get("/v1/tasks", p, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Options for creating task.
// Please note that only one of the Due* fields can be used at the same time (DueLang is a special case).
type CreateTaskOptions struct {
	// A description for the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article (https://todoist.com/help/articles/text-formatting) in the Help Center.
	Description *string
	// Task project ID.
	// If not set, task is put to user's Inbox.
	ProjectID *int
	// ID of section to put task into.
	SectionID *int
	// Parent task ID.
	ParentID *int
	// Non-zero integer value used by clients to sort tasks under the same parent.
	Order *int
	// IDs of labels associated with the task.
	LabelIDs []int
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
	// The responsible user ID (if set, and only for shared tasks).
	Assignee *int
}

// Creates a new task and returns it.
func (cl *Client) CreateTask(content string) (*Task, error) {
	return cl.CreateTaskWithOptions(content, nil)
}

// Creates a new task with options and returns it.
func (cl *Client) CreateTaskWithOptions(content string, opts *CreateTaskOptions) (*Task, error) {
	p := map[string]interface{}{"content": content}
	var reqID *string
	if opts != nil {
		addOptionalStringToMap(p, "description", opts.Description)
		addOptionalIntToMap(p, "project_id", opts.ProjectID)
		addOptionalIntToMap(p, "section_id", opts.SectionID)
		addOptionalIntToMap(p, "parent_id", opts.ParentID)
		addOptionalIntToMap(p, "order", opts.Order)
		if len(opts.LabelIDs) > 0 {
			ids := strings.Join(intsToStrings(opts.LabelIDs), ",")
			addOptionalStringToMap(p, "label_ids", &ids)
		}
		addOptionalIntToMap(p, "priority", opts.Priority)
		addOptionalStringToMap(p, "due_string", opts.DueString)
		addOptionalStringToMap(p, "due_date", opts.DueDate)
		addOptionalStringToMap(p, "due_datetime", opts.DueDatetime)
		addOptionalStringToMap(p, "due_lang", opts.DueLang)
		addOptionalIntToMap(p, "assignee", opts.Assignee)
	}

	task := Task{}
	if err := cl.post("/v1/tasks", p, reqID, &task); err != nil {
		return nil, err
	}
	return &task, nil
}
