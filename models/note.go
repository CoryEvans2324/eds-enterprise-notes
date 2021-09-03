package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID              uint         `gorm:"primarykey autoIncrement"`
	Name            string       `json:"name"`
	Content         string       `json:"content"`
	Status          string       `json:"status"`
	DueDate         *time.Time   `json:"dueDate"`
	OwnerID         *int         `json:"ownerID"`
	Owner           *User        `json:"owner" gorm:"foreignKey:OwnerID"`
	DelegatedUserID *int         `json:"delegatedUserID"`
	DelegatedUser   *User        `json:"delegatedUser" gorm:"foreignKey:DelegatedUserID;default:NULL"`
	SharedUsers     []Permission `json:"sharedUsers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foriegnKey:NoteID"`
}

func (n *Note) FormattedDate() string {
	if n.DueDate == nil {
		return ""
	}

	return n.DueDate.Format("02 Jan")
}

func (n *Note) FormattedTime() string {
	if n.DueDate == nil {
		return ""
	}

	return n.DueDate.Format("3:04pm")
}
