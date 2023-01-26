package dto

import (
	"time"

	"github.com/rigmas/joybox/internal/core/domain"
)

type OrderDtoRequest struct {
	Key            string `json:"key,omitempty"`
	PickupSchedule string `json:"pickup_schedule,omitempty"`
}

type OrderDtoResponse struct {
	Book           BookDtoResponse `json:"book,omitempty"`
	PickupSchedule time.Time       `json:"pickup_schedule,omitempty"`
}

func ToOrderDtoRequest(key, pickupSchedule string) OrderDtoRequest {
	return OrderDtoRequest{
		Key:            key,
		PickupSchedule: pickupSchedule,
	}
}

func ToOrderDtoResponse(order domain.Order) OrderDtoResponse {
	return OrderDtoResponse{
		Book:           BookDtoResponse(order.Book),
		PickupSchedule: order.PickupSchedule,
	}
}
