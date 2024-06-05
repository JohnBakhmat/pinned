package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/patrickmn/go-cache"
	"johnbakhmat.tech/pinned/graphql"
)

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "80"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err)
	}

	c := cache.New(5*time.Minute, 10*time.Minute)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendStatus(200)
	})
	app.Get("/projects/:username", func(ctx fiber.Ctx) error {
		username := ctx.Params("username", "johnbakhmat")

		if projects_cache, cache_hit := c.Get(username); cache_hit {
			log.Println("Cache hit")
			projects := projects_cache.(*[]graphql.Project)
			return ctx.JSON(*projects)
		}

		projects, err := graphql.GetProjects(username)
		if err != nil {
			log.Println(err)
			return ctx.SendStatus(500)
		}
		log.Printf("%v", projects)
		c.Set(username, &projects, cache.DefaultExpiration)
		return ctx.JSON(projects)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
