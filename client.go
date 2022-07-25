package todoist

type Client struct {
	token string
}

func New(token string) *Client {
	return &Client{token}
}
