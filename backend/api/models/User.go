package models

// User - User model
type User struct {
	ID       uint32 `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
