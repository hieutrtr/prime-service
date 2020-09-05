package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(iss string, JWTSecret string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = iss
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	t, _ := token.SignedString([]byte(JWTSecret))
	return "Bearer "+ t
}
