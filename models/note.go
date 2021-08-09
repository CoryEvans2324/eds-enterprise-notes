package models

import "time"

type Note struct {
	NoteID        int          `json:"noteID"`
	Name          string       `json:"name"`
	Content       string       `json:"content"`
	Status        string       `json:"status"`
	DueDate       *time.Time   `json:"dueDate"`
	Owner         *User        `json:"owner"`
	DelegatedUser *User        `json:"delegatedUser"`
	SharedUsers   []Permission `json:"sharedUsers"`
}
