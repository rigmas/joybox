package port

import (
	"github.com/rigmas/joybox/internal/core/domain"
)

type OrderRepository interface {
	Get() []domain.Order
	Add(domain.Order) (domain.Order, error)
}

type BookRepository interface {
	Get() ([]domain.Book, error)
}
