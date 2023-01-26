package domain

type Book struct {
	Key           string
	Title         string
	Author        string
	EditionNumber string
}

type RawAuthor struct {
	Name string `json:"name"`
}

type RawBook struct {
	Key             string       `json:"key"`
	Title           string       `json:"title"`
	CoverEditionKey string       `json:"cover_edition_key"`
	Authors         []*RawAuthor `json:"authors"`
}

type RawBookData struct {
	Books []*RawBook `json:"works"`
}
