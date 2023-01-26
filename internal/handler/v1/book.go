package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rigmas/joybox/internal/core/port"
	"github.com/rigmas/joybox/internal/handler/v1/dto"
)

type bookRoute struct {
	BookSrv port.BookService
}

func NewBookRoute(bookSrv port.BookService) *bookRoute {
	return &bookRoute{
		BookSrv: bookSrv,
	}
}

func (b *bookRoute) Get(c echo.Context) error {
	books, err := b.BookSrv.Get()
	if err != nil {
		resp := dto.NewResponseDto(
			err.Error(),
			dto.ToBooksDto(nil...),
			"book",
		)
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := dto.NewResponseDto(dto.MSG_SUCCESS, dto.ToBooksDto(books...), "books")

	return c.JSON(http.StatusOK, resp)
}
