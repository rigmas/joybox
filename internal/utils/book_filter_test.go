package utils

import (
	"fmt"
	"testing"

	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func MockGetBooks() []domain.Book {
	var mockBooks []domain.Book
	mockBook1 := domain.Book{
		Key:           "/works/OL21177W",
		Title:         "Wuthering Heights",
		Author:        "Emily Brontë",
		EditionNumber: "OL38586477M",
	}

	mockBook2 := domain.Book{
		Key:           "/works/OL362427W",
		Title:         "Romeo and Juliet",
		Author:        "William Shakespeare",
		EditionNumber: "OL26501367M",
	}

	mockBooks = append(mockBooks, mockBook1, mockBook2)

	return mockBooks
}

func TestBookFilter_NoError(t *testing.T) {
	key := "/works/OL362427W"
	mockBooks := MockGetBooks()

	res, err := BookFilter(key, mockBooks)
	assert.NoError(t, err)
	assert.Equal(t, "Romeo and Juliet", res.Title)
}

func TestBookFilter_ErrorInvalidBookKey(t *testing.T) {
	key := "/works/OL362427Wxxx"
	mockBooks := MockGetBooks()

	res, err := BookFilter(key, mockBooks)
	assert.Equal(t, fmt.Errorf("Invalid key of book: %v", key), err)
	assert.Equal(t, domain.Book{}, res)
}
