package models

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primarykey autoIncrement"`
	Username     string `gorm:"unique"`
	PasswordHash string
	Role         string `gorm:"default:standard"`
}

type JWTUser struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
