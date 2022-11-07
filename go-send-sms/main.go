package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    password := "secret"
		// hash := "$2a$14$yPHbzBrBrnS05rywflpndO1HomfZxnimuyfS0Vk3cI9gaB.IttA3K"
		// hash := "$2a$14$TYj36PkEe7TmdB7lctTvQOJW.kqs0i1d57EJZqQtMc/K8TMUIc4sG"
		hash := "$2a$14$cmBRI6wbhTGlh5hKyhf63Otp8xSiRuDRKqEqPzRJiZGrDuKS.IcjS"
    // hash, _ := HashPassword(password) // ignore error for the sake of simplicity

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Println("Match:   ", match)
}
