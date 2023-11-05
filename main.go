package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/matheustgf10/miniature-potato/interfaces/crm"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/v1")

	//* CRM Routes aggregator
	crmRoutes := v1.Group("/crm")
	crm.Router(crmRoutes)

	//* Server Runner
	log.Fatal(app.Listen(":3000"))
}
