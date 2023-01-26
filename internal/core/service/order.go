package service

import (
	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/rigmas/joybox/internal/core/port"
	"github.com/rigmas/joybox/internal/utils"
)

type orderService struct {
	BookSrv   port.BookService
	RepoOrder port.OrderRepository
}

func NewOrderService(bookSrv port.BookService,
	repo port.OrderRepository) (*orderService, error) {
	return &orderService{
		BookSrv:   bookSrv,
		RepoOrder: repo,
	}, nil
}

func (o *orderService) Add(key string, pickupSchedule string) (domain.Order, error) {
	timeSchedule, err := utils.ScheduleConverter(pickupSchedule)
	if err != nil {
		return domain.Order{}, err
	}

	books, err := o.BookSrv.Get()
	if err != nil {
		return domain.Order{}, err
	}

	book, err := utils.BookFilter(key, books)
	if err != nil {
		return domain.Order{}, err
	}

	order := domain.Order{
		PickupSchedule: timeSchedule,
		Book:           book,
	}

	return order, nil
}

func (o *orderService) Get() []domain.Order {
	return o.RepoOrder.Get()
}
