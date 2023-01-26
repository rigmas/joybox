package utils

import (
	"testing"

	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func MockGetRawBooks() domain.RawBookData {
	var mockAuthors []*domain.RawAuthor
	mockAuthor := domain.RawAuthor{
		Name: "William Shakespeare",
	}
	mockAuthors = append(mockAuthors, &mockAuthor)

	var mockBooks []*domain.RawBook
	mockBook1 := domain.RawBook{
		Key:             "TEST/123",
		Title:           "Romeo and Juliet",
		CoverEditionKey: "NVM/123",
		Authors:         mockAuthors,
	}
	mockBooks = append(mockBooks, &mockBook1)

	books := domain.RawBookData{
		Books: mockBooks,
	}

	return books
}

func TestBookConverter_NoError(t *testing.T) {
	mockBooks := MockGetRawBooks()

	res := BookConverter(mockBooks)
	assert.Equal(t, "Romeo and Juliet", res[0].Title)
	assert.Equal(t, 1, len(res))
}
