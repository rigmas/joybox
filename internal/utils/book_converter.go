package utils

import "github.com/rigmas/joybox/internal/core/domain"

func BookConverter(rawData domain.RawBookData) []domain.Book {
	var books []domain.Book

	for _, rawBook := range rawData.Books {
		book := domain.Book{
			Key:           rawBook.Key,
			Title:         rawBook.Title,
			Author:        rawBook.Authors[0].Name,
			EditionNumber: rawBook.CoverEditionKey,
		}

		books = append(books, book)
	}

	return books
}
