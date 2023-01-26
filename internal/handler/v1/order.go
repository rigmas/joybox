package v1

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/rigmas/joybox/internal/core/port"
	"github.com/rigmas/joybox/internal/handler/v1/dto"
)

type orderRoute struct {
	OrderSrv port.OrderService
}

func NewOrderRoute(orderSrv port.OrderService) *orderRoute {
	return &orderRoute{
		OrderSrv: orderSrv,
	}
}

func (o *orderRoute) Add(c echo.Context) error {
	order := new(dto.OrderDtoRequest)
	if err := c.Bind(order); err != nil {
		resp := dto.NewResponseDto(err.Error(), *order, "orders")
		return c.JSON(http.StatusBadRequest, resp)
	}

	res, err := o.OrderSrv.Add(order.Key, order.PickupSchedule)
	if err != nil {
		resp := dto.NewResponseDto(err.Error(), dto.ToOrderDtoRequest(order.Key, order.PickupSchedule), "orders")
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := dto.NewResponseDto(dto.MSG_SUCCESS, dto.ToOrderDtoResponse(res), "orders")
	return c.JSON(http.StatusCreated, resp)
}

func (o *orderRoute) Get(c echo.Context) error {
	orders := o.OrderSrv.Get()

	resp := dto.NewResponseDto(dto.MSG_SUCCESS, dto.ToOrdersDto(orders...), "orders")

	return c.JSON(http.StatusOK, resp)
}
