package api

import (
	"net/http"
	"testing"

	"github.com/rigmas/joybox/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestFetchBookApi_NoError(t *testing.T) {
	var client *http.Client

	bookApi, err := NewBookApi(client, "http://openlibrary.org/subjects/love.json?published_in=1500-1900")
	assert.NoError(t, err)

	res, err := bookApi.Fetch()
	assert.NoError(t, err)
	assert.Equal(t, 12, len(res.Books))
}

func TestFetchBookApi_Error(t *testing.T) {
	var result domain.RawBookData
	var client *http.Client

	bookApi, err := NewBookApi(client, "asd")
	assert.NoError(t, err)

	res, err := bookApi.Fetch()
	assert.Equal(t, result, res)
	assert.NotNil(t, err)
}
