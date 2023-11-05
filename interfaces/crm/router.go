package crm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheustgf10/miniature-potato/interfaces/crm/order"
)

func Router(r fiber.Router) fiber.Router {
	orderRoutes := r.Group("/order")
	return order.Router(orderRoutes)
}
