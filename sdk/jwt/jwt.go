package jwt

import (
	"os"
	"strconv"
	"time"

	"live-easy-backend/sdk/errors"
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(data interface{}) (string, error) {
	expTime, err := strconv.ParseInt(os.Getenv("JWT_EXPIRED_TIME_SEC"), 10, 64)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"exp":  expTime + time.Now().Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func DecodeToken(token string) (map[string]interface{}, error) {
	decoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, errors.NewWithCode(401, "Invalid token", "HTTPStatusUnauthorized")
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.NewWithCode(500, "Failed to decode token", "HTTPStatusInternalServerError")
	}
	if !decoded.Valid {
		return nil, errors.NewWithCode(401, "Invalid token", "HTTPStatusUnauthorized")
	}

	return claims, nil
}
