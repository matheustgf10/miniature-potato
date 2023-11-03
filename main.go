package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/v1")
	v1 := api.Group("/api", handler)
	v1.Get("list", handler)

	log.Fatal(app.Listen(":3000"))
}
