package models

type Permission struct {
	Permission string `json:"permission"`
	User       User   `json:"user"`
}
