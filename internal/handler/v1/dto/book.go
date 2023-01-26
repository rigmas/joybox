package dto

import "github.com/rigmas/joybox/internal/core/domain"

type BookDtoResponse struct {
	Key           string `json:"key,omitempty"`
	Title         string `json:"title,omitempty"`
	Author        string `json:"author,omitempty"`
	EditionNumber string `json:"edition_number,omitempty"`
}

func ToBookDtoResponse(book domain.Book) BookDtoResponse {
	return BookDtoResponse{
		Key:           book.Key,
		Title:         book.Title,
		Author:        book.Author,
		EditionNumber: book.EditionNumber,
	}
}

func ToBooksDto(books ...domain.Book) []BookDtoResponse {
	if len(books) == 0 || books == nil {
		return make([]BookDtoResponse, 0)
	}
	var result []BookDtoResponse
	for _, book := range books {
		temp := ToBookDtoResponse(book)
		result = append(result, temp)
	}

	return result
}
