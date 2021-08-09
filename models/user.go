package models

type User struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
