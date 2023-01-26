package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/rigmas/joybox/internal/core/port"
)

func Route(e *echo.Group, bookSrv port.BookService,
	orderSrv port.OrderService) {

	bookRoute := NewBookRoute(bookSrv)
	orderRoute := NewOrderRoute(orderSrv)

	eBook := e.Group("/books")
	eBook.GET("/", bookRoute.Get)

	eOrder := e.Group("/orders")
	eOrder.POST("/", orderRoute.Add)
	eOrder.GET("/", orderRoute.Get)
}
