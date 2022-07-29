package todoist

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
