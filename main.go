package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Go'nun kütüphaneleri "kullanılmıyor" diye silmemesi için hepsini sahte olarak çağırıyoruz
	_ = validator.New()
	_ = fiber.New()
	_ = jwt.New(jwt.SigningMethodHS256)
	_ = godotenv.Load()
	_, _ = bcrypt.GenerateFromPassword([]byte("test"), 10)
	_, _ = gorm.Open(postgres.Open("dsn"), &gorm.Config{})

	fmt.Println("Tüm kütüphaneler başarıyla import edildi!")
}
