package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int, error) {

}
