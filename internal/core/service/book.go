package service

import (
	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/rigmas/joybox/internal/core/port"
	"github.com/rigmas/joybox/internal/utils"
)

type bookService struct {
	apiBook port.BookApi
}

func NewBookService(api port.BookApi) (*bookService, error) {
	return &bookService{
		apiBook: api,
	}, nil
}

func (b *bookService) Get() ([]domain.Book, error) {
	var books []domain.Book
	rawData, err := b.apiBook.Fetch()
	if err != nil {
		return nil, err
	}

	books = utils.BookConverter(rawData)

	return books, nil
}
