/*
Package repository holds event sourced repositories
*/
package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hsynrtn/dashboard-management/cmd/dashboard/internal/domain/city"
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	"github.com/hsynrtn/dashboard-management/pkg/eventbus"
	"github.com/hsynrtn/dashboard-management/pkg/eventstore"
)

type cityRepository struct {
	eventStore eventstore.EventStore
	eventBus   eventbus.EventBus
}

// Save current city changes to event store and publish each event with an event bus
func (r *cityRepository) Save(ctx context.Context, u city.City) error {
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

// Get city with current state applied
func (r *cityRepository) Get(ctx context.Context, id uuid.UUID) (city.City, error) {
	events, err := r.eventStore.GetStream(ctx, id, city.StreamName)
	if err != nil {
		return city.City{}, apperrors.Wrap(err)
	}

	if len(events) == 0 {
		return city.City{}, apperrors.ErrNotFound
	}

	return city.FromHistory(ctx, events)
}

// NewTokenRepository creates new city event sourced repository
func NewCityRepository(store eventstore.EventStore, bus eventbus.EventBus) city.Repository {
	return &cityRepository{store, bus}
}
