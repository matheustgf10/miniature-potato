package order

import (
	"github.com/matheustgf10/miniature-potato/config/database"
	"github.com/matheustgf10/miniature-potato/utils"
)

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
	tx, err := database.NewTx()
	if err != nil {
		return
	}
	defer tx.Rollback()

	print("// REGISTRAR NO BANCO DE DADOS")

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func UpdateOrder(order *Order) (err error) {

	return
}

func DeleteOrder(order *Order) (err error) {
	return
}
