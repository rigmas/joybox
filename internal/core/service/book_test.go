package service

import (
	"testing"

	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/rigmas/joybox/internal/core/port"
	"github.com/stretchr/testify/assert"
)

type MockBookApi struct{}

func NewMockBookApi() port.BookApi {
	return &MockBookApi{}
}

func (b *MockBookApi) Fetch() (domain.RawBookData, error) {
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

	return books, nil
}

func TestGetBookList_NoError(t *testing.T) {
	mockBookApi := NewMockBookApi()

	bookSrv, err := NewBookService(mockBookApi)
	assert.NoError(t, err)

	result, err := bookSrv.Get()
	assert.NoError(t, err)
	assert.Equal(t, "Romeo and Juliet", result[0].Title)
}
