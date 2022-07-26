package todoist

// User.
type User struct {
	// User ID.
	ID int `json:"id"`
	// User name.
	Name string `json:"name"`
	// User email address.
	Email string `json:"email"`
}

// List of users.
type Users []*User
