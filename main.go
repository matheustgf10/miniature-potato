package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Group("/v1", crm)

	log.Fatal(app.Listen(":3000"))
}
