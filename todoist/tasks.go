package todoist

type Task struct {
	// Task ID.
	ID int
	// Task's project ID (read-only).
	ProjectID int
	// ID of section task belongs to.
	SectionID int
	// Task content.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article (https://todoist.com/help/articles/text-formatting) in the Help Center.
	Content string
	// A description for the task.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article (https://todoist.com/help/articles/text-formatting) in the Help Center.
	Description string
	// Flag to mark completed tasks.
	Completed bool
	// Array of label IDs, associated with a task.
	LabelIDs []int
	// ID of parent task (read-only, absent for top-level tasks).
	ParentID *int
	// Position under the same parent or project for top-level tasks (read-only).
	Order int
	// Task priority from 1 (normal, default value) to 4 (urgent).
	Priority int
	// object representing task due date/time.
	Due *Due
	// URL to access this task in the Todoist web or mobile applications.
	URL string
	// Number of task comments.
	CommentCount int
	// The responsible user ID (if set, and only for shared tasks).
	Assignee *int
	// The ID of the user who assigned the task. 0 if the task is unassigned.
	Assigner int
}

type Tasks []*Task

type Due struct {
	// Human defined date in arbitrary format.
	String string
	// Date in format YYYY-MM-DD corrected to user's timezone.
	Date string
	// Whether the task has a recurring due date (https://todoist.com/help/articles/set-a-recurring-due-date).
	Recurring bool
	// Only returned if exact due time set (i.e. it's not a whole-day task), date and time in RFC3339 (https://www.ietf.org/rfc/rfc3339.txt) format in UTC.
	Datetime *string
	// Only returned if exact due time set, user's timezone definition either in tzdata-compatible format ("Europe/Berlin") or as a string specifying east of UTC offset as "UTC±HH:MM" (i.e. "UTC-01:00").
	Timezone *string
}
