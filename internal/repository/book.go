package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BookRepository struct {
	Client   *http.Client
	Endpoint string
}

func NewBookRepository(client *http.Client, endpoint string) (*BookRepository, error) {
	return &BookRepository{
		Client:   client,
		Endpoint: endpoint,
		// Endpoint: "http://openlibrary.org/subjects/love.json?published_in=1500-1900",
	}, nil
}

func (b *BookRepository) Get() (map[string]string, error) {
	// var books []domain.Book

	c := http.Client{}
	response, err := c.Get(b.Endpoint)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	// err = json.NewDecoder(response.Body).Decode(&books)

	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	err = json.Unmarshal(body, &data)

	return data, nil

	// request, err := http.NewRequest("GET", "http://openlibrary.org/subjects/love.json?published_in=1500-1900", nil)

	// if err != nil {
	// 	return books, err
	// }

	// response, err := b.client.Do(request)
	// if err != nil {
	// 	return books, err
	// }

	// defer response.Body.Close()

	// err = json.NewDecoder(response.Body).Decode(&books)
	// if err != nil {
	// 	return books, err
	// }

	// return books, nil
}
