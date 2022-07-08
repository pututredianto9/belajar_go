package jwtAuth

import "github.com/dgrijalva/jwt-go"

type JwtData struct {
	UserID   string `json:"uid"`
	UserName string `json:"uname"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
