package search

import (
	"context"

	"github.com/leogsouza/meower/schema"
)

// Repository is an interface for handling search operations
type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

// SetRepository sets the repository through dependency injection
func SetRepository(repository Repository) {
	impl = repository
}

// Close ends the elastic search connection
func Close() {
	impl.Close()
}

// InsertMeow inserts the type into search engine
func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

// SearchMeows seeks for Meows into search engine
func SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.SearchMeows(ctx, query, skip, take)
}
