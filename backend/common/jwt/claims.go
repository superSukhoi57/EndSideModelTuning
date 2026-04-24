package jwt

import "github.com/golang-jwt/jwt/v4"

type BaseClaims struct {
	jwt.RegisteredClaims
}

type CustomClaims struct {
	BaseClaims
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
	Username string `json:"username"`
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}
