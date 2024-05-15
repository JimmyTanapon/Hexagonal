package core

type OrderRepository interface {
	SaveOrder(order Order) error
}
