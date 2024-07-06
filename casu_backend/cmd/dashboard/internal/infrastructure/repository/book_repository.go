/*
Package repository holds event sourced repositories
*/
package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/book"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"github.com/hsynrtn/dashboard-management/pkg/eventstore"
)

type bookRepository struct {
	eventStore eventstore.EventStore
	eventBus   eventbus.EventBus
}

// Save current book changes to event store and publish each event with an event bus
func (r *bookRepository) Save(ctx context.Context, u book.Book) error {
	if err := r.eventStore.Store(ctx, u.Changes()); err != nil {
		return apperrors.Wrap(err)
	}

	for _, event := range u.Changes() {
		if err := r.eventBus.Publish(ctx, event); err != nil {
			return apperrors.Wrap(err)
		}
	}

	return nil
}

// Get book with current state applied
func (r *bookRepository) Get(ctx context.Context, id uuid.UUID) (book.Book, error) {
	events, err := r.eventStore.GetStream(ctx, id, book.StreamName)
	if err != nil {
		return book.Book{}, apperrors.Wrap(err)
	}

	if len(events) == 0 {
		return book.Book{}, apperrors.ErrNotFound
	}

	return book.FromHistory(ctx, events)
}

// NewBookRepository creates new book event sourced repository
func NewBookRepository(store eventstore.EventStore, bus eventbus.EventBus) book.Repository {
	return &bookRepository{store, bus}
}
