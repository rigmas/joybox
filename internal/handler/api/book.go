package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rigmas/joybox/internal/core/domain"
)

type BookApi struct {
	Client   *http.Client
	Endpoint string
}

func NewBookApi(client *http.Client, endpoint string) (*BookApi, error) {

	return &BookApi{
		Client:   client,
		Endpoint: endpoint,
	}, nil
}

func (b *BookApi) Fetch() (domain.RawBookData, error) {
	var data domain.RawBookData
	c := http.Client{}
	response, err := c.Get(b.Endpoint)
	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return data, err
	}

	return data, nil
}
