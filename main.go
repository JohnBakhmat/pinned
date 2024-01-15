package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portStr := os.Getenv("PORT")
    if portStr == "" {
        portStr = "7000"
    }

    port, err := strconv.Atoi(portStr)
    if err != nil {
        log.Fatal(err)
    }

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendStatus(200)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
