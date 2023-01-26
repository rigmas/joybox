package repository

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBook(t *testing.T) {
	// server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter,
	// 	req *http.Request) {
	// 	rw.Write([]byte(`{Result: [{Cluster_name: "cl1", Pings: 2}]}`))
	// }))

	// defer server.Close()

	var client *http.Client

	bookRepo, err := NewBookRepository(client, "http://openlibrary.org/subjects/love.json?published_in=1500-1900")
	assert.NoError(t, err)

	res, err := bookRepo.Get()
	assert.NoError(t, err)
	assert.Equal(t, "/subjects/love", res["key"])
}
