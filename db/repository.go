package db

import (
	"context"

	"github.com/leogsouza/meower/schema"
)

// Repository represents database interactions
type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

// SetRepository sets a repository. By using Repository interface you allow any concrete
// implementation to be injected at runtime, and all function calls will be delegated to the impl object.
func SetRepository(repository Repository) {
	impl = repository
}

// Close closes database connection
func Close() {
	impl.Close()
}

// InsertMeow inserts a new Meow into database
func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

// ListMeows list all meows from database
// Considering skip and take params
func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.ListMeows(ctx, skip, take)
}
