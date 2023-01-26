package generator

import (
	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/rigmas/joybox/internal/utils"
)

func Order() []domain.Order {
	var mockOrders []domain.Order
	mockBook1 := domain.Book{
		Key:           "/works/OL21177W",
		Title:         "Wuthering Heights",
		Author:        "Emily Brontë",
		EditionNumber: "OL38586477M",
	}
	timeSchedule1, _ := utils.ScheduleConverter("2023-01-26")
	mockOrder1 := domain.Order{
		Book:           mockBook1,
		PickupSchedule: timeSchedule1,
	}

	mockBook2 := domain.Book{
		Key:           "/works/OL362427W",
		Title:         "Romeo and Juliet",
		Author:        "William Shakespeare",
		EditionNumber: "OL26501367M",
	}
	timeSchedule2, _ := utils.ScheduleConverter("2023-01-30")
	mockOrder2 := domain.Order{
		Book:           mockBook2,
		PickupSchedule: timeSchedule2,
	}

	mockOrders = append(mockOrders, mockOrder1, mockOrder2)

	return mockOrders

}
