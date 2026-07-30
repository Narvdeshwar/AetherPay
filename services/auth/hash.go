package main

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes plain password using Bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

