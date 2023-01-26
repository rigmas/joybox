package service

import (
	"testing"
	"time"

	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/rigmas/joybox/internal/core/port"
	"github.com/rigmas/joybox/pkg/generator"
	"github.com/stretchr/testify/assert"
)

type MockBookSrv struct{}

func NewMockBookSrv() port.BookService {
	return &MockBookSrv{}
}

func (b *MockBookSrv) Get() ([]domain.Book, error) {
	var mockBooks []domain.Book
	mockBook1 := domain.Book{
		Key:           "/works/OL21177W",
		Title:         "Wuthering Heights",
		Author:        "Emily Brontë",
		EditionNumber: "OL38586477M",
	}

	mockBook2 := domain.Book{
		Key:           "/works/OL362427W",
		Title:         "Romeo and Juliet",
		Author:        "William Shakespeare",
		EditionNumber: "OL26501367M",
	}

	mockBooks = append(mockBooks, mockBook1, mockBook2)

	return mockBooks, nil
}

func (b *MockBookSrv) GetOne(key string) (domain.Book, error) {
	return domain.Book{}, nil
}

type MockOrderRepo struct{}

func NewMockOrderRepo() port.OrderRepository {
	return &MockOrderRepo{}
}

func (o *MockOrderRepo) Add(order domain.Order) (domain.Order, error) {
	return domain.Order{}, nil
}

func (o *MockOrderRepo) Get() []domain.Order {
	orders := generator.Order()
	return orders
}

func TestAddOrder_NoError(t *testing.T) {
	mockBookSrv := NewMockBookSrv()
	mockOrderRepo := NewMockOrderRepo()

	orderSrv, err := NewOrderService(mockBookSrv, mockOrderRepo)
	assert.NoError(t, err)

	result, err := orderSrv.Add("/works/OL21177W", "2023-01-26")
	assert.NoError(t, err)
	assert.Equal(t, "Wuthering Heights", result.Book.Title)
}

func TestAddOrder_ErrorInvalidDate(t *testing.T) {
	mockBookSrv := NewMockBookSrv()
	mockOrderRepo := NewMockOrderRepo()

	orderSrv, err := NewOrderService(mockBookSrv, mockOrderRepo)
	assert.NoError(t, err)

	result, err := orderSrv.Add("/works/OL21177W", "2023-02-31")
	assert.NotNil(t, err)
	assert.Equal(
		t,
		domain.Order{Book: domain.Book{}, PickupSchedule: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		result,
	)
}

func TestAddOrder_ErrorInvalidBookKey(t *testing.T) {
	mockBookSrv := NewMockBookSrv()
	mockOrderRepo := NewMockOrderRepo()

	orderSrv, err := NewOrderService(mockBookSrv, mockOrderRepo)
	assert.NoError(t, err)

	result, err := orderSrv.Add("/TEST/123", "2023-01-30")
	assert.NotNil(t, err)
	assert.Equal(
		t,
		domain.Order{Book: domain.Book{}, PickupSchedule: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		result,
	)
}

func TestFetchOrder_NoError(t *testing.T) {
	mockBookSrv := NewMockBookSrv()
	mockOrderRepo := NewMockOrderRepo()

	orderSrv, err := NewOrderService(mockBookSrv, mockOrderRepo)
	assert.NoError(t, err)

	result := orderSrv.Get()
	assert.NoError(t, err)
	assert.Equal(t, "Romeo and Juliet", result[1].Book.Title)
}
