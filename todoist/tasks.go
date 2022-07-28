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
