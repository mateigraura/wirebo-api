package crypto

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	payload, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	return string(payload), err
}

func CheckEqual(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
