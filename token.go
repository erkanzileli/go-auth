package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SignedString is secret key for decoding
var SignedString = []byte("secret-string")

// CreateToken takes email and password and gives valid JWT
func CreateToken(email, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expires": time.Now().Add(time.Minute * 5),
	})
	tokenString, err := token.SignedString(SignedString)
	if err != nil {
		fmt.Printf("%s", err)
		return tokenString, err
	}
	return tokenString, err
}
