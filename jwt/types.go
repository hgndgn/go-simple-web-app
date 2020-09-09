package jwt

import jwt "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JwtOptions struct {
	Secret    []byte
	ExpiresAt int64
}
