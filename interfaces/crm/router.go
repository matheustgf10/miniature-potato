package crm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheustgf10/miniature-potato/application/crm/order"
)

func Router(r fiber.Router) fiber.Router {
	orderRoute := order.Router()
	r.Group("/crm")

	return r
}
