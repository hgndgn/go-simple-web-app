package jwt

import (
	"fmt"
	"hgndgn/api/jwt-authentication/config"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtConfig JwtOptions

func Initialize() {
	secret := config.Get("jwt-secret", "")
	if secret == "" {
		panic("'jwt-secret' not set in '.env'")
	}

	jwtExp := config.Get("jwt-exp", "")
	if jwtExp == "" {
		panic("'jwt-secret' not set in '.env'")
	}
	expSeconds, err := strconv.ParseInt(jwtExp, 10, 0)
	if err != nil {
		panic(err)
	}

	JwtConfig.ExpiresAt = time.Now().Add(time.Duration(expSeconds) * time.Second).Unix()
	JwtConfig.Secret = []byte(secret)
}

func TokenString(token *jwt.Token) string {
	tokenString, err := token.SignedString(JwtConfig.Secret)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func CreateToken(username string) *jwt.Token {
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: JwtConfig.ExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token
}

func ParseToken(tokenString string) *jwt.Token {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		var expiresAt = token.Claims.(Claims).ExpiresAt
		var now = time.Now().Unix()
		if now > expiresAt {
			return nil, fmt.Errorf("Token expired")
		}

		if token.Claims.(Claims).Username != "admin" {
			return nil, fmt.Errorf("Invalid user")
		}

		return token, nil
	})

	return token
}

func TokenClaims(token *jwt.Token) Claims {
	return token.Claims.(Claims)
}
