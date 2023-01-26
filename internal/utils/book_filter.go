package utils

import (
	"fmt"

	"github.com/rigmas/joybox/internal/core/domain"
)

func BookFilter(key string, books []domain.Book) (domain.Book, error) {

	for _, book := range books {
		if book.Key == key {
			return book, nil
		}
	}

	return domain.Book{}, fmt.Errorf("Invalid key of book: %v", key)
}
