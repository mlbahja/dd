package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Age        string    `json:"age"`
	// First_name string    `json:"first_name"`
	// Last_name  string    `json:"last_name"`
	// CreatedAt  time.Time `json:"createdat"`
}
