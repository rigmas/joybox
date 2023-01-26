package repository

import (
	"testing"

	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/rigmas/joybox/internal/utils"
	"github.com/rigmas/joybox/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func TestFetchOrder_NoError(t *testing.T) {
	mockOrders := generator.Order()
	orderRepo, err := NewOrderRepo(mockOrders)
	assert.NoError(t, err)

	res := orderRepo.Get()
	assert.Equal(t, 2, len(res))
	assert.Equal(t, "Romeo and Juliet", res[1].Book.Title)
	assert.Equal(t, "William Shakespeare", res[1].Book.Author)
	assert.Equal(t, "William Shakespeare", res[1].Book.Author)
}

func TestAddOrder_NoError(t *testing.T) {
	mockOrders := generator.Order()
	orderRepo, err := NewOrderRepo(mockOrders)
	assert.NoError(t, err)

	previousOrders := orderRepo.Get()
	assert.Equal(t, 2, len(previousOrders))

	mockBook := domain.Book{
		Key:           "/works/OL27776452W",
		Title:         "The Importance of Being Earnest",
		Author:        "Oscar Wilde",
		EditionNumber: "OL9694914M",
	}

	timeSchedule, _ := utils.ScheduleConverter("2023-02-01")
	mockOrder := domain.Order{
		Book:           mockBook,
		PickupSchedule: timeSchedule,
	}

	res, err := orderRepo.Add(mockOrder)
	assert.Equal(t, "Oscar Wilde", res.Book.Author)
	assert.Equal(t, "/works/OL27776452W", res.Book.Key)

	afterOrders := orderRepo.Get()
	assert.Equal(t, 3, len(afterOrders))
}
