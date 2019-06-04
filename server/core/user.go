package core

// User ...
type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Admin     bool   `json:"admin"`
	Active    bool   `json:"active"`
	Avatar    string `json:"avatar"`
	Created   int64  `json:"created"`
	Updated   int64  `json:"updated"`
	LastLogin int64  `json:"last_login"`
}

// UserStore ...
type UserStore interface {
	// Count returns a count of active users.
	Count() (int64, error)

	// Create persists a new user to the datastore.
	Create(*User) error
}