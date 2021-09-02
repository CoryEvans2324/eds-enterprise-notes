package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	ID         uint   `gorm:"primarykey autoIncrement"`
	Permission string `json:"permission"`
	UserID     int    `json:"userID"`
	User       User   `json:"user" gorm:"foreignKey:UserID"`
	NoteID     int    `json:"noteID" gorm:"foreignKey:NoteID"`
	Note       Note   `json:"note"`
}
