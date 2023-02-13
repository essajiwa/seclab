package model

// create a struct for the user
type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password,omitempty" db:"password"`
	Role     string `json:"role" db:"role"`
}
