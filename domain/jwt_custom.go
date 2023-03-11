package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}
