package models

import "github.com/golang-jwt/jwt"

type User struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
