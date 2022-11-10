package password

import (
	"os"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	saltRound, err := strconv.Atoi(os.Getenv("BCRYPT_SALT_ROUND"))
	if err != nil {
		return "", err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), saltRound)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Compare(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func IsValid(password string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9]{8,}$`)
	return regex.MatchString(password)
}
