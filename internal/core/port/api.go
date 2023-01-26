package port

import "github.com/rigmas/joybox/internal/core/domain"

type BookApi interface {
	Fetch() (domain.RawBookData, error)
}
