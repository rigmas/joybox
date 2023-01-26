package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/rigmas/joybox/internal/core/port"
)

func Route(e *echo.Group, bookSrv port.BookService) {

	bookRoute := NewBookRoute(bookSrv)

	eBook := e.Group("/books")
	eBook.GET("/", bookRoute.Get)
}
