package order

import (
	"net/http"

	app "github.com/matheustgf10/miniature-potato/application/crm/order"

	"github.com/gofiber/fiber/v2"
)

// @Router crm/order/list [get]
func listOrders(ctx *fiber.Ctx) error {
	res, err := app.ListOrders()
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Router crm/order [post]
func addOrder(ctx *fiber.Ctx) (err error) {
	var order app.Order
	if err = app.AddOrder(&order); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(nil)
}

// @Router crm/order/:order_id [put]
func updateOrder(ctx *fiber.Ctx) (err error) {
	var order app.Order
	if err = app.UpdateOrder(&order); err != nil {
		return err
	}

	return ctx.Status(http.StatusNoContent).JSON(nil)
}

// @Router crm/order/:order_id [delete]
func deleteOrder(ctx *fiber.Ctx) (err error) {
	var order app.Order
	if err = app.DeleteOrder(&order); err != nil {
		return err
	}

	return ctx.Status(http.StatusNoContent).JSON(nil)
}
