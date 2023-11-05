package order

import "github.com/matheustgf10/miniature-potato/utils"

func ListOrders() (res *ResOrder, err error) {
	res = new(ResOrder)

	//! Mock
	res.Data = append(res.Data, &Order{
		ID:    utils.GetStringPointer("550e8400-e29b-41d4-a716-446655440000"),
		Title: utils.GetStringPointer("MC Donalds Food"),
		Type:  utils.GetStringPointer("Food"),
	})

	return
}

func AddOrder(order *Order) (err error) {
	return
}

func UpdateOrder(order *Order) (err error) {
	return
}

func DeleteOrder(order *Order) (err error) {
	return
}
