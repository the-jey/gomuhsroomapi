package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResponseTokenPayload struct {
	Token string `json:"token"`
}

func CreateJWTToken(id primitive.ObjectID) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("an't creata a JWT token, cause 'JWT_SECRET' in .env is missing ❌")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id.String(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tString, nil
}
