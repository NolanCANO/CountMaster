package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Hasher un mot de passe
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Comparer un mot de passe en clair avec un mot de passe hashé
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
