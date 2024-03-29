package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"johnbakhmat.tech/pinned/graphql"
)

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	
    app.Use(cors.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendStatus(200)
	})
	app.Get("/projects/:username", func(c fiber.Ctx) error {
		username := c.Params("username", "johnbakhmat")
		projects, err := graphql.GetProjects(username)
		if err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}
		log.Printf("%v", projects)
		return c.JSON(projects)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
