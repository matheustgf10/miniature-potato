package order

type IOrder interface {
	ListOrders() (*ResOrder, error)
	AddOrder(*Order) error
	UpdateOrder(*Order) error
	DeleteOrder(*Order) error
}
