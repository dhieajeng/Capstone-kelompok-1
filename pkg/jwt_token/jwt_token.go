package jwt_token

import (
	"github.com/golang-jwt/jwt/v5"
)

type jwtToken struct {
	secretKey string
}

type JwtCustomClaims struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Avatar string `json:"avatar"`
	jwt.RegisteredClaims
}

func (t *jwtToken) GenerateAccessToken(claims JwtCustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(t.secretKey))

	if err != nil {
		return "", err
	}

	return encodedToken, nil
}

type JwtToken interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
}

func NewJwtToken(secretKey string) JwtToken {
	return &jwtToken{
		secretKey: secretKey,
	}
}
