package port

import (
	"github.com/rigmas/joybox/internal/core/domain"
)

type BookService interface {
	Get() ([]domain.Book, error)
	GetOne(key string) (domain.Book, error)
}

type OrderService interface {
	Add(key string) (domain.Order, error)
	Get() ([]domain.Order, error)
}
