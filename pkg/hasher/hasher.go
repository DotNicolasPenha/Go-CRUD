package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		10,
	)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func ComparePassword(passwordHash string, passwordInput string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(passwordInput),
	)
	return err == nil
}
