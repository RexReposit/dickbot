package models

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DickSize  int    `json:"dick_size"`
	IsBlocked bool   `json:"is_blocked"`
	IsAdmin   bool   `json:"is_admin"`
}
