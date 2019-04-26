package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SignedString is secret key for decoding
var SignedString = []byte("secret-string")

var Expire = time.Duration(5)

// InitializeTokenService take the JWT parameters from Environment and
// does what it need
func InitializeTokenService(e *Environment) {
	if e.JwtSecret != "" {
		SignedString = []byte(e.JwtSecret)
	}
	if e.JwtExpire != 0 {
		Expire = time.Duration(e.JwtExpire)
	}
}

// CreateToken takes email and password and gives valid JWT
func CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().UTC().Add(time.Minute * Expire).Unix(),
	})
	tokenString, err := token.SignedString(SignedString)
	if err != nil {
		fmt.Printf("%s", err)
		return tokenString, err
	}
	return tokenString, err
}
