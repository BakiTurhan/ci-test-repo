package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email string `validate:"required,email"`
}

func main() {
	// Validator'ı kullanalım
	validate := validator.New()
	user := User{Email: "test@test.com"}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("Validation error:", err)
	}

	// Crypto'yu kullanalım
	password := []byte("gizlisifre")
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	fmt.Println("Hashed password:", string(hashedPassword))
}
