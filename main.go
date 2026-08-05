package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = validator.New()
	_ = jwt.New(jwt.SigningMethodHS256)
	_ = godotenv.Load()
	_, _ = bcrypt.GenerateFromPassword([]byte("test"), 10)
	_, _ = gorm.Open(postgres.Open("dsn"), &gorm.Config{})

	corsOrigin := os.Getenv("CORS_ALLOWED_ORIGINS")
	if corsOrigin == "" {

		log.Fatalf("Error: Required environment variable missing or invalid: CORS_ALLOWED_ORIGINS")
	}

	app := fiber.New()

	// k6
	app.Post("/api/v1/auth/register", func(c *fiber.Ctx) error {
		return c.Status(201).JSON(fiber.Map{"message": "sahte kayit basarili"})
	})

	fmt.Println("API 3000 portunda başlıyor!")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
