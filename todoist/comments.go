package todoist

import (
	"fmt"
	"strconv"
)

type Comment struct {
	// Comment ID.
	ID int `json:"id"`
	// Comment's task ID (for task comments).
	TaskID *int `json:"task_id"`
	// Comment's project ID (for project comments).
	ProjectID *int `json:"project_id"`
	// Date and time when comment was added, RFC3339 (https://www.ietf.org/rfc/rfc3339.txt) format in UTC.
	Posted string `json:"posted"`
	// Comment content.
	// This value may contain markdown-formatted text and hyperlinks.
	// Details on markdown support can be found in the Text Formatting article in the Help Center.
	Content string `json:"content"`
	// Attachment file.
	Attachment *Attachment `json:"attachment"`
}

// List of comments.
type Comments []*Comment

type Attachment struct {
	// The type of the file (for example image, video, audio, file, etc.)
	ResourceType string `json:"resource_type"`
	// The name of the file.
	FileName *string `json:"file_name"`
	// The size of the file in bytes.
	FileSize *int `json:"file_size"`
	// MIME type (i.e. text/plain, image/png).
	FileType *string `json:"file_type"`
	// The URL where the file is located (a string value representing an HTTP URL).
	// Note that we don't cache the remote content on our servers and stream or expose files directly from third party resources.
	// In particular this means that you should avoid providing links to non-encrypted (plain HTTP) resources, as exposing this files in Todoist may issue a browser warning.
	FileURL *string `json:"file_url"`
	// If you upload an audio file, you may provide an extra attribute file_duration (duration of the audio file in seconds, which takes an integer value).
	FileDuration *int `json:"file_duration"`
	// Upload completion state (either pending or completed).
	UploadState *string `json:"upload_state"`
	// Image file URL.
	Image *string `json:"image"`
	// Image width.
	ImageWidth *int `json:"image_width"`
	// Image height.
	ImageHeight *int `json:"image_height"`
	// Large thumbnail (a list that contains the URL, the width and the height of the thumbnail).
	TnL []interface{} `json:"tn_l"`
	// Medium thumbnail (a list that contains the URL, the width and the height of the thumbnail).
	TnM []interface{} `json:"tn_m"`
	// Small thumbnail (a list that contains the URL, the width and the height of the thumbnail).
	TnS []interface{} `json:"tn_s"`
}

// Gets list of all comments for a project.
func (cl *Client) GetProjectComments(projectID int) (Comments, error) {
	cmts := Comments{}
	if err := cl.get("/v1/comments", map[string]string{"project_id": strconv.Itoa(projectID)}, &cmts); err != nil {
		return nil, err
	}
	return cmts, nil
}

// Gets list of all comments for a task.
func (cl *Client) GetTaskComments(taskID int) (Comments, error) {
	cmts := Comments{}
	if err := cl.get("/v1/comments", map[string]string{"task_id": strconv.Itoa(taskID)}, &cmts); err != nil {
		return nil, err
	}
	return cmts, nil
}

// Gets a comment.
func (cl *Client) GetComment(id int) (*Comment, error) {
	cmt := Comment{}
	if err := cl.get(fmt.Sprintf("/v1/comments/%d", id), nil, &cmt); err != nil {
		return nil, err
	}

	return &cmt, nil
}

// Options for creating attachment.
type UploadAttachmentOptions struct {
	ResourceType *string
	FileName     *string
	FileURL      *string
	FileType     *string
}

// Options for creating a comment for a project.
type CreateProjectCommentOptions struct {
	RequestID *string

	// Object for attachment object.
	Attachment *UploadAttachmentOptions
}

// Creates a comment for a project.
func (cl *Client) CreateProjectComment(projectID int, content string) (*Comment, error) {
	return cl.CreateProjectCommentWithOptions(projectID, content, nil)
}

// Creates a comment for a project with options.
func (cl *Client) CreateProjectCommentWithOptions(projectID int, content string, opts *CreateProjectCommentOptions) (*Comment, error) {
	p := map[string]interface{}{"project_id": projectID, "content": content}
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
		if opts.Attachment != nil {
			a := map[string]string{}
			addOptionalStringToStringMap(a, "resource_type", opts.Attachment.ResourceType)
			addOptionalStringToStringMap(a, "file_name", opts.Attachment.FileName)
			addOptionalStringToStringMap(a, "file_url", opts.Attachment.FileURL)
			addOptionalStringToStringMap(a, "file_type", opts.Attachment.FileType)
			p["attachment"] = a
		}
	}

	cmt := Comment{}
	if err := cl.post("/v1/comments", p, reqID, &cmt); err != nil {
		return nil, err
	}

	return &cmt, nil
}

// Options for creating a comment for a task.
type CreateTaskCommentOptions struct {
	RequestID *string

	// Object for attachment object.
	Attachment *UploadAttachmentOptions
}

// Creates a comment for a task.
func (cl *Client) CreateTaskComment(taskID int, content string) (*Comment, error) {
	return cl.CreateTaskCommentWithOptions(taskID, content, nil)
}

// Creates a comment for a task with options.
func (cl *Client) CreateTaskCommentWithOptions(taskID int, content string, opts *CreateTaskCommentOptions) (*Comment, error) {
	p := map[string]interface{}{"task_id": taskID, "content": content}
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
		if opts.Attachment != nil {
			a := map[string]string{}
			addOptionalStringToStringMap(a, "resource_type", opts.Attachment.ResourceType)
			addOptionalStringToStringMap(a, "file_name", opts.Attachment.FileName)
			addOptionalStringToStringMap(a, "file_url", opts.Attachment.FileURL)
			addOptionalStringToStringMap(a, "file_type", opts.Attachment.FileType)
			p["attachment"] = a
		}
	}

	cmt := Comment{}
	if err := cl.post("/v1/comments", p, reqID, &cmt); err != nil {
		return nil, err
	}

	return &cmt, nil
}

// Options for updating a comment.
type UpdateCommentOptions struct {
	RequestID *string
}

// Updates a comment.
func (cl *Client) UpdateComment(id int, content string) error {
	return cl.UpdateCommentWithOptions(id, content, nil)
}

// Updates a comment with options.
func (cl *Client) UpdateCommentWithOptions(id int, content string, opts *UpdateCommentOptions) error {
	p := map[string]interface{}{"content": content}
	var reqID *string
	if opts != nil {
		reqID = opts.RequestID
	}

	if err := cl.postWithoutBind(fmt.Sprintf("/v1/comments/%d", id), p, reqID); err != nil {
		return err
	}

	return nil
}
