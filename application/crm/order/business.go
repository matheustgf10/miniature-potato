package order

import (
	"log"

	"github.com/matheustgf10/miniature-potato/config/database"
)

func ListOrders() (res *ResOrder, err error) {
	res = new(ResOrder)

	tx, err := database.NewTx()
	if err != nil {
		return
	}
	defer tx.Rollback()

	order := new(Order)
	if err = tx.
		QueryRow("SELECT O.id, O.title, O.type FROM public.order AS O").
		Scan(
			&order.ID,
			&order.Title,
			&order.Type,
		); err != nil {
		log.Fatalln(err)
		return
	}

	res.Data = append(res.Data, order)

	return
}

func AddOrder(order *Order) (err error) {
	tx, err := database.NewTx()
	if err != nil {
		return
	}
	defer tx.Rollback()
	if _, err = tx.Exec(
		"INSERT INTO public.order(title, description, type) VALUES('teste', 'teste', 'Food')"); err != nil {
		log.Fatalln(err)
		return
	}

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
