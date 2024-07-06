/*
Package repository holds event sourced repositories
*/
package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/country"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"github.com/hsynrtn/dashboard-management/pkg/eventstore"
)

type countryRepository struct {
	eventStore eventstore.EventStore
	eventBus   eventbus.EventBus
}

// Save current country changes to event store and publish each event with an event bus
func (r *countryRepository) Save(ctx context.Context, u country.Country) error {
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

// Get country with current state applied
func (r *countryRepository) Get(ctx context.Context, id uuid.UUID) (country.Country, error) {
	events, err := r.eventStore.GetStream(ctx, id, country.StreamName)
	if err != nil {
		return country.Country{}, apperrors.Wrap(err)
	}

	if len(events) == 0 {
		return country.Country{}, apperrors.ErrNotFound
	}

	return country.FromHistory(ctx, events)
}

// NewCountryRepository creates new country event sourced repository
func NewCountryRepository(store eventstore.EventStore, bus eventbus.EventBus) country.Repository {
	return &countryRepository{store, bus}
}
