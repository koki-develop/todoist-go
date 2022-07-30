package todoist

func newClientForTest() (*Client, *mockRestAPI) {
	api := &mockRestAPI{}
	return &Client{token: "TOKEN", restAPI: api}, api
}
