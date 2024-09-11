package utils

import (
	"errors"
	"fmt"
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
		return "", errors.New("can't create a JWT token, cause 'JWT_SECRET' in .env is missing ❌")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tString, nil
}

func VerifyJWTToken(tString string) error {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return errors.New("can't verify a JWT token, cause 'JWT_SECRET' in .env is missing ❌")
	}

	// Parse the JWT token
	token, err := jwt.Parse(tString, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return errors.New("can't verify a JWT token, error during parsing ❌")
	}

	// Verify token is valid
	if !token.Valid {
		return errors.New("JWT token is invalid ❌")
	}

	return nil
}

func ParseClaimsToken(tString string) (primitive.ObjectID, error) {
	t, _ := jwt.Parse(tString, nil)
	c, _ := t.Claims.(jwt.MapClaims)
	id := c["id"].(string)
	fmt.Println(id)

	return primitive.ObjectIDFromHex(id)
}
