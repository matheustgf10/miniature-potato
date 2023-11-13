package order

import (
	"database/sql"

	domain "github.com/matheustgf10/miniature-potato/domain/crm/order"
)

type PGOrder struct {
	//! Tx é uma transação
	Tx *sql.Tx
}

func (pg *PGOrder) NewRepo(tx *sql.Tx) domain.IOrder {
	return &PGOrder{Tx: tx}
}

func (pg *PGOrder) ListOrders() (res *domain.ResOrder, err error) { return pg.ListOrders() }
func (pg *PGOrder) AddOrder(*domain.Order) error                  { return nil }
func (pg *PGOrder) UpdateOrder(*domain.Order) error               { return nil }
func (pg *PGOrder) DeleteOrder(*domain.Order) error               { return nil }
