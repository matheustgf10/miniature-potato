package order

import "github.com/gofiber/fiber/v2"

func Router(r fiber.Router) fiber.Router {
	r.Get("/list", listOrders)
	r.Post("/", addOrder)
	r.Put("/:order_id", updateOrder)
	r.Delete("/:order_id", deleteOrder)

	return r
}
