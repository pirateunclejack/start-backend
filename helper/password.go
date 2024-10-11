package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        log.Fatalf("faield to generate hashed password: %v\n", err)
    }
    return string(bytes)
}

func VerifyPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err!= nil {
        log.Println("failed to verify hashed password: ", err)
        return false
    }
    return true
}
