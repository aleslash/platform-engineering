package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// Dummy hash function for illustration purposes
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
