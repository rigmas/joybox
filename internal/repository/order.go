package repository

import "github.com/rigmas/joybox/internal/core/domain"

type OrderRepo struct {
	Orders []domain.Order
}

func NewOrderRepo(orders []domain.Order) (*OrderRepo, error) {
	return &OrderRepo{
		Orders: orders,
	}, nil
}

func (o *OrderRepo) Add(order domain.Order) (domain.Order, error) {
	o.Orders = append(o.Orders, order)

	return order, nil
}

func (o *OrderRepo) Get() []domain.Order {
	return o.Orders
}
