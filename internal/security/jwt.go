package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken() *string {
	claims := jwt.MapClaims{
		"email": "jdoe@gmail.com",
		"exp":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token_string, err := token.SignedString("4af77f338270ae9df3c941744ffcc01bd2b3253bbbe8f2b6beae213a9ee3defa")
	if err != nil {
		return nil
	}
	return &token_string
}
