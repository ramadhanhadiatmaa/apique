package middlewares

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func AuthMiddleware(ctx *fiber.Ctx) error {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env files :%v", err)
	}

	secret_key := os.Getenv("SECRET_KEY")

	header := ctx.Get("apiKey")
	if header == "" || header != secret_key {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized Access Request",
		})
	}
	return ctx.Next()
}